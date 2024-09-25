package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"pcbuilder/internal/domain"
	"reflect"
	"slices"
	"strings"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PCItemMongo struct {
	db *mongo.Database
}

func NewPCItemMongo(db *mongo.Database) *PCItemMongo {
	return &PCItemMongo{db: db}
}

func (r *PCItemMongo) Create(ctx context.Context, item interface{}) (string, error) {
	collectionName := strings.ToLower(reflect.TypeOf(item).Name())
	if !slices.Contains(itemCollectionsNames, collectionName) {
		return "", domain.ErrCollectioNotFound
	}

	result, err := r.db.Collection(collectionName).InsertOne(ctx, item)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(bson.ObjectID).Hex(), nil
}
