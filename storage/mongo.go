package storage

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type MongoStorage struct {
	client *mongo.Client
}

func NewMongoStorage(connectionString string) (Storage, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, connectionString)

	return MongoStorage{client: client}, err
}
