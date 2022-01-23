package database

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TrackURL(c *gin.Context, urlID string) {
    var response bson.M

    err := urls.FindOne(ctx, bson.M{"id": urlID}).Decode(&response)
    if err != nil {
        c.JSON(400, gin.H {
            "error": "URL/ID not founded on database",
        })

        return
    }

    c.JSON(200, response)
}
