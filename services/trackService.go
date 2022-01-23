package services

import (
    ct "github.com/anonurl/anonurl-api/controllers"
    "github.com/gin-gonic/gin"
)

func Track(c *gin.Context) {
    urlID := c.Param("id")

    ct.TrackURL(c, urlID)
}
