package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ihksanghazi/api-matriks/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/add", controllers.MatriksPertambahan)
	r.Post("/api/subtract", controllers.MatriksPengurangan)
	r.Post("/api/multiply", controllers.MatriksPerkalian)
	r.Post("/api/transpose", controllers.MatriksTranspose)
	r.Post("/api/inverse", controllers.MatriksInverse)
	r.Post("/api/determinan", controllers.MatriksDeterminan)
	http.ListenAndServe(":3000", r)
}
