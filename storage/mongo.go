package storage

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoStorage the mongo storage
type MongoStorage struct {
	client *mongo.Client
}

// NewMongoStorage MongoStorage's initializer
func NewMongoStorage(connectionString string) (MongoStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, connectionString)

	defer cancel()

	return MongoStorage{client: client}, err
}
