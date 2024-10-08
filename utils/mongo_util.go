// mongo_utils.go
package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient interface for better testability
type MongoDBClient interface {
	InsertOne(collection string, doc interface{}) error
	FindByFilter(collection string, filter interface{}, limit int) ([]bson.M, error)
}

// MongoDBUtil is the implementation of MongoDBClient
type MongoDBUtil struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewMongoDBUtil creates a new MongoDBUtil instance
func NewMongoDBUtil(client *mongo.Client, dbName string) *MongoDBUtil {
	return &MongoDBUtil{
		client: client,
		db:     client.Database(dbName),
	}
}

// InsertOne inserts a new document into the collection
func (m *MongoDBUtil) InsertOne(collection string, doc interface{}) error {

	_, err := m.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}
	fmt.Println("Document inserted successfully!")
	return nil
}

// FindByFilter retrieves documents by filter with a limit
func (m *MongoDBUtil) FindByFilter(collection string, filter interface{}, limit int) ([]bson.M, error) {
	options := options.Find()
	options.SetLimit(int64(limit))

	cursor, err := m.db.Collection(collection).Find(context.TODO(), filter, options)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve documents: %w", err)
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("failed to decode document: %w", err)
		}
		results = append(results, result)
	}

	return results, nil
}
