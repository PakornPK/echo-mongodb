package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name     string `json:"name" bson:"name"`
	LastName string `json:"last_name" bson:"last_name"`
}

func InsertUser(client *mongo.Client, data bson.D) error {
	db := client.Database("DB_PX")
	users := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := users.InsertOne(ctx, data)
	if err != nil {
		return errors.New("InsertUser: Can not Insert data")
	}
	return nil
}

func ReadUser(client *mongo.Client) (*User, error) {
	db := client.Database("DB_PX")
	users := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result := new(User)
	err := users.FindOne(ctx, bson.D{}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ReadAllUser(client *mongo.Client) ([]bson.D, error) {
	db := client.Database("DB_PX")
	coll := db.Collection("users")
	result, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []bson.D
	if err = result.All(context.TODO(), &users); err != nil {
    	panic(err)
	}
	if err := result.Close(context.TODO()); err != nil {
    	panic(err)
	}
	return users, nil
}
