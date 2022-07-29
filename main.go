package main

import (
	"gininterface/Config"
	"gininterface/Routing"
)

func main() {

	Config.DataMigration()
	Routing.Handler()
}
