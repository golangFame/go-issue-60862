package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riandyrn/otelchi"
)

func main() {
	router := chi.NewRouter()
	router.Use(otelchi.Middleware("my-service", otelchi.WithChiRoutes(router)))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "hello world"
		w.Write([]byte(msg))
		log.Println(msg)
	})

	server := http.Server{
		Handler: router,
		Addr:    ":3804",
	}
	log.Printf("server started at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err) // never gets called
	}
}
