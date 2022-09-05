package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct{}

type GetReportBody struct {
	 ShelterId int64 `json:"user_shelter_id"`
}

func (ctrl ReportController) GetReports(context *gin.Context) {
	var reports []*models.Report

	body := GetReportBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	err := database.DB.Table("\"AP_Animal_Reports\" AR").
		Select("AR.id, AR.created_at, AR.is_approved, SH.name shelter_name, FO.name form_name, US.first_name reporter_name, US.last_name reporter_lastname").
		Joins("INNER JOIN \"AP_Forms\" FO ON AR.form_id = FO.id").
		Joins("INNER JOIN \"AP_Shelters\" SH ON FO.shelter_id = SH.id").
		Joins("INNER JOIN \"AP_Users\" US ON AR.reporter_id = US.id").
		Where("FO.shelter_id = ?", body.ShelterId).
		Find(&reports).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": reports})
}
