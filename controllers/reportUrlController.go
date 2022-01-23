package controllers

import (
    "github.com/gin-gonic/gin"
    db "github.com/anonurl/anonurl-api/database"
)

func ReportURL(c *gin.Context, message string) {
    reportCreate := Report {
        message,
    }

    db.Reports.InsertOne(db.Ctx, reportCreate)

    c.JSON(200, gin.H {
        "message": "Thanks for contributing.",
    })
}
