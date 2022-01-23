package services

import (
    ct "github.com/anonurl/anonurl-api/controllers"
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

    ct.CreateURL(c, createBody.URL)
}
