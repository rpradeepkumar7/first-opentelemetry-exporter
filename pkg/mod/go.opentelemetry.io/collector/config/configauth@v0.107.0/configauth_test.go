// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configauth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/auth"
)

var mockID = component.MustNewID("mock")

func TestNewDefaultAuthentication(t *testing.T) {
	auth := NewDefaultAuthentication()
	assert.NotNil(t, auth)
	assert.Empty(t, auth)
}

func TestGetServer(t *testing.T) {
	testCases := []struct {
		desc          string
		authenticator extension.Extension
		expected      error
	}{
		{
			desc:          "obtain server authenticator",
			authenticator: auth.NewServer(),
			expected:      nil,
		},
		{
			desc:          "not a server authenticator",
			authenticator: auth.NewClient(),
			expected:      errNotServer,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// prepare
			cfg := &Authentication{
				AuthenticatorID: mockID,
			}
			ext := map[component.ID]component.Component{
				mockID: tC.authenticator,
			}

			authenticator, err := cfg.GetServerAuthenticator(context.Background(), ext)

			// verify
			if tC.expected != nil {
				assert.ErrorIs(t, err, tC.expected)
				assert.Nil(t, authenticator)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, authenticator)
			}
		})
	}
}

func TestGetServerFails(t *testing.T) {
	cfg := &Authentication{
		AuthenticatorID: component.MustNewID("does_not_exist"),
	}

	authenticator, err := cfg.GetServerAuthenticator(context.Background(), map[component.ID]component.Component{})
	assert.ErrorIs(t, err, errAuthenticatorNotFound)
	assert.Nil(t, authenticator)
}

func TestGetClient(t *testing.T) {
	testCases := []struct {
		desc          string
		authenticator extension.Extension
		expected      error
	}{
		{
			desc:          "obtain client authenticator",
			authenticator: auth.NewClient(),
			expected:      nil,
		},
		{
			desc:          "not a client authenticator",
			authenticator: auth.NewServer(),
			expected:      errNotClient,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// prepare
			cfg := &Authentication{
				AuthenticatorID: mockID,
			}
			ext := map[component.ID]component.Component{
				mockID: tC.authenticator,
			}

			authenticator, err := cfg.GetClientAuthenticator(context.Background(), ext)

			// verify
			if tC.expected != nil {
				assert.ErrorIs(t, err, tC.expected)
				assert.Nil(t, authenticator)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, authenticator)
			}
		})
	}
}

func TestGetClientFails(t *testing.T) {
	cfg := &Authentication{
		AuthenticatorID: component.MustNewID("does_not_exist"),
	}
	authenticator, err := cfg.GetClientAuthenticator(context.Background(), map[component.ID]component.Component{})
	assert.ErrorIs(t, err, errAuthenticatorNotFound)
	assert.Nil(t, authenticator)
}
