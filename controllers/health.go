package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (ctrl HealthController) Check(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"response": "OK"})
}
