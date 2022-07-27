package Controller

import (
	"gininterface/Config"
	"gininterface/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	var companies []Model.Company

	Config.Database.Find(&companies)
	c.JSON(http.StatusOK, gin.H{"data": companies})

}

func CreateCompany(c *gin.Context) {
	var company Model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company = Model.Company{Name: company.Name, Id: company.Id, Location: company.Location}
	Config.Database.Create(&company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

func GetCompanyById(c *gin.Context) {
	var company Model.Company

	if err := Config.Database.Where("id=?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}

func UpdateCompany(c *gin.Context) {
	var company Model.Company
	if err := Config.Database.Where("id=?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found"})
		return
	}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Config.Database.Model(&company).Updates(company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

func DeleteCompany(c *gin.Context) {
	var company Model.Company
	if err := Config.Database.Where("id=?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found"})
		return
	}

	Config.Database.Delete(&company)
	c.JSON(http.StatusOK, gin.H{"data": "Record Deleted Sucessfully"})
}
