package main

import (
	"gestor/Config"
	db "gestor/Config/database"
	"gestor/Routes"
	"log"
)

func main() {

	if err := db.ConfigurarDB(); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	router := Config.SetupServer()
	Routes.SetupRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
