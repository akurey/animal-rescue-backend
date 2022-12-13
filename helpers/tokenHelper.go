package helpers

import (
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct{
	Username	string
	jwt.StandardClaims 
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateUserTokens(username string) (signedToken string, signedRefreshToken string, err error){
	claims := &SignedDetails{
		Username : username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
			
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	return token, refreshToken, err
}


