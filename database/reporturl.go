package database

import "github.com/gin-gonic/gin"

func ReportURL(c *gin.Context, message string) {
    reportCreate := Report {
        message,
    }

    reports.InsertOne(ctx, reportCreate)

    c.JSON(200, gin.H {
        "message": "Thanks for contributing.",
    })
}
