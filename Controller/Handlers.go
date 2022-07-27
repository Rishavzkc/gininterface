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
