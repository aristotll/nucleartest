package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
)

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("fib"),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

func Init() (*trace.TracerProvider, error) {
	jaegerURL := "http://localhost:14268a/api/traces"
	//exp := new(Exporter)
	//if err := exp.InitWithJaeger("http://localhost:14268a/api/traces"); err != nil {
	// return nil, err
	//}

	//endpoint := jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL))
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL)))
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
		trace.WithSampler(trace.AlwaysSample()),
	)
	//otel.SetTracerProvider(tp)
	return tp, nil
}

// 0,1,1,2,3,5,8,13,24,36,60,96,
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	x, y := 0, 1
	for i := 2; i < n; i++ {
		x, y = y, x+y
	}
	return x + y
}

func main() {
	tp, err := Init()
	if err != nil {
		panic(err)
	}

	otel.SetTracerProvider(tp)

	if len(os.Args) != 2 {
		fmt.Println("Usage: main [count to the number], e.g main 5")
		return
	}
	n := os.Args[1]
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	ctx := context.Background()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		// ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		// defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	tr := tp.Tracer("hello-otel")
	ctx, span := tr.Start(ctx, "calculate fibonacci")
	defer span.End()

	span.AddEvent("start calculate")
	span.SetAttributes(attribute.Key("n").String(n))

	_ = n
	v := fibonacci(100)
	fmt.Println(v)
	span.AddEvent("calculate completed")
	span.SetAttributes(attribute.Key("result").Int(v))
}
