package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnimalController struct{}

type AnimalBody struct {
	AnimalName string `json:"animal_name"` 
	ScientificName string `json:"scientific_name"` 
	ConservationStatusId int `json:"conservation_status_id"` 
	ClassificationId int `json:"classification_id"`
}

func (ctrl AnimalController) GetAnimals(context *gin.Context) {
	var animals []*models.Animal
	err := database.DB.Table("\"AP_Animals\" AA").
		Select("AA.id, AA.name, AA.scientific_name, AA.created_at, ACS.name conservation_status, ACS.abbreviation conservation_abbreviation, AAC.name classification_name").
		Joins("INNER JOIN \"AP_Conservation_Status\" ACS ON ACS.id = AA.conservation_status_id").
		Joins("INNER JOIN \"AP_Animal_Classification\" AAC ON AAC.id = AA.classification_id").
		Where("AA.is_deleted = ?", 0).
		Find(&animals).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": animals})
}

func (ctrl AnimalController) AddAnimal(context *gin.Context) {
	var animal models.Animal

	body := AnimalBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	err := database.DB.Raw("SELECT * FROM AFN_AddAnimal(?,?,?,?);",
	       body.AnimalName, body.ScientificName, body.ConservationStatusId, body.ClassificationId).Scan(&animal).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": animal})
}

func (ctrl AnimalController) UpdateAnimal(context *gin.Context) {
	var animal models.Animal

	animalId := context.Param("id")
	body := AnimalBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	err := database.DB.Raw("SELECT * FROM public.AFN_UpdateAnimal(?, ?, ?, ?, ?);", 
	       animalId, body.AnimalName, body.ScientificName, body.ConservationStatusId, body.ClassificationId).Scan(&animal).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": animal})
}

func (ctrl AnimalController) DeleteAnimal(context *gin.Context) {
	var animal models.Animal

	animalId := context.Param("id")
	err := database.DB.Raw("SELECT * FROM public.AFN_DeleteAnimal(?);", animalId).Scan(&animal).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": animal})
}
