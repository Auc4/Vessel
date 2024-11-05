package config

import (
	"log"

	"github.com/Auc4/Vessel/routes"
	"github.com/gin-gonic/gin"
)

func ConnectGin() {

	r := gin.Default()

	routes.SetUpRoutes(r)

	if err := r.Run(":9090"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}

}
