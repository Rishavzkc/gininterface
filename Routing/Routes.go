package Routing

import (
	"gininterface/Controller"
	"gininterface/Service"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	companyService    Service.CompanyService       = Service.New()
	companyController Controller.CompanyController = Controller.New(companyService)
)

func Handler() {
	r := gin.Default()
	r.GET("/company", func(c *gin.Context) {
		c.JSON(http.StatusOK, companyController.GetCompany(c))
	})
	r.POST("/company", func(c *gin.Context) {
		c.JSON(http.StatusOK, companyController.CreateCompany(c))
	})

	r.PATCH("/company/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, companyController.UpdateCompany(c))
	})

	r.DELETE("/company/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, companyController.DeleteCompany(c))
	})

	
	r.Run()
}
