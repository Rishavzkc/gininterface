package Config

import (
	"fmt"
	"gininterface/Model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlDSN = "root:Quest1234@tcp(localhost:3306)/newdb?parseTime=true"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Connection Failed: ")
	}
	Database.AutoMigrate(&Model.Company{})
}
