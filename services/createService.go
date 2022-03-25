package services

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

func Create(c *gin.Context) {
    var createBody CreateBody
    
    err := c.BindJSON(&createBody)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "Invalid request body.",
        })

        return
    }

    id := generateID()
    db.Urls.InsertOne(db.Ctx, URL {
        URL: createBody.URL,
        ID: id,
        Create: time.Now().Format("01/02/2006 15:04"),
    })
    
    c.JSON(200, gin.H {
        "id": id,
    }) 
}
