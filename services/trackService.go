package services

import (
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    db "github.com/anonurl/anonurl-api/database"
)

func Track(c *gin.Context) {
    var response bson.M

    err := db.Urls.FindOne(db.Ctx, bson.M{"id": c.Param("id")}).Decode(&response)
    if err != nil {
        c.JSON(400, gin.H {
            "error": "URL/ID not founded on database",
        })

        return
    }

    c.JSON(200, response)
}
