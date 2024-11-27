package routes

import (
	"github.com/Auc4/Vessel/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {

	usuario := router.Group("/api/usuario")
	{
		usuario.GET("/obtener_usuario/:id", controllers.GetUsuarioByID)
		usuario.POST("/crear_usuario", controllers.PostUsuario)
		usuario.PUT("/actualizar_usuario/:id", controllers.PutUsuario)
		usuario.DELETE("/borrar_usuario/:id", controllers.DeleteUsuario)
	}

	libro := router.Group("/api/libros")
	{
		libro.GET("/obtener_libros/:id")
	}
}
