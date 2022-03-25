package database

import (
    "context"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_URI"))
    Ctx = context.TODO()
    client, _ = mongo.Connect(Ctx, clientOptions)
    Urls = client.Database("anonurl").Collection("urls")
    Reports = client.Database("anonurl").Collection("reports")
)
