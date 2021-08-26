package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func GetContext(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)
	return ctx
}

func init() {
	/*
	   Connect to my cluster
	*/
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://store_user:ODg7FC1IJUaMiSYB@cluster0.dmbnx.mongodb.net/store?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := GetContext(15)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database successfully !")
	}
}
