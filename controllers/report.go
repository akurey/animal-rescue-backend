package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct{}

func (ctrl ReportController) GetAnimalRecord(context *gin.Context) {
	var formField []*models.Report
	reportId := context.Param("id")
	err := database.DB.Exec("CALL public.ASP_GetAnimalReport($1, $2); fetch all in $3;", reportId, "test", "test").
	Find(&formField).Error

	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": formField})
}