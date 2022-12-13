package models

type User struct {
	ID           int64 `db:"id, primarykey, autoincrement" json:"id"`
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Username    string `json:"username" binding:"required,alphanum"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	Identification   string `json:"identification" binding:"required"`
	Sinac_registry string `json:"sinac_registry""`
	Token         string   `json:"token"`
	Refresh_token string   `json:"refresh_token"`
}

type AddUserBody struct{
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Username    string `json:"username" binding:"required,alphanum"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	Identification   string `json:"identification" binding:"required"`
	Sinac_registry string `json:"sinac_registry""`
	Token         string   `json:"token"`
	Refresh_token string   `json:"refresh_token"`
}

type UserResponse struct{
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Username    string `json:"username" binding:"required,alphanum"`
	Email        string `json:"email" binding:"required,email"`
	Identification   string `json:"identification" binding:"required"`
	Sinac_registry string `json:"sinac_registry""`
	Token         string   `json:"token"`
	Refresh_token string   `json:"refresh_token"`
}

type LoginUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password     string `json:"password" binding:"required,min=6"`
}