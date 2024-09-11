// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry // import "go.opentelemetry.io/collector/service/telemetry"

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace/noop"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/internal/globalgates"
	"go.opentelemetry.io/collector/service/telemetry/internal"
)

func TestAttributes(t *testing.T) {
	tests := []struct {
		name           string
		cfg            Config
		buildInfo      component.BuildInfo
		wantAttributes map[string]interface{}
	}{
		{
			name:           "no build info and no resource config",
			cfg:            Config{},
			wantAttributes: map[string]interface{}{"service.name": "", "service.version": ""},
		},
		{
			name:           "build info and no resource config",
			cfg:            Config{},
			buildInfo:      component.BuildInfo{Command: "otelcoltest", Version: "0.0.0-test"},
			wantAttributes: map[string]interface{}{"service.name": "otelcoltest", "service.version": "0.0.0-test"},
		},
		{
			name:           "no build info and resource config",
			cfg:            Config{Resource: map[string]*string{"service.name": ptr("resource.name"), "service.version": ptr("resource.version"), "test": ptr("test")}},
			wantAttributes: map[string]interface{}{"service.name": "resource.name", "service.version": "resource.version", "test": "test"},
		},
		{
			name:           "build info and resource config",
			buildInfo:      component.BuildInfo{Command: "otelcoltest", Version: "0.0.0-test"},
			cfg:            Config{Resource: map[string]*string{"service.name": ptr("resource.name"), "service.version": ptr("resource.version"), "test": ptr("test")}},
			wantAttributes: map[string]interface{}{"service.name": "resource.name", "service.version": "resource.version", "test": "test"},
		},
		{
			name:           "deleting a nil value",
			buildInfo:      component.BuildInfo{Command: "otelcoltest", Version: "0.0.0-test"},
			cfg:            Config{Resource: map[string]*string{"service.name": nil, "service.version": ptr("resource.version"), "test": ptr("test")}},
			wantAttributes: map[string]interface{}{"service.version": "resource.version", "test": "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrs := attributes(internal.Settings{BuildInfo: tt.buildInfo}, tt.cfg)
			require.Equal(t, tt.wantAttributes, attrs)
		})
	}
}

func TestNewTracerProvider(t *testing.T) {
	tests := []struct {
		name               string
		wantTracerProvider any
		noopTracer         bool
	}{
		{
			name:               "noop tracer provider",
			noopTracer:         true,
			wantTracerProvider: noop.TracerProvider{},
		},
		{
			name:               "tracer provider",
			wantTracerProvider: &sdktrace.TracerProvider{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			previousValue := globalgates.NoopTracerProvider.IsEnabled()
			err := featuregate.GlobalRegistry().Set(globalgates.NoopTracerProvider.ID(), tt.noopTracer)
			require.NoError(t, err)
			defer func() {
				err = featuregate.GlobalRegistry().Set(globalgates.NoopTracerProvider.ID(), previousValue)
				require.NoError(t, err)
			}()
			provider, err := newTracerProvider(context.TODO(), internal.Settings{}, Config{})
			require.NoError(t, err)
			require.IsType(t, tt.wantTracerProvider, provider)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
