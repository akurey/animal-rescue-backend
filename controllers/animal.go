package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnimalController struct{}

func (ctrl AnimalController) GetAnimals(context *gin.Context) {
	var animals []*models.Animal
	err := database.DB.Table("\"AP_Animals\" AA").
		Select("AA.id, AA.name, AA.scientific_name, ACS.name conservation_status, ACS.abbreviation conservation_abbreviation, AAC.name classification_name").
		Joins("INNER JOIN \"AP_Conservation_Status\" ACS ON ACS.id = AA.conservation_status_id").
		Joins("INNER JOIN \"AP_Animal_Classification\" AAC ON AAC.id = AA.classification_id").
		Where("AA.is_deleted = ?", 0).
		Find(&animals).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": animals})
}
