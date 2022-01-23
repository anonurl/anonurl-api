package database

import (
	"context"
        "os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_URI"))
    client, _ = mongo.Connect(context.TODO(), clientOptions)
    Ctx, _ = context.WithTimeout(context.Background(), 10 * time.Second)
    Urls = client.Database("anonurl").Collection("urls")
    Reports = client.Database("anonurl").Collection("reports")
)
