package controllers

import (
	"animal-rescue-be/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)
type ErrorController struct{}

func (ctrl ErrorController) Error(context *gin.Context) {
	var tokenUrl = "/error"
	var response, err = errors.CodeRunner(tokenUrl,false); 
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"response": err})
	} else {
		context.JSON(http.StatusOK, gin.H{"response": response})
	}
}
