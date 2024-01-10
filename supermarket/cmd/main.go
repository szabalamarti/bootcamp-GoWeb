package main

import (
	"net/http"
	"supermarket/internal/handler"
	"supermarket/internal/repository"
	"supermarket/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Load repository from file
	ps := &repository.ProductRepository{}
	err := ps.LoadProducts("supermarket/docs/db/products.json")
	if err != nil {
		panic(err)
	}
	ps.PrintProductsInfo()

	// Create service and handler
	service := &service.ProductService{ProductRepository: ps}
	handler := &handler.Handler{ProductService: service}

	// Create router and routes
	router := chi.NewRouter()
	router.Get("/ping", handler.GetPingHandler)
	router.Route("/products", func(router chi.Router) {
		router.Get("/", handler.GetProductsHandler)
		router.Get("/{id}", handler.GetProductHandler)
		router.Get("/search", handler.SearchProductsHandler)
		router.Post("/", handler.AddProductHandler)
	})

	// Start server
	http.ListenAndServe(":8080", router)
}
