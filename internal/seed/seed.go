package seed

import (
	"context"
	"intergation_service/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func SeedDatabase(log *zap.Logger, client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dataBase := client.Database("integration_db")
	collection := dataBase.Collection("proxy_requests")

	count, err := collection.CountDocuments(ctx, struct{}{})
	if err != nil {
		log.Error("failed to count documents", zap.Error(err))
		return err
	}

	if count > 0 {
		log.Info("seed skipped: documents already exist")
		return nil
	}

	testRequest := model.ProxyRequest{
		ID:          primitive.NewObjectID(),
		Method:      "GET",
		URL:         "https://jsonplaceholder.typicode.com/posts/1",
		Headers:     map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
		Body:        nil,
		QueryParams: map[string]string{},
	}

	_, err = collection.InsertOne(ctx, testRequest)
	if err != nil {
		log.Error("failed to insert test document", zap.Error(err))
		return err
	}

	log.Info("seed completed: test document inserted")
	return nil
}
