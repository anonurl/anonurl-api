package services

type URL struct {
    URL string `bson:"url"`
    ID string `bson:"id"`
    Create string `bson:"create"`
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
