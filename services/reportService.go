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

    reportCreate := Reports {
        reportBody.Message,
    }

    db.Reports.InsertOne(db.Ctx, reportCreate)

    c.JSON(200, gin.H {
        "message": "Thanks for contributing.",
    })
}
