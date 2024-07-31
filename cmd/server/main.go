package main

import (
	"fmt"
	"net/http"
	"toughleaf/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	config.ConnectDB()

	routes(r)

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
