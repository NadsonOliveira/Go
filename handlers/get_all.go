package handlers

import (
	"Go-teste/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("error ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
