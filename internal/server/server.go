package server

import (
	"intergation_service/internal/config"
	"intergation_service/internal/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func NewFiberServer(log *zap.Logger, mongoClient *mongo.Client, cfg *config.Config) *fiber.App {
	app := fiber.New()

	handler := handler.NewHandler(log, mongoClient)
	app.Post("/api/proxy/:requestId", handler.ProxyRequest)

	return app
}
