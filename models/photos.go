package models

type Photo struct {
    Typename     string `json:"__typename"`
    ID           int    `json:"id"`
    PhotoUri     string `json:"photoUri"`
    Ranking      int    `json:"ranking"`
    ThumbnailUri string `json:"thumbnailUri"`
}

type Data struct {
    Photos []Photo `json:"photos"`
    Prefix string  `json:"prefix"`
}

type ResponsePhoto struct {
    Data    Data   `json:"data"`
    Message string `json:"message"`
    Status  bool   `json:"status"`
}
