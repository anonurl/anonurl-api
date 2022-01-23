package services

import (
    "github.com/anonurl/anonurl-api/database"
    "github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
    urlID := c.Param("id")

    database.RedirectURL(c, urlID) 
}
