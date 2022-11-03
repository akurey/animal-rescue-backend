package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormController struct{}

func (ctrl FormController) GetFormFields(context *gin.Context) {
	var formField []*models.FormField
	formId := context.Param("id")
	database.DB.Raw("SELECT * FROM public.AFN_GetFormFields(?);", formId).Scan(&formField);

	context.JSON(http.StatusOK, gin.H{"response": formField})
}

func (ctrl FormController) GetAddressOptions(context *gin.Context) {
	var addressField []*models.AdressField
	database.DB.Raw("SELECT * FROM public.AFN_GetAddressOptions();").Scan(&addressField);

	context.JSON(http.StatusOK, gin.H{"response": addressField})
}
