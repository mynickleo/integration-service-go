package main

import (
	"context"
	"fmt"
	"intergation_service/internal/config"
	"intergation_service/internal/database/mongo"
	"intergation_service/internal/seed"
	"intergation_service/internal/server"
	"intergation_service/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	logg := logger.NewLogger()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	mongoClient, err := mongo.Connect(cfg.MongoURI)
	if err != nil {
		logg.Fatal("failed to connect to MongoDB", logger.ZapError(err))
	}
	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			logg.Error("failed to disconnect MongoDB", logger.ZapError(err))
		}
	}()

	err = seed.SeedDatabase(logg, mongoClient)
	if err != nil {
		logg.Fatal("failed to seed database", zap.Error(err))
	}

	app := server.NewFiberServer(logg, mongoClient, cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
			logg.Fatal("failed to start server", logger.ZapError(err))
		}
	}()

	logg.Info("server started", logger.ZapField("port", cfg.Port))

	<-quit

	logg.Info("server shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		logg.Fatal("failed to shutdown server", logger.ZapError(err))
	}
}
