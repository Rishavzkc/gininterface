package main

import (
	"gininterface/Routing"
	"ginnterface/Config"
)

func main() {

	Config.DataMigration()
	Routing.Handler()
}
