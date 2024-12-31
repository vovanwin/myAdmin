package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type AdminRouter struct {
	*mux.Router
}

func NewRouter() AdminRouter {
	// Создаём подмаршруты для /admin
	r := mux.NewRouter().PathPrefix("/admin").Subrouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/admin/static/", http.FileServer(http.Dir("resources"))))

	return AdminRouter{
		r,
	}
}
