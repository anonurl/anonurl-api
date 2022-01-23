package services

import (
    "github.com/anonurl/anonurl-api/database"
    "github.com/gin-gonic/gin"
)

type ReportBody struct {
    Message string
}

type Reports struct {
    Message string `json:"message"`
}

func Report(c *gin.Context) {
    var reportBody ReportBody

    err := c.BindJSON(&reportBody)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "Invalid request body.",
        })

        return
    }

    database.ReportURL(c, reportBody.Message)
}
