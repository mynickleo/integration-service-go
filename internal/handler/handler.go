package handler

import (
	"intergation_service/internal/repository"
	"intergation_service/internal/service"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Handler struct {
	log          *zap.Logger
	proxyService *service.ProxyService
}

func NewHandler(log *zap.Logger, mongoClient *mongo.Client) *Handler {
	dataBase := mongoClient.Database("integration_db")
	repo := repository.NewRequestRepository(dataBase)
	proxyService := service.NewProxyService(repo)

	return &Handler{
		log:          log,
		proxyService: proxyService,
	}
}

func (handler *Handler) ProxyRequest(ctx *fiber.Ctx) error {
	requestId := ctx.Params("requestId")

	resp, err := handler.proxyService.Proxy(ctx.Context(), requestId)
	if err != nil {
		handler.log.Error("failed to proxy request", zap.Error(err))
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		handler.log.Error("failed to read response body", zap.Error(err))
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(resp.StatusCode).Send(body)
}
