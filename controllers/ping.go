package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"animal-rescue-be/models"
)

//PingController ...
type PingController struct{}

//Ping ...
func (ctrl PingController) Ping(c *gin.Context) {

	var pongModel = new(models.Pong)

	pongModel.ID = 13
	pongModel.Message = "Pong"

	c.JSON(http.StatusOK, gin.H{"response": pongModel})
}
