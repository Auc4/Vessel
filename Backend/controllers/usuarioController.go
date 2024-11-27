package controllers

import (
	"log"
	"net/http"

	"github.com/Auc4/Vessel/entities"
	"github.com/gin-gonic/gin"
)

func GetUsuarioByID(c *gin.Context) {

	var usuario entities.Usuario

	UsuarioID := c.Param("id")

	err := DB.QueryRow("SELECT Nombre_usuario, Password_usuario, Email FROM usuario WHERE Usuario_ID = ?;", UsuarioID).Scan(&usuario.NombreUsuario, &usuario.PasswordUsuario, &usuario.Email)

	if err != nil {
		log.Println("Error al consultar los datos del usuario", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ha ocurrido un fallo al consultar los datos del usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usuario": usuario})

}

func PostUsuario(c *gin.Context) {

	var usuario entities.Usuario

	// Pasar el formato JSON a la estructura del usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	// Validaciones para verificar si el email ya ha sido usado
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM usuario WHERE Email = ?;", usuario.Email).Scan(&count)
	if err != nil {
		log.Println("Error al verificar si el usuario existe:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar usuario"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Este correo ya ha sido utilizado"})
		return
	}

	// Validaciones para verificar si el nombre de usuario ya existe
	err = DB.QueryRow("SELECT COUNT(*) FROM usuario WHERE Nombre_usuario = ?;", usuario.NombreUsuario).Scan(&count)
	if err != nil {
		log.Println("Error al verificar si el usuario existe: ", err)
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Este nombre de usuario ya ha sido utilizado"})
		return
	}

	// Inserción del usuario en la base de datos
	_, err = DB.Exec("INSERT INTO usuario (Nombre_usuario, Password_usuario, Email) VALUES (?, ?, ?);", usuario.NombreUsuario, usuario.PasswordUsuario, usuario.Email)
	if err != nil {
		log.Println("Error al insertar el usuario:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"usuario": usuario,
	})
}

func PutUsuario(c *gin.Context) {

	UsuarioID := c.Param("id")

	var usuario entities.Usuario
	var errores []string

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	if usuario.NombreUsuario == "" {
		errores = append(errores, "El nombre de usuario no puede estar vacío")
	}
	if usuario.PasswordUsuario == "" {
		errores = append(errores, "La contraseña no puede estar vacía")
	}
	if usuario.Email == "" {
		errores = append(errores, "El email no puede estar vacío")
	}

	if len(errores) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errores})
		return
	}

	_, err := DB.Exec("UPDATE usuario SET Nombre_usuario = ?, Password_usuario = ?, Email = ? WHERE Usuario_ID = ?;", usuario.NombreUsuario, usuario.PasswordUsuario, usuario.Email, UsuarioID)
	if err != nil {
		log.Println("Error al actualizar el usuario", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}

func DeleteUsuario(c *gin.Context) {

	UsuarioID := c.Param("id")

	//Verifica si el usuario existe antes de eliminarlo
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM usuario WHERE Usuario_ID = ?", UsuarioID).Scan(&count)
	if err != nil {
		log.Println("Error al verificar la existencia del usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "El usuario no existe"})
		return
	}

	//Ejecución del Query
	_, err = DB.Exec("DELETE FROM usuario WHERE Usuario_ID = ?", UsuarioID)
	if err != nil {
		log.Println("Error al eliminar el usuario", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "El usuario se ha eliminado exitosamente"})
}
