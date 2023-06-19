package helpers

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
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

	intToken, err := strconv.ParseInt(os.Getenv("TTL_TOKEN"), 0, 64)
	intRefreshToken, err := strconv.ParseInt(os.Getenv("TTL_REFRESHTOKEN"), 0, 64)
	
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

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(SECRET_KEY)), nil
	})
	if err != nil {
		fmt.Println(err, 123)
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}


