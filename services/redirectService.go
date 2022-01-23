package services

import (
    ct "github.com/anonurl/anonurl-api/controllers"
    "github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
    urlID := c.Param("id")

    ct.RedirectURL(c, urlID) 
}
