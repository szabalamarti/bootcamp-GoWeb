package main

import (
	"net/http"
	"supermarket/internal/handlers"
	"supermarket/internal/repository"
	"supermarket/internal/service"

	"github.com/go-chi/chi/v5"
)

func initializeApplication() (*handlers.ProductHandler, error) {
	// Load repository from file
	ps := &repository.ProductRepository{}
	err := ps.LoadProducts("docs/db/products.json")
	if err != nil {
		return &handlers.ProductHandler{}, err
	}
	ps.PrintProductsInfo()

	// Create service and handler
	service := &service.ProductService{ProductRepository: ps}
	handler := &handlers.ProductHandler{ProductService: service}

	return handler, nil
}

func main() {
	// Create handler, initialize application.
	handler, err := initializeApplication()
	if err != nil {
		panic(err)
	}

	// Create router and routes
	router := chi.NewRouter()
	router.Get("/ping", handler.GetPingHandler)
	router.Route("/products", func(router chi.Router) {
		router.Get("/", handler.GetProductsHandler)
		router.Get("/{id}", handler.GetProductHandler)
		router.Get("/search", handler.SearchProductsHandler)
		router.Post("/", handler.CreateProductHandler)
	})

	// Start server
	http.ListenAndServe(":8080", router)
}
