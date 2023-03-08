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
	err := database.DB.Raw("SELECT * FROM public.AFN_GetFormFields(?);", formId).Scan(&formField).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": formField})
}


func (ctrl FormController) groupByProvince(data []*models.AddressField) []models.ProvinceModel {
	var provinceModel []models.ProvinceModel
	provinceMap := make(map[string]int)
	
	for i, province := range data{
		provinceMap[province.Province] = i
	}

	for province := range provinceMap {
		var cantonModelValue = ctrl.groupByCanton(data, province)
		var provinceModelValue = models.ProvinceModel{Province: province, Cantons: cantonModelValue}
		provinceModel = append(provinceModel, provinceModelValue)
	}

	
	return provinceModel
}


func (ctrl FormController) groupByCanton(data []*models.AddressField, province string) []models.CantonModel{
	var cantonByProvinceList []*models.AddressField
	var cantonModel []models.CantonModel
	cantonByProvinceListMap := make(map[string]int)

	for index, address := range data {
		if address.Province == province {
			cantonByProvinceListMap[address.Canton] = index
			cantonByProvinceList = append(cantonByProvinceList, address)
		}
	}

	for canton := range cantonByProvinceListMap {
		 var districtByCanton = ctrl.groupByDistrict(cantonByProvinceList, canton)
		 var cantonModelValue = models.CantonModel{ Canton: canton, Districts: districtByCanton }
		 cantonModel = append(cantonModel, cantonModelValue)
	}

	return cantonModel
	
}

func (ctrl FormController) groupByDistrict(data []*models.AddressField, canton string) []models.DistrictModel{
	var districtByCanton []models.DistrictModel
	for _, address := range data {
		if address.Canton == canton {
			districtByCanton = append(districtByCanton, models.DistrictModel{Id: address.Id, District: address.District} )
		}
	}

	return districtByCanton
}

func (ctrl FormController) GetAddressOptions(context *gin.Context) {
	var addressField []*models.AddressField
	err := database.DB.Raw("SELECT * FROM public.AFN_GetAddressOptions();").Scan(&addressField).Error
	helpers.HandleErr(err)
	var provinceModel = ctrl.groupByProvince(addressField)
	context.JSON(http.StatusOK, gin.H{"response": provinceModel})
}
