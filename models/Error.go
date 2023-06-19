package models

type ErrorResponse struct {
    Error struct {
        Message string `json:"message"`
        Code    string `json:"code"`
    } `json:"error"`
    Code int `json:"code"`
}