package cmd

import (
	"log"
	"myAdmin/router"
	"net/http"
	"time"
)

func NewServerAdmin(r router.AdminRouter) {
	// Создаём сервер с таймаутами
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server is running on http://127.0.0.1:8000/admin")
	log.Fatal(srv.ListenAndServe())
}
