package controllers

import (
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: remove this file after create the first controller.
type PingController struct{}

func (ctrl PingController) Ping(context *gin.Context) {
	var pongModel = new(models.Pong)

	pongModel.ID = 13
	pongModel.Message = "Pong"

	context.JSON(http.StatusOK, gin.H{"response": pongModel})
}
