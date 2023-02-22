package main

import (
	"Go-teste/configs"
	"Go-teste/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.GetAll)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
