package database

import (
    "context"
    "time"

    "github.com/gin-gonic/gin"
)

func ReportURL(c *gin.Context, message string) {
    ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
    reports := client.Database("anonurl").Collection("reports")

    reportCreate := Report {
        message,
    }

    reports.InsertOne(ctx, reportCreate)

    c.JSON(200, gin.H {
        "message": "Thanks for contributing.",
    })
}
