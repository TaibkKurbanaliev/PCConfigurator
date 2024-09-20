package storage

import "go.mongodb.org/mongo-driver/v2/mongo"

type Repository struct {
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{}
}
