package services

import "time"

type URLs struct {
    URL string `bson:"url"`
    ID string `bson:"id"`
    Create time.Time `bson:"create"`
}

type Reports struct {
    Message string `bson:"message"`
}

type ReportBody struct {
    Message string
}

type CreateBody struct {
    URL string
}
