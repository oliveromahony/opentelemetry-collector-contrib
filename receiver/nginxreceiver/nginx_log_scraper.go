// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nginxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"

import (
	"context"
	// "net/http"
	"time"

	// "github.com/nginxinc/nginx-prometheus-exporter/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	// "go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/metadata"
)

type nginxLogScraper struct {
	settings component.TelemetrySettings
	cfg      *Config
	mb       *metadata.MetricsBuilder
}

func newNginxLogScraper(
	settings receiver.CreateSettings,
	cfg *Config,
) *nginxLogScraper {
	mb := metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings)
	return &nginxLogScraper{
		settings: settings.TelemetrySettings,
		cfg:      cfg,
		mb:       mb,
	}
}

func (r *nginxLogScraper) start(_ context.Context, host component.Host) error {
	// httpClient, err := r.cfg.ToClient(host, r.settings)
	// if err != nil {
	// 	return err
	// }
	// r.httpClient = httpClient

	return nil
}

func (r *nginxLogScraper) scrape(context.Context) (pmetric.Metrics, error) {
	// Init client in scrape method in case there are transient errors in the constructor.
	// if r.client == nil {
	// 	var err error
	// 	r.client, err = client.NewNginxClient(r.httpClient, r.cfg.ClientConfig.Endpoint)
	// 	if err != nil {
	// 		r.client = nil
	// 		return pmetric.Metrics{}, err
	// 	}
	// }

	// stats, err := r.client.GetStubStats()
	// if err != nil {
	// 	r.settings.Logger.Error("Failed to fetch nginx stats", zap.Error(err))
	// 	return pmetric.Metrics{}, err
	// }

	now := pcommon.NewTimestampFromTime(time.Now())

	r.mb.RecordNginxUpstreamsResponseDataPoint(now, 0.0, metadata.AttributeResponsesBuffered)
	r.mb.RecordNginxUpstreamsResponseDataPoint(now, 0.0, metadata.AttributeResponsesFailed)

	return r.mb.Emit(), nil
}