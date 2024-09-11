// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extension

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
)

type nopExtension struct {
	component.StartFunc
	component.ShutdownFunc
	Settings
}

func TestNewFactory(t *testing.T) {
	var testType = component.MustNewType("test")
	defaultCfg := struct{}{}
	nopExtensionInstance := new(nopExtension)

	factory := NewFactory(
		testType,
		func() component.Config { return &defaultCfg },
		func(context.Context, Settings, component.Config) (Extension, error) {
			return nopExtensionInstance, nil
		},
		component.StabilityLevelDevelopment)
	assert.EqualValues(t, testType, factory.Type())
	assert.EqualValues(t, &defaultCfg, factory.CreateDefaultConfig())

	assert.Equal(t, component.StabilityLevelDevelopment, factory.ExtensionStability())
	ext, err := factory.CreateExtension(context.Background(), Settings{}, &defaultCfg)
	assert.NoError(t, err)
	assert.Same(t, nopExtensionInstance, ext)
}

func TestMakeFactoryMap(t *testing.T) {
	type testCase struct {
		name string
		in   []Factory
		out  map[component.Type]Factory
	}

	p1 := NewFactory(component.MustNewType("p1"), nil, nil, component.StabilityLevelAlpha)
	p2 := NewFactory(component.MustNewType("p2"), nil, nil, component.StabilityLevelAlpha)
	testCases := []testCase{
		{
			name: "different names",
			in:   []Factory{p1, p2},
			out: map[component.Type]Factory{
				p1.Type(): p1,
				p2.Type(): p2,
			},
		},
		{
			name: "same name",
			in:   []Factory{p1, p2, NewFactory(component.MustNewType("p1"), nil, nil, component.StabilityLevelAlpha)},
		},
	}
	for i := range testCases {
		tt := testCases[i]
		t.Run(tt.name, func(t *testing.T) {
			out, err := MakeFactoryMap(tt.in...)
			if tt.out == nil {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestBuilder(t *testing.T) {
	var testType = component.MustNewType("test")
	defaultCfg := struct{}{}
	testID := component.NewID(testType)
	unknownID := component.MustNewID("unknown")

	factories, err := MakeFactoryMap([]Factory{
		NewFactory(
			testType,
			func() component.Config { return &defaultCfg },
			func(_ context.Context, settings Settings, _ component.Config) (Extension, error) {
				return nopExtension{Settings: settings}, nil
			},
			component.StabilityLevelDevelopment),
	}...)
	require.NoError(t, err)

	cfgs := map[component.ID]component.Config{testID: defaultCfg, unknownID: defaultCfg}
	b := NewBuilder(cfgs, factories)

	testIDSettings := createSettings(testID)
	testIDModuleInfo := ModuleInfo{
		Extension: map[component.Type]string{
			testType: "go.opentelemetry.io/collector/extension/extensiontest v1.2.3",
		},
	}
	testIDSettings.ModuleInfo = testIDModuleInfo

	e, err := b.Create(context.Background(), testIDSettings)
	assert.NoError(t, err)
	assert.NotNil(t, e)

	// Check that the extension has access to the resource attributes.
	nop, ok := e.(nopExtension)
	assert.True(t, ok)
	assert.Equal(t, nop.Settings.Resource.Attributes().Len(), 0)

	// Check that the extension has access to the module info.
	assert.Equal(t, testIDModuleInfo, nop.ModuleInfo)

	missingType, err := b.Create(context.Background(), createSettings(unknownID))
	assert.EqualError(t, err, "extension factory not available for: \"unknown\"")
	assert.Nil(t, missingType)

	missingCfg, err := b.Create(context.Background(), createSettings(component.NewIDWithName(testType, "foo")))
	assert.EqualError(t, err, "extension \"test/foo\" is not configured")
	assert.Nil(t, missingCfg)
}

func TestBuilderFactory(t *testing.T) {
	factories, err := MakeFactoryMap([]Factory{NewFactory(component.MustNewType("foo"), nil, nil, component.StabilityLevelDevelopment)}...)
	require.NoError(t, err)

	cfgs := map[component.ID]component.Config{component.MustNewID("foo"): struct{}{}}
	b := NewBuilder(cfgs, factories)

	assert.NotNil(t, b.Factory(component.MustNewID("foo").Type()))
	assert.Nil(t, b.Factory(component.MustNewID("bar").Type()))
}

func createSettings(id component.ID) Settings {
	return Settings{
		ID:                id,
		TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		BuildInfo:         component.NewDefaultBuildInfo(),
	}
}
