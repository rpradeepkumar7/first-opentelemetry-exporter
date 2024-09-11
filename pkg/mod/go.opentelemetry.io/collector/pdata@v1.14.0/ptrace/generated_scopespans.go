// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ScopeSpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewScopeSpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeSpans struct {
	orig  *otlptrace.ScopeSpans
	state *internal.State
}

func newScopeSpans(orig *otlptrace.ScopeSpans, state *internal.State) ScopeSpans {
	return ScopeSpans{orig: orig, state: state}
}

// NewScopeSpans creates a new empty ScopeSpans.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewScopeSpans() ScopeSpans {
	state := internal.StateMutable
	return newScopeSpans(&otlptrace.ScopeSpans{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ScopeSpans) MoveTo(dest ScopeSpans) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlptrace.ScopeSpans{}
}

// Scope returns the scope associated with this ScopeSpans.
func (ms ScopeSpans) Scope() pcommon.InstrumentationScope {
	return pcommon.InstrumentationScope(internal.NewInstrumentationScope(&ms.orig.Scope, ms.state))
}

// SchemaUrl returns the schemaurl associated with this ScopeSpans.
func (ms ScopeSpans) SchemaUrl() string {
	return ms.orig.SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ScopeSpans.
func (ms ScopeSpans) SetSchemaUrl(v string) {
	ms.state.AssertMutable()
	ms.orig.SchemaUrl = v
}

// Spans returns the Spans associated with this ScopeSpans.
func (ms ScopeSpans) Spans() SpanSlice {
	return newSpanSlice(&ms.orig.Spans, ms.state)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ScopeSpans) CopyTo(dest ScopeSpans) {
	dest.state.AssertMutable()
	ms.Scope().CopyTo(dest.Scope())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.Spans().CopyTo(dest.Spans())
}
