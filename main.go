package main

import (
	"os"

	"github.com/anonurl/anonurl-api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())

    router.POST("/api/create", func(c *gin.Context) { services.Create(c) })
    router.POST("/api/report", func(c *gin.Context) { services.Report(c) })
    router.GET("/api/redirect/:id", func(c *gin.Context) { services.Redirect(c) })
    router.GET("/api/track/:id", func(c *gin.Context) { services.Track(c) })

    router.Run(os.Getenv("PORT"))
}
