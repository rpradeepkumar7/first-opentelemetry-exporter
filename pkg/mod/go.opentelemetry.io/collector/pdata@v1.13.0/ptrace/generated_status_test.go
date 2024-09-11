// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
)

func TestStatus_MoveTo(t *testing.T) {
	ms := generateTestStatus()
	dest := NewStatus()
	ms.MoveTo(dest)
	assert.Equal(t, NewStatus(), ms)
	assert.Equal(t, generateTestStatus(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newStatus(&otlptrace.Status{}, &sharedState)) })
	assert.Panics(t, func() { newStatus(&otlptrace.Status{}, &sharedState).MoveTo(dest) })
}

func TestStatus_CopyTo(t *testing.T) {
	ms := NewStatus()
	orig := NewStatus()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestStatus()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newStatus(&otlptrace.Status{}, &sharedState)) })
}

func TestStatus_Code(t *testing.T) {
	ms := NewStatus()
	assert.Equal(t, StatusCode(0), ms.Code())
	testValCode := StatusCode(1)
	ms.SetCode(testValCode)
	assert.Equal(t, testValCode, ms.Code())
}

func TestStatus_Message(t *testing.T) {
	ms := NewStatus()
	assert.Equal(t, "", ms.Message())
	ms.SetMessage("cancelled")
	assert.Equal(t, "cancelled", ms.Message())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newStatus(&otlptrace.Status{}, &sharedState).SetMessage("cancelled") })
}

func generateTestStatus() Status {
	tv := NewStatus()
	fillTestStatus(tv)
	return tv
}

func fillTestStatus(tv Status) {
	tv.orig.Code = 1
	tv.orig.Message = "cancelled"
}
