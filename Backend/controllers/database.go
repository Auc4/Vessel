package controllers

import "database/sql"

var DB *sql.DB

// SetDB asigna la conexión de base de datos a la variable DB en el paquete controllers
func SetDB(db *sql.DB) {
	DB = db
}
