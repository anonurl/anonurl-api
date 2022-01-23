package database

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TrackURL(c *gin.Context, urlID string) {
    var response bson.M

    ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
    urls := client.Database("anonurl").Collection("urls")

    err := urls.FindOne(ctx, bson.M{"id": urlID}).Decode(&response)
    if err != nil {
        c.JSON(400, gin.H {
            "error": "URL/ID not founded on database",
        })

        return
    }

    c.JSON(200, response)
}
