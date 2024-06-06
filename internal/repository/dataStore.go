package repository

import (
	"context"
	"fmt"
	"github.com/mhsbz/xiaohan/internal/schemas"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	db *mongo.Database
}

func NewMongoClient() *MongoClient {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	db := client.Database("xiaohan")
	return &MongoClient{
		db,
	}
}

func (d *MongoClient) CreateUser(user *schemas.User) error {
	results, err := d.db.Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	fmt.Println(results.InsertedID)
	return nil
}

func (d *MongoClient) GetUser(memberID string) (user *schemas.User, err error) {
	err = d.db.Collection("user").FindOne(context.Background(), map[string]string{"member_id": memberID}).Decode(&user)
	d.db.Collection("user").FindOne(context.Background(), map[string]string{"member_id": memberID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
