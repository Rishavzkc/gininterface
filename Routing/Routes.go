package Routing

import (
	"gininterface/Controller"

	"github.com/gin-gonic/gin"
)

func Handler() {
	r := gin.Default()
	r.GET("/company", Controller.GetCompany)
	r.POST("/company", Controller.CreateCompany)
	r.GET("/company/:id", Controller.GetCompanyById)
	r.PATCH("/company/:id", Controller.UpdateCompany)
	r.DELETE("comny/:id", Controller.DeleteCompany)

	r.Run()
}
