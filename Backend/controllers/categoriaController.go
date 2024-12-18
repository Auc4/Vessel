package controllers

import (
	"log"
	"net/http"

	"github.com/Auc4/Vessel/entities"
	"github.com/gin-gonic/gin"
)

var categoria entities.Categoria

func GetCategorias(c *gin.Context) {

	Query := "SELECT Categoria_ID, Descripcion_Categoria FROM categoria WHERE Estado_Categoria = true;"

	var categorias []entities.Categoria

	rows, err := DB.Query(Query)
	if err != nil {
		log.Println("Error al obtener las categorías", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las categorías"})
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&categoria.CategoriaID, &categoria.DescripcionCategoria); err != nil {
			log.Println("Error al leer las categorías existentes", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los datos de las categorías"})
			return
		}
		categorias = append(categorias, categoria)
	}

	c.JSON(http.StatusOK, gin.H{"categorias": categorias})

}
