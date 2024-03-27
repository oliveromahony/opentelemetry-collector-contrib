// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nginxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/experimental/storage"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/metadata"
)

type nginxLogScraper struct {
	cancel   context.CancelFunc
	cfg      *Config
	emit     *adapter.LogEmitter
	logger   *zap.Logger
	mb       *metadata.MetricsBuilder
	pipe     *pipeline.DirectedPipeline
	settings component.TelemetrySettings
	wg       *sync.WaitGroup
}

func newNginxLogScraper(
	settings receiver.CreateSettings,
	cfg *Config,
) (*nginxLogScraper, error) {
	mb := metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings)
	inputCfg := operator.NewConfig(&cfg.InputConfig)
	sugaredLogger := settings.Logger.Sugar()

	operators := append([]operator.Config{inputCfg}, cfg.BaseConfig.Operators...)

	emitter := adapter.NewLogEmitter(sugaredLogger)
	pipe, err := pipeline.Config{
		Operators:     operators,
		DefaultOutput: emitter,
	}.Build(sugaredLogger)
	if err != nil {
		return nil, err
	}

	return &nginxLogScraper{
		cfg:      cfg,
		emit:     emitter,
		logger:   settings.Logger,
		mb:       mb,
		pipe:     pipe,
		settings: settings.TelemetrySettings,
		wg:       &sync.WaitGroup{},
	}, nil
}

func (nls *nginxLogScraper) start(parentCtx context.Context, host component.Host) error {
	err := nls.pipe.Start(storage.NewNopClient())
	if err != nil {
		return fmt.Errorf("start stanza: %w", err)
	}

	// nls.wg.Add(1)
	// go nls.runProducer(ctx)

	// nls.wg.Add(1)
	// go nls.runConsumer(ctx)

	return nil
}

func (nls *nginxLogScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	now := pcommon.NewTimestampFromTime(time.Now())
	entryChan := nls.emit.OutChannel()
	// Unclear whether just a `select` will be enough here. May need to e.g. consume events from channel with a timer.
	select {
	case <-ctx.Done():
		return pmetric.Metrics{}, nil
	case entries, ok := <-entryChan:
		if !ok {
			nls.logger.Debug("Emitter channel closed!")
			return pmetric.Metrics{}, nil
		}

		for _, entry := range entries {
			nls.recordMetrics(ctx, entry, now)
		}
	}

	return nls.mb.Emit(), nil
}

func (nls *nginxLogScraper) recordMetrics(ctx context.Context, e *entry.Entry, now pcommon.Timestamp) {
	// TODO: Transform `stanza` entry into `pmetric.Metrics`

	nls.mb.RecordNginxUpstreamsResponseDataPoint(now, 0.0, metadata.AttributeResponsesFailed)
}
