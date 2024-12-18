package controllers

import (
	"log"
	"net/http"

	"github.com/Auc4/Vessel/entities"
	"github.com/gin-gonic/gin"
)

func GetUsuarioByID(c *gin.Context) {

	var usuario entities.Usuario

	Query := "SELECT Usuario_ID, Nombre_usuario, Password_usuario, Email FROM usuario WHERE Usuario_ID = ?;"

	UsuarioID := c.Param("id")

	err := DB.QueryRow(Query, UsuarioID).Scan(&usuario.UsuarioID, &usuario.NombreUsuario, &usuario.PasswordUsuario, &usuario.Email)

	if err != nil {
		log.Println("Error al consultar los datos del usuario", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ha ocurrido un fallo al consultar los datos del usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usuario": usuario})

}

func PostUsuario(c *gin.Context) {

	var usuario entities.Usuario

	var errores []string

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

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

	err = DB.QueryRow("SELECT COUNT(*) FROM usuario WHERE Nombre_usuario = ?;", usuario.NombreUsuario).Scan(&count)
	if err != nil {
		log.Println("Error al verificar si el usuario existe: ", err)
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Este nombre de usuario ya ha sido utilizado"})
		return
	}

	if usuario.NombreUsuario == "" {
		errores = append(errores, "Se debe ubicar un nombre de usuario")
	}
	if usuario.Email == "" {
		errores = append(errores, "Se debe ubicar un email para registrarse")
	}

	if len(errores) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errores})
		return
	}

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

	Query := "UPDATE usuario SET Nombre_usuario = ?, Password_usuario = ?, Email = ? WHERE Usuario_ID = ?;"

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

	_, err := DB.Exec(Query, usuario.NombreUsuario, usuario.PasswordUsuario, usuario.Email, UsuarioID)
	if err != nil {
		log.Println("Error al actualizar el usuario", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}

func DeleteUsuario(c *gin.Context) {

	UsuarioID := c.Param("id")

	Query := "SELECT COUNT(*) FROM usuario WHERE Usuario_ID = ?"

	var count int
	err := DB.QueryRow(Query, UsuarioID).Scan(&count)
	if err != nil {
		log.Println("Error al verificar la existencia del usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "El usuario no existe"})
		return
	}

	Query = "DELETE FROM usuario WHERE Usuario_ID = ?"

	_, err = DB.Exec(Query, UsuarioID)
	if err != nil {
		log.Println("Error al eliminar el usuario", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "El usuario se ha eliminado exitosamente"})
}
