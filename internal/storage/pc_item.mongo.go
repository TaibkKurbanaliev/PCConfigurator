package storage

import (
	"context"
	"pcbuilder/internal/domain"
	"pcbuilder/internal/domain/items"
	"reflect"
	"slices"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"

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

func (r *PCItemMongo) GetByID(ctx context.Context, id string, itemType string) (interface{}, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.db.Collection(strings.ToLower(itemType)).FindOne(
		ctx,
		bson.D{
			{"_id", objID},
		})

	return DecodeByType(result, itemType)
}

func DecodeByType(result *mongo.SingleResult, itemType string) (interface{}, error) {
	switch strings.ToLower(itemType) {
	case "cpu":
		var item items.CPU
		err := result.Decode(&item)
		return item, err
	case "cpu_cooler":
		var item items.CPUCooler
		err := result.Decode(&item)
		return item, err
	case "case":
		var item items.Case
		err := result.Decode(&item)
		return item, err
	case "gpu":
		var item items.GPU
		err := result.Decode(&item)
		return item, err
	case "memory":
		var item items.Memory
		err := result.Decode(&item)
		return item, err
	case "motherboard":
		var item items.Motherboard
		err := result.Decode(&item)
		return item, err
	case "power_supply":
		var item items.PowerSupply
		err := result.Decode(&item)
		return item, err
	case "storage":
		var item items.Storage
		err := result.Decode(&item)
		return item, err
	}

	return nil, domain.ErrCollectioNotFound
}
