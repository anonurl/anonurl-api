package services

import (
	"github.com/anonurl/anonurl-api/database"
	"github.com/gin-gonic/gin"
)

type CreateBody struct {
    URL string
}

func Create(c *gin.Context) {
    var createBody CreateBody

    err := c.BindJSON(&createBody)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "Invalid request body.",
        })

        return
    }

    database.CreateURL(c, createBody.URL)
}
