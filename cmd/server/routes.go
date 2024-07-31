package main

import (
	"encoding/json"
	"log"
	"net/http"
	"toughleaf/config"
	"toughleaf/internal/models"
	"toughleaf/internal/templates"

	"github.com/go-chi/chi/v5"
)

func routes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Home()
		component.Render(r.Context(), w)
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := config.Db.Create(&user).Error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
		log.Println("User created")
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		config.Db.Find(&users)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	})

}
