package application

import (
	"fmt"
	"net/http"
	"os"
	"supermarket/internal/handler"
	"supermarket/internal/repository"
	"supermarket/internal/service"
	"supermarket/internal/storage"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	address string
}

func NewServer(address string) *Server {
	defaultAddress := ":8080"
	if address == "" {
		address = defaultAddress
	}

	return &Server{
		address: address,
	}
}

func (s *Server) Start() error {
	// token for auth
	os.Setenv("Token", "themostsecrettoken")

	// create Repository
	storage := storage.NewProductStorage("docs/db/products.json")
	repository := repository.NewProductRepository(storage)

	// create service and handler
	service := service.NewProductService(repository)
	handler := handler.NewProductHandler(service)

	// create router and routes
	router := chi.NewRouter()
	router.Get("/ping", handler.GetPingHandler)
	router.Route("/products", func(router chi.Router) {
		router.Get("/", handler.GetProductsHandler)
		router.Get("/{id}", handler.GetProductHandler)
		router.Get("/search", handler.SearchProductsByPriceHandler)
		router.Post("/", handler.CreateProductHandler)
		router.Patch("/{id}", handler.UpdateProductHandler)
		router.Delete("/{id}", handler.DeleteProductHandler)
		router.Put("/{id}", handler.UpdateOrCreateProductHandler)
	})

	// start server
	fmt.Printf("Server started on port %s\n", s.address)
	http.ListenAndServe(":8080", router)
	return nil
}
