// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statustest // import "go.opentelemetry.io/collector/service/internal/status/statustest"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service/internal/status"
)

func NewNopStatusReporter() status.Reporter {
	return &nopStatusReporter{}
}

type nopStatusReporter struct{}

func (r *nopStatusReporter) Ready() {}

func (r *nopStatusReporter) ReportStatus(*component.InstanceID, *component.StatusEvent) {}

func (r *nopStatusReporter) ReportOKIfStarting(*component.InstanceID) {}
