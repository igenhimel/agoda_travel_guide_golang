package models

type Description struct {
    Typename     string `json:"__typename"`
    Description  string `json:"description"`
    LanguageCode string `json:"languageCode"`
}

type Response struct {
    Data    []Description `json:"data"`
    Message string       `json:"message"`
    Status  bool         `json:"status"`
}