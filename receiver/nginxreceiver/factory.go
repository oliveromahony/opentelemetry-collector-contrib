// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nginxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/file"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/metadata"
)

// NewFactory creates a factory for nginx receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

func createDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()
	cfg.CollectionInterval = 10 * time.Second

	return &Config{
		ControllerConfig: cfg,
		ClientConfig: confighttp.ClientConfig{
			Endpoint: "http://localhost:80/status",
			Timeout:  10 * time.Second,
		},
		InputConfig:          *file.NewConfig(),
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
	}
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	rConf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	cfg := rConf.(*Config)

	ns := newNginxStubStatusScraper(params, cfg)
	stubStatusScraper, err := scraperhelper.NewScraper(metadata.Type.String(), ns.scrape, scraperhelper.WithStart(ns.start))
	if err != nil {
		return nil, err
	}

	nals := newNginxLogScraper(params, cfg)
	accessLogScraper, err := scraperhelper.NewScraper(metadata.Type.String(), nals.scrape, scraperhelper.WithStart(nals.start))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewScraperControllerReceiver(
		&cfg.ControllerConfig, params, consumer,
		scraperhelper.AddScraper(stubStatusScraper),
		scraperhelper.AddScraper(accessLogScraper),
		// scraperhelper.AddScraper(errorLogScraper),
	)
}
