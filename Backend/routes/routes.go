package routes

import (
	"github.com/Auc4/Vessel/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {

	router.GET("/api/usuarios", controllers.GetUsuarios)
	router.POST("/api/crear_usuario", controllers.PostUsuario)
	router.GET("/api/libros/usuario/:id_usuario", controllers.GetLibros)
	router.POST("/api/crear_libro", controllers.PostLibro)

}
