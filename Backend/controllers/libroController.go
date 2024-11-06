package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Auc4/Vessel/entities"
	"github.com/gin-gonic/gin"
)

func GetLibros(c *gin.Context) {

	id_usuario := c.Param("id_usuario")

	url := fmt.Sprintf("http://127.0.0.1:8090/api/collections/Libro/records?filter=(id_usuario='%s')", id_usuario)

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

func PostLibro(c *gin.Context) {
	var nuevoLibro entities.Libro

	url := "http://127.0.0.1:8090/api/collections/Libro/records"

	if err := c.ShouldBindJSON(&nuevoLibro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de JSON inválido"})
		return
	}

	jsonData, err := json.Marshal(nuevoLibro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los datos"})
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el libro"})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusCreated, gin.H{"message": "Libro creado con éxito"})
}
