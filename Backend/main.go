package main

import (
	"github.com/Auc4/Vessel/config"
)

func main() {

	config.ConnectDB()

	config.ConnectGin()

}
