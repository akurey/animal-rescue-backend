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
	var report models.AnimalReport

	reportId := context.Param("id")
	database.DB.Raw("SELECT * FROM public.AFN_GetAnimalReport(?) r(ID BIGINT, IdAnimal BIGINT, AnimalName VARCHAR, ScientificName VARCHAR, ConservationStatusName VARCHAR, Abbreviaton VARCHAR, ClassificationName VARCHAR, Fields JSONB);", reportId).Scan(&report)

	context.JSON(http.StatusOK, gin.H{"response": report})
}
type GetReportBody struct {
	 ShelterId int64 `json:"user_shelter_id"`
}

func (ctrl ReportController) GetReports(context *gin.Context) {
	var reports []*models.Report

	body := GetReportBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	err := database.DB.Table("\"AP_Animal_Reports\" AR").
		Select("AR.id, to_char(AR.created_at, 'DD/MM/YYYY') created_at, AR.is_approved, AN.name animal_name, CONCAT(DIS.name,', ',CA.name,', ',PR.name) place_of_rescue").
		Joins("INNER JOIN \"AP_Forms\" FO ON AR.form_id = FO.id").
		Joins("INNER JOIN \"AP_Animals\" AN ON AR.animal_id = AN.id").
		Joins("INNER JOIN \"AP_Fields\" FI ON FI.name = 'Dirección'").
		Joins("INNER JOIN \"AP_Report_Field_Values\" FV ON AR.id = FV.report_id AND FI.id = FV.field_id").
		Joins("INNER JOIN \"AP_Directions\" DIR ON CAST(FV.value AS BIGINT) = DIR.id").
		Joins("INNER JOIN \"AP_Districts\" DIS ON DIR.district_id = DIS.id").
		Joins("INNER JOIN \"AP_Cantons\" CA ON DIS.canton_id = CA.id").
		Joins("INNER JOIN \"AP_Provinces\" PR ON CA.province_id = PR.id").
		Where("FO.shelter_id = ?", body.ShelterId).
		Find(&reports).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": reports})
}
