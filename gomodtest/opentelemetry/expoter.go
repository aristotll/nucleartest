package main

import (
    "io"

    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/trace"
)

type Exporter struct {
    spanExp trace.SpanExporter
}

func (e *Exporter) InitWithJaeger(jaegerURL string) error {
    // "http://localhost:14268/api/traces"
    endpoint := jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL))
    exp, err := jaeger.New(endpoint)
    if err != nil {
       return err
    }
    e.spanExp = exp
    return nil
}

func (e *Exporter) InitWithStdout(w io.Writer) error {
    exp, err := stdouttrace.New(
       stdouttrace.WithWriter(w),
       // Use human readable output.
       stdouttrace.WithPrettyPrint(),
       // Do not print timestamps for the demo.
       stdouttrace.WithoutTimestamps(),
    )
    if err != nil {
       return err
    }
    e.spanExp = exp
    return nil
}