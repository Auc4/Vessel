package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Auc4/Vessel/entities"
	"github.com/gin-gonic/gin"
)

func GetUsuarios(c *gin.Context) {

	url := "http://127.0.0.1:8090/api/collections/Usuario/records"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con Pocketbase"})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la respuesta de Pocketbase"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(body))

}

func PostUsuario(c *gin.Context) {

	url := "http://127.0.0.1:8090/api/collections/Usuario/records"

	var nuevoUsuario entities.CrearUsuario

	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Formato JSON inválido"})
	}

	jsonData, err := json.Marshal(nuevoUsuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los datos"})
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	defer resp.Body.Close()

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado con éxito"})

}
