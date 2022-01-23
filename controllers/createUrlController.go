package controllers

import (
	"crypto/rand"
	"math/big"
	"time"
        "github.com/gin-gonic/gin"
        db "github.com/anonurl/anonurl-api/database"
)

func generateID() string {
    letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
    result := make([]byte, 6)

    for i := 0; i < 6; i++ {
        n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))

        result[i] = letters[n.Int64()]
    }

    return string(result)
}

func CreateURL(c *gin.Context, url string) {
    id := generateID()

    urlCreate := URL {
        url,
        id,
        time.Now(),
    }

    db.Urls.InsertOne(db.Ctx, urlCreate)
    
    c.JSON(200, gin.H {
        "url": url,
        "id": id,
        "create": time.Now(),
    }) 
}
