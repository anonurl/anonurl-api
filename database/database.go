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

var (
    clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_URI"))
    client, _ = mongo.Connect(context.TODO(), clientOptions)
    ctx, _ = context.WithTimeout(context.Background(), 10 * time.Second)
    urls = client.Database("anonurl").Collection("urls")
    reports = client.Database("anonurl").Collection("reports")
)
