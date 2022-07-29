package Controller

import (
	"gininterface/Config"
	"gininterface/Model"
	"gininterface/Service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController interface {
	GetCompany(c *gin.Context) []Model.Company
	CreateCompany(c *gin.Context) Model.Company
	UpdateCompany(c *gin.Context) Model.Company
	DeleteCompany(c *gin.Context) string
}

type controller struct {
	service Service.CompanyService
}

func (con *controller) GetCompany(c *gin.Context) []Model.Company {
	var companies []Model.Company
	Config.Database.Find(&companies)
	c.JSON(http.StatusOK, gin.H{"data": companies})
	return con.service.GetCompany()
}
func (con *controller) CreateCompany(c *gin.Context) Model.Company {
	var company Model.Company
	c.BindJSON(&company)
	Config.Database.Create(&company)
	return company
}

func (con *controller) UpdateCompany(c *gin.Context) Model.Company {
	var company Model.Company

	if err := Config.Database.Where("id=?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found"})
	}
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	Config.Database.Model(&company).Updates(company)
	return company
}

func (con *controller) DeleteCompany(c *gin.Context) string {
	var company Model.Company
	if err := Config.Database.Where("id=?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found"})
	}

	Config.Database.Model(&company).Delete(company)
	return "Record deleted sucessfully"
}

func New(service Service.CompanyService) CompanyController {
	return &controller{
		service: service,
	}
}
