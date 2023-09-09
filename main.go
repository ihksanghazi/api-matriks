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
	r.Post("/add", controllers.MatriksPertambahan)
	r.Post("/subtract", controllers.MatriksPengurangan)
	http.ListenAndServe(":3000", r)
}
