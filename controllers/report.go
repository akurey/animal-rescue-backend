package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct{}

type AddReportBody struct {
	AnimalId    int64  `json:"animal_id"`
	ReporterId  int64  `json:"reporter_id"`
	FormId      int64  `json:"form_id"`
	FieldValues string `json:"field_values"`
}

type UpdateReportBody struct {
	AnimalId    int64  `json:"animal_id"`
	FieldValues string `json:"field_values"`
}

func (ctrl ReportController) GetAnimalRecord(context *gin.Context) {
	var report models.AnimalReport

	reportId := context.Param("id")
	err := database.DB.Raw("SELECT * FROM public.AFN_GetAnimalReport(?);", reportId).Scan(&report).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": report})
}

func (ctrl ReportController) GetReports(context *gin.Context) {
	var reports []*models.Report
	err := database.DB.Raw("SELECT * FROM public.AFN_GetAnimalReports();").Scan(&reports).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": reports})
}

func (ctrl ReportController) AddReport(context *gin.Context) {
	var report []*models.Report

	body := AddReportBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	err := database.DB.Raw("SELECT * FROM public.afn_addanimalreport(?,?,?,?) AS r(ID bigint, created_at text, is_approved bit(1), animal_name varchar(100));",
		body.AnimalId, body.ReporterId, body.FormId, body.FieldValues).Scan(&report).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": report})
}

func (ctrl ReportController) UpdateReport(context *gin.Context) {
	var report_field_values []*models.ReportFieldValue

	reportId := context.Param("id")
	body := UpdateReportBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	err := database.DB.Raw("SELECT * FROM public.AFN_UpdateAnimalReport(?, ?, ?);",
		reportId, body.AnimalId, body.FieldValues).Scan(&report_field_values).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": report_field_values})
}

func (ctrl ReportController) DeleteReport(context *gin.Context) {
	var report models.Report

	reportId := context.Param("id")
	err := database.DB.Raw("SELECT * FROM public.AFN_DeleteAnimalReport(?);", reportId).Scan(&report).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": report.ID})
}
