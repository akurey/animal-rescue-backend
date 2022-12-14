package helpers

import (
	"os"
	"time"
	"strconv"
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct{
	Username	string
	jwt.StandardClaims 
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")
var TOKEN string = os.Getenv("TOKEN")
var REFRESHTOKEN string = os.Getenv("REFRESHTOKEN")

func GenerateUserTokens(username string) (signedToken string, signedRefreshToken string, err error){

	intToken, err := strconv.ParseInt(TOKEN, 0, 64)
	intRefreshToken, err := strconv.ParseInt(REFRESHTOKEN, 0, 64)

	claims := &SignedDetails{
		Username : username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(intToken)).Unix(),
	
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(intRefreshToken)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	return token, refreshToken, err
}


