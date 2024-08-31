package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Menu struct {
	Id    string
	Name  string
	Info  string
	Harga string
}

func NewIndexmenu(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, info, harga FROM menu")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var menus []Menu
		for rows.Next() {
			var m Menu

			err = rows.Scan(
				&m.Id,
				&m.Name,
				&m.Info,
				&m.Harga,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			menus = append(menus, m)
		}

		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["menus"] = menus

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
