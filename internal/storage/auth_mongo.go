package storage

import (
	"context"
	"errors"
	"fmt"
	"pcbuilder/internal/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthMongo struct {
	db *mongo.Collection
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db.Collection("auth_mongo")}
}

func (r *AuthMongo) CreateUser(ctx context.Context, user domain.User) (int, error) {
	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
	return 0, err
}

func (r *AuthMongo) GetUser(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User
	filter := bson.D{
		{"name", username},
		{"password", password},
	}

	if err := r.db.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, domain.ErrUserNotFound
		}

		return domain.User{}, err
	}

	return user, nil
}
