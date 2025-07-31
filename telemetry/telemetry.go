package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var tp *sdktrace.TracerProvider
var tExp *otlptrace.Exporter

func GetInstance() *sdktrace.TracerProvider {
	return tp
}

func NewTracerProvider(grpcHost string, serviceName, namespace string) error {
	var err error
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientConn, err := grpc.DialContext(timeoutCtx, grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithGRPCConn(clientConn),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)

	tExp, err = otlptrace.New(ctx, traceClient)
	if err != nil {
		return err
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.K8SNamespaceName(namespace),
		),
	)
	if err != nil {
		return err
	}
	bsp := sdktrace.NewBatchSpanProcessor(tExp)
	tp = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tp)
	return nil
}
