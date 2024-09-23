package storage

import (
	"context"
	"pcbuilder/internal/domain"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	GetUser(ctx context.Context, username, password string) (domain.User, error)
}

type PCList interface {
}

type PCItem interface {
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}
