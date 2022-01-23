package controllers

import "time"

type URL struct {
    URL string `bson:"url"`
    ID string `bson:"id"`
    Create time.Time `bson:"create"`
}

type Report struct {
    Message string `bson:"message"`
}
