package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewAddMenuController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			name := r.Form["name"][0]
			info_menu := r.Form["info_menu"][0]
			harga_menu := r.Form["harga_menu"][0]

			_, err := db.Exec("INSERT INTO menu (name,info,harga) VALUES(?, ?, ?)", name, info_menu, harga_menu)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/menu", http.StatusMovedPermanently)
			return

		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
