// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package accesslog

import (
	"context"
	"fmt"
	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/experimental/storage"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/model"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/record"
)

type (
	grokParser interface {
		ParseString(text string) map[string]string
	}

	NginxLogScraper struct {
		cancel  context.CancelFunc
		cfg     *config.Config
		entries []*entry.Entry
		grok    grokParser
		logger  *zap.SugaredLogger
		mb      *metadata.MetricsBuilder
		mut     sync.Mutex
		outChan <-chan []*entry.Entry
		pipe    *pipeline.DirectedPipeline
		wg      *sync.WaitGroup
	}
)

var _ scraperhelper.Scraper = (*NginxLogScraper)(nil)

func NewScraper(
	settings receiver.CreateSettings,
	cfg *config.Config,
) (*NginxLogScraper, error) {
	mb := metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings)
	inputCfg := operator.NewConfig(&cfg.InputConfig)
	logger := settings.Logger.Sugar()

	stanzaPipeline, outChan, err := initStanzaPipeline(cfg, inputCfg, logger)
	if err != nil {
		return nil, fmt.Errorf("init stanza pipeline: %w", err)
	}

	logFormat, err := logFormatFromNginxConf(cfg.NginxConfigPath)
	if err != nil {
		return nil, fmt.Errorf("NGINX log format missing: %w", err)
	}
	logger.Debugf("Using log format: %s", logFormat)

	grok, err := newGrok(logFormat, logger)
	if err != nil {
		return nil, fmt.Errorf("grok init: %w", err)
	}

	return &NginxLogScraper{
		cfg:     cfg,
		grok:    grok,
		logger:  logger,
		mb:      mb,
		mut:     sync.Mutex{},
		outChan: outChan,
		pipe:    stanzaPipeline,
		wg:      &sync.WaitGroup{},
	}, nil
}

func (nls *NginxLogScraper) ID() component.ID {
	return component.NewID(metadata.Type)
}

func (nls *NginxLogScraper) Start(parentCtx context.Context, _ component.Host) error {
	nls.logger.Debug("NGINX access log scraper started")
	ctx, cancel := context.WithCancel(parentCtx)
	nls.cancel = cancel

	err := nls.pipe.Start(storage.NewNopClient())
	if err != nil {
		return fmt.Errorf("stanza pipeline start: %w", err)
	}

	nls.wg.Add(1)
	go nls.runConsumer(ctx)

	return nil
}

func (nls *NginxLogScraper) Scrape(_ context.Context) (pmetric.Metrics, error) {
	nls.mut.Lock()
	for _, ent := range nls.entries {
		strBody, ok := ent.Body.(string)
		if !ok {
			nls.logger.Debugf("Failed to cast log entry to string, %v", ent.Body)
			continue
		}

		mappedResults := nls.grok.ParseString(strBody)

		ai, err := newNginxAccessItem(mappedResults)
		if err != nil {
			return pmetric.Metrics{}, fmt.Errorf("cast grok map to access item: %w", err)
		}

		err = record.Item(ai, nls.mb)
		if err != nil {
			nls.logger.Debugf("Recording metric failed, %+v", ai)
			continue
		}

	}
	nls.entries = make([]*entry.Entry, 0)
	nls.mut.Unlock()

	return nls.mb.Emit(), nil
}

func (nls *NginxLogScraper) Shutdown(_ context.Context) error {
	nls.logger.Info("Shutting down NGINX access log scraper")
	nls.cancel()
	nls.wg.Wait()

	return nls.pipe.Stop()
}

func initStanzaPipeline(
	baseCfg *config.Config,
	opCfg operator.Config,
	logger *zap.SugaredLogger,
) (*pipeline.DirectedPipeline, <-chan []*entry.Entry, error) {
	operators := append([]operator.Config{opCfg}, baseCfg.BaseConfig.Operators...)

	emitter := adapter.NewLogEmitter(logger)
	pipe, err := pipeline.Config{
		Operators:     operators,
		DefaultOutput: emitter,
	}.Build(logger)

	return pipe, emitter.OutChannel(), err
}

func (nls *NginxLogScraper) runConsumer(ctx context.Context) {
	nls.logger.Debug("Starting NGINX access log receiver's consumer")
	defer nls.wg.Done()

	entryChan := nls.outChan
	for {
		select {
		case <-ctx.Done():
			nls.logger.Debug("Closing NGINX access log receiver consumer")
			return
		case entries, ok := <-entryChan:
			if !ok {
				nls.logger.Debug("Emitter channel closed, shutting down NGINX access log consumer")
				return
			}

			nls.mut.Lock()
			nls.entries = append(nls.entries, entries...)
			nls.mut.Unlock()
		}
	}
}

func newNginxAccessItem(v map[string]string) (*model.NginxAccessItem, error) {
	res := &model.NginxAccessItem{}
	if err := mapstructure.Decode(v, res); err != nil {
		return nil, err
	}
	return res, nil
}
