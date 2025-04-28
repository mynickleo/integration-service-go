package repository

import (
	"context"
	"intergation_service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestRepository struct {
	collection *mongo.Collection
}

func NewRequestRepository(db *mongo.Database) *RequestRepository {
	return &RequestRepository{
		collection: db.Collection("proxy_requests"),
	}
}

func (r *RequestRepository) GetRequestByID(ctx context.Context, id string) (*model.ProxyRequest, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result model.ProxyRequest
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
