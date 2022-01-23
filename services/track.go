package services

import (
    "github.com/anonurl/anonurl-api/database"
    "github.com/gin-gonic/gin"
)

func Track(c *gin.Context) {
    urlID := c.Param("id")

    database.TrackURL(c, urlID)
}
