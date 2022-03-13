package database

import (
	"context"
	"errors"
	"time"

	"github.com/echo_CRUD/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
)

var Client *mongo.Client
var err error
var ctx context.Context

func ConnectDB()  error {
	Client, err = mongo.NewClient(options.Client().ApplyURI(config.URI))
	if err != nil {
		return errors.New("ConnectDB: database not availabel")
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		return errors.New("ConnectDB: can not connect database")
	}

	return nil
}
