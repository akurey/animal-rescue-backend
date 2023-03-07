package controllers

import (
	"animal-rescue-be/database"
	"animal-rescue-be/helpers"
	"animal-rescue-be/models"
	"net/http"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FormController struct{}

func (ctrl FormController) GetFormFields(context *gin.Context) {
	var formField []*models.FormField
	formId := context.Param("id")
	err := database.DB.Raw("SELECT * FROM public.AFN_GetFormFields(?);", formId).Scan(&formField).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": formField})
}

func (ctrl FormController) getProvinces(context *gin.Context) []models.ProvinceModel{
	var provinceField []*models.ProvinceField
	var provinceModel []models.ProvinceModel
	err := database.DB.Raw("SELECT * FROM public.AFN_GetProvinces();").Scan(&provinceField).Error
	helpers.HandleErr(err)


	for _, province := range provinceField {
		cantons := ctrl.getCantonByProvinceID(context, province.Id)
		provinceModelValue := models.ProvinceModel{Id: province.Id, Province: province.Province, Cantons: cantons}
		provinceModel = append(provinceModel, provinceModelValue)
	} 

	return provinceModel
	
}

func (ctrl FormController) getCantonByProvinceID(context *gin.Context, id int) []models.CantonModel{
	var cantonFields []*models.CantonField
	var cantonModel []models.CantonModel
	query := fmt.Sprintf("SELECT * FROM public.AFN_GetCantonByProvince(%s);",strconv.Itoa(id))
	err := database.DB.Raw(query).Scan(&cantonFields).Error
	helpers.HandleErr(err)

	for _, canton := range cantonFields {
		districts := ctrl.getDistrictByCantonID(context, canton.Id)
		cantonModelValue := models.CantonModel{Id: canton.Id, Canton: canton.Canton, Districts: districts}
		cantonModel = append(cantonModel, cantonModelValue)
	}

	return cantonModel
	
}

func (ctrl FormController) getDistrictByCantonID(context *gin.Context, id int) []*models.DistrictField{
	var districtFields []*models.DistrictField
	query := fmt.Sprintf("SELECT * FROM public.AFN_GetDistrictByCanton(%s);",strconv.Itoa(id))
	err := database.DB.Raw(query).Scan(&districtFields).Error
	helpers.HandleErr(err)

	return districtFields
	
}

func (ctrl FormController) GetAddressOptions(context *gin.Context) {
	var addressFields []models.ProvinceModel = ctrl.getProvinces(context)

	context.JSON(http.StatusOK, gin.H{"response": addressFields})
}
