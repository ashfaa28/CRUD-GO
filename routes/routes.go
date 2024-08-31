package routes

import (
	"CRUD-GO/controller"
	"database/sql"
	"net/http"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewIndexmenu(db))
	server.HandleFunc("/menu", controller.NewIndexmenu(db))
	server.HandleFunc("/menu/add", controller.NewAddMenuController(db))
	server.HandleFunc("/menu/update", controller.NewUpdateMenuController(db))
	server.HandleFunc("/menu/delete", controller.NewDeleteMenuController(db))
}
