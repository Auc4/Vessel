package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Auc4/Vessel/entities"
	"github.com/gin-gonic/gin"
)

var libro entities.Libro

func GetLibros(c *gin.Context) {

	UsuarioID := c.Param("id")

	Query := "SELECT Titulo, Autor, Año_publicacion, favorito FROM usuario INNER JOIN libros ON usuario.Usuario_ID = libros.Usuario_ID WHERE usuario.Usuario_ID = ?;"

	var libros []entities.Libro

	//Ejecución del Query con Join
	rows, err := DB.Query(Query, UsuarioID)
	if err != nil {
		log.Println("Error al obtener los libros:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los libros"})
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&libro.TituloLibro, &libro.AutorLibro, &libro.FechaPublicacion, &libro.Favorito); err != nil {
			log.Println("Error al leer los datos de los libros:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los libros"})
			return
		}
		libros = append(libros, libro)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al iterar las filas:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los libros"})
		return
	}

	if len(libros) == 0 {
		log.Println("No se encontraron libros para el usuario:", UsuarioID)
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron libros para este usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"libros": libros})
}

func PostLibro(c *gin.Context) {

	var errores []string

	Query := "INSERT INTO libros (Usuario_ID, Titulo, Autor, Año_publicacion, Favorito) VALUES (?, ?, ?, ?, ?);"

	// Pasar el formato JSON a la estructura del libro
	if err := c.ShouldBindJSON(&libro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	//Validaciones para verificar si el libro tiene los datos necesarios
	if libro.TituloLibro == "" {
		errores = append(errores, "Se debe ubicar un título de libro")
	}
	if libro.AutorLibro == "" {
		errores = append(errores, "Se debe ubicar un autor del libro")
	}

	if len(errores) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errores})
		return
	}

	//Inserción del libro en la base de datos
	_, err := DB.Exec(Query, libro.Usuario_ID, libro.TituloLibro, libro.AutorLibro, libro.FechaPublicacion, libro.Favorito)
	if err != nil {
		log.Println("Error al insertar el libro", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar el libro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Libro creado exitosamente",
		"libro":   libro,
	})

}

func PutLibro(c *gin.Context) {

	UsuarioID := c.Param("usuario_id")
	LibroID := c.Param("libro_id")

	var libro entities.Libro

	var errores []string

	Query := "UPDATE libros SET Titulo = ?, Autor = ?, Año_publicacion = ?, Favorito = ? WHERE Libro_ID = ? AND Usuario_ID = ?;"

	if err := c.ShouldBindJSON(&libro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	//Validaciones para verificar si el libro tiene los datos necesarios
	if libro.TituloLibro == "" {
		errores = append(errores, "Se debe ubicar un título de libro")
	}
	if libro.AutorLibro == "" {
		errores = append(errores, "Se debe ubicar un autor del libro")
	}

	if len(errores) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errores})
		return
	}

	_, err := DB.Exec(Query, libro.TituloLibro, libro.AutorLibro, libro.FechaPublicacion, libro.Favorito, LibroID, UsuarioID)
	if err != nil {
		log.Println("Error al acutlaizar el libro", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el libro"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Libro actualizado exitosamente",
		"libro":   libro,
	})
}

func GetLibroById(c *gin.Context) {

	UsuarioID := c.Param("usuario_id")
	LibroID := c.Param("libro_id")

	var focusedLibro entities.FocusedLibro

	Query := `
        SELECT 
            Titulo, 
            Autor, 
            Año_publicacion, 
            Favorito,
            GROUP_CONCAT(categoria.Descripcion_categoria) AS Etiquetas
        FROM 
            libros

        LEFT JOIN 
            categorialibro ON libros.Libro_ID = categorialibro.Libro_ID
        LEFT JOIN 
            categoria ON categorialibro.Categoria_ID = categoria.Categoria_ID
        
		WHERE libros.Usuario_ID = ? AND libros.Libro_ID = ?
        
		GROUP BY libros.Libro_ID;
    `

	row := DB.QueryRow(Query, UsuarioID, LibroID)

	var etiquetas string // GROUP_CONCAT devuelve una cadena separada por comas

	if err := row.Scan(&focusedLibro.TituloLibro, &focusedLibro.AutorLibro, &focusedLibro.FechaPublicacion, &focusedLibro.Favorito, &etiquetas); err != nil {
		log.Println("Error al obtener los datos del libro:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos del libro"})
		return
	}

	// Convierte las etiquetas separadas por comas en un slice
	focusedLibro.Etiquetas = []string{}
	if etiquetas != "" {
		focusedLibro.Etiquetas = strings.Split(etiquetas, ",")
	}

	c.JSON(http.StatusOK, gin.H{"libro": libro})
}

func DeleteLibro(c *gin.Context) {
	UsuarioID := c.Param("usuario_id")
	LibroID := c.Param("libro_id")

	Query := "DELETE FROM libros WHERE Libro_ID = ? AND Usuario_ID = ?"
	result, err := DB.Exec(Query, LibroID, UsuarioID)
	if err != nil {
		log.Println("Error al eliminar el libro:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el libro"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "El libro no existe o no pertenece al usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "El libro se ha eliminado exitosamente"})
}
