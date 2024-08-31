package controller

import "net/http"

func HelloWolrdController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}
