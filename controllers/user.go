package controllers

import (
	// "context"
	"fmt"
	"log"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"github.com/lib/pq"

	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"

)

type UserController struct{}

type AddUserBody struct{
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Username    string `json:"username" binding:"required,alphanum"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	Identification   string `json:"identification" binding:"required"`
	Sinac_registry string `json:"sinac_registry" binding:"required"`
	Token         string   `json:"token"`
	Refresh_token string   `json:"refresh_token"`
}

type UserResponse struct{
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Username    string `json:"username" binding:"required,alphanum"`
	Email        string `json:"email" binding:"required,email"`
	Identification   string `json:"identification" binding:"required"`
	Sinac_registry string `json:"sinac_registry" binding:"required"`
	Token         string   `json:"token"`
	Refresh_token string   `json:"refresh_token"`
}

type LoginUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password     string `json:"password" binding:"required,min=6"`
}
	
// HashPassword return the bcrypt hash of the password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) (bool,string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("Password is incorrect")
		check = false
	}

	return check,msg
}

func GetUser(context *gin.Context, username string ) (models.User, error, string) {
	var user models.User
	msg := ""
	err := database.DB.Raw("SELECT * FROM public.AFN_GetUser(?) r(ID BIGINT, first_name varchar(50), last_name varchar(50), username varchar(100), email varchar(200), password varchar(500), identification varchar(20), sinac_registry varchar(20), token varchar(500), refresh_token varchar(500));", 
	       username).Scan(&user).Error

	if err != nil {
		msg = fmt.Sprintf("Username not found")
	}
	
	return user, err, msg
}

func (ctrl UserController) SignUpUser(ctx *gin.Context) {
	var user []*models.User

	body := AddUserBody{}
	err_body := ctx.BindJSON(&body)
	helpers.HandleErr(err_body)

	hashedPassword := HashPassword(body.Password)
	body.Password = hashedPassword

	token, refreshToken, _ := helpers.GenerateAllTokens(body.Username)
	body.Token = token
	body.Refresh_token = refreshToken

	err := database.DB.Raw("SELECT * FROM public.afn_adduser(?,?,?,?,?,?,?,?,?) AS (first_name varchar(50), last_name varchar(50), username varchar(100), email varchar(200), password varchar(500), identification varchar(20), sinac_registry varchar(20), token varchar(500), refresh_token varchar(500));",
	       body.First_name, body.Last_name, body.Username, body.Email, body.Password, body.Identification, body.Sinac_registry, body.Token, body.Refresh_token).Scan(&user).Error
	
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
				case "unique_violation":
					ctx.JSON(http.StatusForbidden, gin.H{"response": err})
					return
			}
		}
			ctx.JSON(http.StatusInternalServerError, gin.H{"response": err})
			return
	}

	rsp := UserResponse{
		First_name: body.First_name,
		Last_name: body.Last_name,
		Username :body.Username,
		Email : body.Email,
		Identification : body.Identification,
		Sinac_registry : body.Sinac_registry,
		Token : body.Token,
		Refresh_token : body.Refresh_token,
	}
	
	ctx.JSON(http.StatusOK, gin.H{"response": rsp}) 
}

//UpdateAllTokens renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, username string) {
	var user models.User

	err := database.DB.Raw("SELECT * FROM public.AFN_UpdateUser(?, ?, ?)r(first_name varchar(50), last_name varchar(50), username varchar(100), email varchar(200),identification varchar(20), sinac_registry varchar(20), token varchar(500), refresh_token varchar(500));", 
	username, signedToken, signedRefreshToken).Scan(&user).Error
	helpers.HandleErr(err)

	return
}

func (ctrl UserController) LoginUser(ctx *gin.Context) {

	body := LoginUserRequest{}
	err_body := ctx.BindJSON(&body)
	helpers.HandleErr(err_body)

	user, err, msg := GetUser(ctx, body.Username)
	if err != nil{
		if user.Username == ""{
			ctx.JSON(http.StatusNotFound, gin.H{"error": msg})
			return	
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return	
	}

	passwordIsValid, msg := CheckPassword(body.Password, user.Password)
	if passwordIsValid != true{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(user.Username)
	
	UpdateAllTokens(token, refreshToken, user.Username)
	
	user, err, msg = GetUser(ctx, body.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"token":         user.Token,
		"refresh_token": user.Refresh_token},
		"message": "Sign in successfull"})

}