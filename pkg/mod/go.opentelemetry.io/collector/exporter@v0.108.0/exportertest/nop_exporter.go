// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exportertest // import "go.opentelemetry.io/collector/exporter/exportertest"

import (
	"context"

	"github.com/google/uuid"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/exporter"
)

var nopType = component.MustNewType("nop")

// NewNopSettings returns a new nop settings for Create*Exporter functions.
func NewNopSettings() exporter.Settings {
	return exporter.Settings{
		ID:                component.NewIDWithName(nopType, uuid.NewString()),
		TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		BuildInfo:         component.NewDefaultBuildInfo(),
	}
}

// NewNopFactory returns an exporter.Factory that constructs nop exporters.
func NewNopFactory() exporter.Factory {
	return exporter.NewFactory(
		nopType,
		func() component.Config { return &nopConfig{} },
		exporter.WithTraces(createTracesExporter, component.StabilityLevelStable),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelStable),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelStable),
	)
}

func createTracesExporter(context.Context, exporter.Settings, component.Config) (exporter.Traces, error) {
	return nopInstance, nil
}

func createMetricsExporter(context.Context, exporter.Settings, component.Config) (exporter.Metrics, error) {
	return nopInstance, nil
}

func createLogsExporter(context.Context, exporter.Settings, component.Config) (exporter.Logs, error) {
	return nopInstance, nil
}

type nopConfig struct{}

var nopInstance = &nopExporter{
	Consumer: consumertest.NewNop(),
}

// nopExporter stores consumed traces and metrics for testing purposes.
type nopExporter struct {
	component.StartFunc
	component.ShutdownFunc
	consumertest.Consumer
}

// NewNopBuilder returns an exporter.Builder that constructs nop receivers.
//
// Deprecated: [v0.108.0] this builder is being internalized within the service module,
// and will be removed soon.
func NewNopBuilder() *exporter.Builder {
	nopFactory := NewNopFactory()
	return exporter.NewBuilder(
		map[component.ID]component.Config{component.NewID(nopType): nopFactory.CreateDefaultConfig()},
		map[component.Type]exporter.Factory{nopType: nopFactory})
}
