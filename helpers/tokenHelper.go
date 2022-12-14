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
var TTL_TOKEN string = os.Getenv("TTL_TOKEN")
var TTL_REFRESHTOKEN string = os.Getenv("TTL_REFRESHTOKEN")

func GenerateUserTokens(username string) (signedToken string, signedRefreshToken string, err error){

	intToken, err := strconv.ParseInt(TTL_TOKEN, 0, 64)
	intRefreshToken, err := strconv.ParseInt(TTL_REFRESHTOKEN, 0, 64)
	
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


