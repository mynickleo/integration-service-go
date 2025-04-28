package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProxyRequest struct {
	ID             primitive.ObjectID     `bson:"_id,omitempty"`
	Method         string                 `bson:"method"`
	URL            string                 `bson:"url"`
	Headers        map[string]string      `bson:"headers,omitempty"`
	QueryParams    map[string]string      `bson:"query_params,omiteempty"`
	Body           map[string]interface{} `bson:"body,omitempty"`
	AuthToken      string                 `bson:"auth_token,omitempty"`
	TimeoutSeconds int                    `bson:"timeout_seconds"`
}
