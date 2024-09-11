package emptyexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/pmetrics"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	typeStr = "emptyexporter"
)

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		component.NewID(typeStr),
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, component.StabilityLevelDevelopment),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelDevelopment),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelDevelopment),
	)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	config component.Config) (exporter.Traces, error) {

	cfg := config.(*Config)
	s := NewEmptyexporter()
	return exporterhelper.NewTracesExporter(ctx, set, cfg, s.pushTraces)
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	config component.Config) (exporter.Metrics, error) {

	cfg := config.(*Config)
	s := NewEmptyexporter()
	return exporterhelper.NewMetricsExporter(ctx, set, cfg, s.pushMetrics)
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	config component.Config) (exporter.Logs, error) {

	cfg := config.(*Config)
	s := NewEmptyexporter()
	return exporterhelper.NewLogsExporter(ctx, set, cfg, s.pushLogs)
}

type Config struct {
}

func createDefaultConfig() component.Config {
	return &Config{}
}

type Emptyexporter struct {
}

func NewEmptyexporter() *Emptyexporter {
	return &Emptyexporter{}
}

func (e *Emptyexporter) pushTraces(ctx context.Context, traces ptrace.Traces) error {
	// Implement trace exporting logic here
	return nil
}

func (e *Emptyexporter) pushMetrics(ctx context.Context, metrics pmetrics.Metrics) error {
	// Implement metric exporting logic here
	return nil
}

func (e *Emptyexporter) pushLogs(ctx context.Context, logs plog.Logs) error {
	// Implement log exporting logic here
	return nil
}
