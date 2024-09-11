wsl shutdown
wsl -l -v (Check for the Distribution) & (Set the exact default Ubuntu)
wsl -d Ubuntu-22.04
wsl
mkdir first-opentelemetry-exporter
ls
cd first-opentelemetry-exporter
mkdir .devcontainer
mkdir .vscode

cd .devcontainer
nano devcontainer.json

#Code Snippet
{
  "name": "Go",
  "image": "mcr.microsoft.com/devcontainers/go:0-1-bullseye",
  "forwardPorts": [4317,4318],
  "postCreateCommand": "bash .devcontainers/postCreateCommand.sh"
}

nano postCreateCommand.sh
#Code Snippet
go install go.opentelemetry.io/collector/cmd/builder@latest
go install github.com/go-delve/delve/cmd/dlv@latest
Or
#Code Snippet
go install go.opentelemetry.io/collector/cmd/builder@v0.107.0
go install github.com/go-delve/delve/cmd/dlv@latest



cd .vscode
nano launch.json
#Code Snippet
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Connect to server",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "port": 2345,
            "host": "127.0.0.1",
            "apiVersion": 2,
            "showLog": true
        }
    ]
}



CTRL + SHIFT + P
Dev Container : Reopen in Container (It will Create the Conatiner Env)
Command Palette (Terminal)
Check the Go - Builder Version
go version
builder version
cd ..
ls

sudo service docker status (Check Docker is Running)
sudo service docker start

nano otelcol-builder.yaml
#Code Snippet
dist:
  name: otelcol-custom
  description: Local OpenTelemetry Collector binary
  output_path: /tmp/dist
  otelcol_version: 0.108.0  # Ensure this matches the builder version
exporters:
  - gomod: go.opentelemetry.io/collector/exporter/debugexporter v0.108.0
receivers:
  - gomod: go.opentelemetry.io/collector/receiver/otlpreceiver v0.108.0

builder --config=otelcol-builder.yaml


nano config.yaml
#Code Snippet
receivers:
  otlp:
    protocols:
      http:
      grpc:
exporters:
  debug:
    verbosity: detailed
service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters:
        - emptyexporter
        - debug
    metrics:
      receivers: [otlp]
      exporters:
        - emptyexporter
        - debug
    logs:
      receivers: [otlp]
      exporters:
        - emptyexporter
        - debug



nano ~./bashrc
#Code Snippet
export GOPATH=/root/first-opentelemetry-exporter
export PATH=$PATH:$GOPATH/bin


source ~./bashrc
otelcol-custom binary in the /tmp/dist
/tmp/dist/otelcol-custom --config=config.yaml


Creating Exporter
mkdir emptyexporter
cd emptyexporter
nano go.mod
module emptyexporter


go 1.22


require (
    go.opentelemetry.io/collector v0.108.0
    go.opentelemetry.io/collector/component v0.108.0
    go.opentelemetry.io/collector/exporter/exporterhelper v0.108.0
    go.opentelemetry.io/collector/pdata v0.14.0
)


require (
    github.com/gogo/protobuf v1.3.2
    github.com/json-iterator/go v1.1.12
    github.com/stretchr/testify v1.9.0
    go.uber.org/goleak v1.3.0
    go.uber.org/multierr v1.11.0
    google.golang.org/grpc v1.65.0
    google.golang.org/protobuf v1.34.2
)


require (
    github.com/davecgh/go-spew v1.1.1 // indirect
    github.com/kr/pretty v0.3.1 // indirect
    github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
    github.com/modern-go/reflect2 v1.0.2 // indirect
    github.com/pmezard/go-difflib v1.0.0 // indirect
    github.com/rogpeppe/go-internal v1.10.0 // indirect
    golang.org/x/net v0.25.0 // indirect
    golang.org/x/sys v0.20.0 // indirect
    golang.org/x/text v0.15.0 // indirect
    google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
    gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
    gopkg.in/yaml.v3 v3.0.1 // indirect
)


retract (
    v1.0.0-rc10 // RC version scheme discovered to be alphabetical, use v1.0.0-rcv0011 instead
    v0.57.1 // Release failed, use v0.57.2
    v0.57.0 // Release failed, use v0.57.2
)


//replace go.opentelemetry.io/collector/config/configtelemetry => ../config/configtelemetry


replace go.opentelemetry.io/collector/pdata => ../pdata


//replace go.opentelemetry.io/collector/pdata/testdata => ../pdata/testdata


//replace go.opentelemetry.io/collector/pdata/pprofile => ../pdata/pprofile


//replace go.opentelemetry.io/collector => github.com/droosma/emptyexporter v0.0.1



nano factory.go
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



Cd ..
nano otelcol-builder.yaml
dist:
  name: otelcol-custom
  description: Local OpenTelemetry Collector binary
  output_path: /tmp/dist
  otelcol_version: 0.108.0  # Ensure this matches the builder version
exporters:
  - gomod: "github.com/droosma/emptyexporter v0.0.1"
    path: emptyexporter



nano config.yaml
receivers:
  otlp:
    protocols:
      http:
      grpc:
exporters:
  debug:
    verbosity: detailed
service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters:
        - emptyexporter
        - debug
    metrics:
      receivers: [otlp]
      exporters:
        - emptyexporter
        - debug
    logs:
      receivers: [otlp]
      exporters:
        - emptyexporter
        - debug



/tmp/dist/otelcol-custom --config=config.yaml

Debug the Exporter
Nano otelcol-builder.yaml
dist:
  debug_compilation: true



dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient --log exec /tmp/dist/otelcol-custom -- --config=config.yaml
Check the .vscode launch code
nano launch.json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Connect to server",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "port": 2345,
            "host": "127.0.0.1",
            "apiVersion": 2,
            "showLog": true
        }
    ]
}
