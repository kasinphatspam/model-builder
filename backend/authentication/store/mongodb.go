package store

import (
	"context"
	"log"
	"mtrain-main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore(dsn string) *MongoDBStore {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	collection := client.Database("authentication").Collection("accounts")

	return &MongoDBStore{Collection: collection}
}

func (s *MongoDBStore) Save(account models.Account) error {
	_, err := s.Collection.InsertOne(context.Background(), &account)
	return err
}

func (s *MongoDBStore) Find(data map[string]interface{}) ([]models.Account, error) {
	// convert map[string]inteface{} to bson format
	bsonData, err := bson.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// query data in collection
	cursor, err := s.Collection.Find(context.TODO(), bsonData)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// decode mongo cursor to type object
	var accounts []models.Account
	for cursor.Next(context.TODO()) {
		var result models.Account
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, result)
	}
	return accounts, nil
}
