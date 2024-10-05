package storage

import (
	"context"
	"pcbuilder/internal/domain"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) (string, error)
	GetUser(ctx context.Context, username, password string) (domain.User, error)
}

type PCList interface {
}

type PCItem interface {
	Create(ctx context.Context, item interface{}) (string, error)
	GetByID(ctx context.Context, id string, itemType string) (interface{}, error)
}

type Repository struct {
	Authorization
	PCItem
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		PCItem:        NewPCItemMongo(db),
	}
}
