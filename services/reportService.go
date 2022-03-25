package services

import (
    db "github.com/anonurl/anonurl-api/database"
    "github.com/gin-gonic/gin"
)

func Report(c *gin.Context) {
    var reportBody ReportBody

    err := c.BindJSON(&reportBody)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "Invalid request body.",
        })

        return
    }

    db.Reports.InsertOne(db.Ctx, Reports {
        reportBody.Message,
    })

    c.Status(200)
}
