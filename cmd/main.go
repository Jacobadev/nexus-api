package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/server"
	"github.com/gateway-address/pkg/db/postgres"
	"github.com/gateway-address/pkg/db/redis"
	"github.com/gateway-address/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

var loggerConfig = log.New(os.Stderr, "zipkin-example", log.Ldate|log.Ltime|log.Llongfile)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer(url string) (func(context.Context) error, error) {
	exporter, err := zipkin.New(
		url,
		zipkin.WithLogger(loggerConfig),
	)
	if err != nil {
		return nil, err
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("zipkin-test"),
		)),
	)
	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil
}

func main() {
	url := flag.String("zipkin", "http://localhost:9411/api/v2/spans", "zipkin url")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initTracer(*url)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer func() {
		if err := shutdown(ctx); err != nil {
			fmt.Print(err.Error())
		}
	}()

	// tr := otel.Tracer("Main-function")
	// ctx, span := tr.Start(ctx, "foo", trace.WithSpanKind(trace.SpanKindServer))
	// <-time.After(6 * time.Millisecond)
	// <-time.After(6 * time.Millisecond)
	// span.End()
	cfgFile, err := config.LoadConfig("./config/config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	// Initialize logger
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	// Initialize database connection
	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	appLogger.Infof("Postgres connected")
	redis := redis.NewRedisClient(cfg)
	defer redis.Close()
	appLogger.Infof("Redis connected")

	// Create and run the server
	s := server.NewServer(cfg, db, appLogger, redis)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
