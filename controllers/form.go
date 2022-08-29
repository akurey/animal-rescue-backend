package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormController struct{}

func (ctrl FormController) GetFormFields(context *gin.Context) {
	var formField []*models.FormField
	formId := context.Param("id")
	err := database.DB.Table("\"AP_Forms\" AF").
		Select("AF.name form_name, AFI.name field_name, AFI.is_required, AFT.name field_type, AFO.option").
		Joins("INNER JOIN \"AP_Form_Fields\" AFF ON AFF.form_id = AF.id").
		Joins("INNER JOIN \"AP_Fields\" AFI ON AFF.field_id = AFI.id").
		Joins("INNER JOIN \"AP_Field_Types\" AFT ON AFT.id = AFI.field_type_id").
		Joins("LEFT JOIN \"AP_Field_Options\" AFO ON AFO.field_id = AFI.id").
		Where("AF.is_deleted = ? and AF.id = ? and AFF.is_deleted = ? and AFI.is_deleted = ? and AFT.is_deleted = ? and (AFO.is_deleted IS NULL OR AFO.is_deleted = ?)", 0, formId, 0, 0, 0, 0).
		Find(&formField).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": formField})
}
