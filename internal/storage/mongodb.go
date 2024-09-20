package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewMongoClient(ctx *context.Context, configDB ConfigDB) (*mongo.Client, error) {
	return mongo.Connect(options.Client().ApplyURI(configDB.Host + configDB.Port))
}
