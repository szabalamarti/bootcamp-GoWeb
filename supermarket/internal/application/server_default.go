package application

import (
	"net/http"
	"os"
	"supermarket/internal/handler"
	"supermarket/internal/repository"
	"supermarket/internal/service"

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
	repository := repository.NewProductRepository()
	err := repository.LoadProducts("docs/db/products.json")
	if err != nil {
		return err
	}
	repository.PrintProductsInfo()

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
	http.ListenAndServe(":8080", router)
	return nil
}
