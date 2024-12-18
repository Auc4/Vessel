package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB maneja la conexión a la base de datos y la retorna
func ConnectDB() *sql.DB {
	// Configura la conexión: usuario, contraseña, servidor y nombre de la base de datos
	dsn := "vessel_user:Vessel.Pass.w0rd-24@tcp(localhost:3306)/Vessel"

	// Abre la conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Verifica la conexión
	if err = db.Ping(); err != nil {
		log.Fatalf("Error al verificar la conexión: %v", err)
	}

	fmt.Println("Conexión a MySQL exitosa")

	return db
}
