package database

import (
	"context"
        "os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URL struct {
    URL string `bson:"url"`
    ID string `bson:"id"`
    Create time.Time `bson:"create"`
}

type Report struct {
    Message string `bson:"message"`
}

var clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_URI"))
var client, _ = mongo.Connect(context.TODO(), clientOptions)
var ctx, _ = context.WithTimeout(context.Background(), 10 * time.Second)
var urls = client.Database("anonurl").Collection("urls")
var reports = client.Database("anonurl").Collection("reports")
