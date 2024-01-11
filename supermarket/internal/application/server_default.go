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
	host   string
	port   string
	dbFile string
	token  string
}

type ServerConfig struct {
	Host   string
	Port   string
	DbFile string
	Token  string
}

func NewServer(config ServerConfig) *Server {
	// default values
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "8080"
	}
	if config.DbFile == "" {
		config.DbFile = "docs/db/products.json"
	}

	return &Server{
		host:   config.Host,
		port:   config.Port,
		dbFile: config.DbFile,
		token:  config.Token,
	}
}

func (s *Server) Start() error {
	// Set auth token from config
	os.Setenv("Token", s.token)

	// create Repository
	storage := storage.NewProductStorage(s.dbFile)
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
	fmt.Printf("Server started on %s:%s\n", s.host, s.port)
	http.ListenAndServe(s.host+":"+s.port, router)
	return nil
}
