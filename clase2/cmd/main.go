package main

import (
	"clase2/internal/handler"
	"clase2/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	ps := &service.ProductService{}
	err := ps.LoadProducts("clase2/products.json")
	if err != nil {
		panic(err)
	}
	ps.PrintProductsInfo()

	h := &handler.Handler{ProductService: ps}

	r := chi.NewRouter()

	r.Get("/ping", h.GetPingHandler)
	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.GetProductsHandler)
		r.Get("/{id}", h.GetProductHandler)
		r.Get("/search", h.SearchProductsHandler)
	})

	http.ListenAndServe(":8080", r)
}
