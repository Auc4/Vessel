package main

import (
	"github.com/Auc4/Vessel/config"
	"github.com/Auc4/Vessel/controllers"
)

func main() {

	db := config.ConnectDB()

	controllers.SetDB(db)

	config.ConnectGin()

}
