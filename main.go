package main

import (
	"CRUD-GO/database"
	"CRUD-GO/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
