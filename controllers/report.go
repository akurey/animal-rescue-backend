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

type AddReportBody struct {
	AnimalId int64 `json:"animal_id"`
	ReporterId int64 `json:"reporter_id"`
	FormId int64 `json:"form_id"`
	FieldValues string `json:"field_values"`
}

type UpdateReportBody struct {
	ReportId int64 `json:"report_id"`
	FieldValues string `json:"field_values"`
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

func (ctrl ReportController) AddReport(context *gin.Context) {
	var report []*models.Report

	body := AddReportBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	database.DB.Raw("SELECT * FROM public.afn_addanimalreport(?,?,?,?) AS r(ID bigint, created_at text, is_approved bit(1), animal_name varchar(100));",
		body.AnimalId, body.ReporterId, body.FormId, body.FieldValues).Scan(&report);

	context.JSON(http.StatusOK, gin.H{"response": report})
}

func (ctrl ReportController) UpdateReport(context *gin.Context) {
	var report_field_values []*models.ReportFieldValue

	body := UpdateReportBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	database.DB.Raw("SELECT * FROM public.AFN_UpdateAnimalReport(?, ?);", 
		body.ReportId, body.FieldValues).Scan(&report_field_values);

	context.JSON(http.StatusOK, gin.H{"response": report_field_values})
}
