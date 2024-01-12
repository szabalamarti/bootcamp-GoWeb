package application

import (
	"fmt"
	"net/http"
	"supermarket/internal/auth"
	"supermarket/internal/auth/middleware"
	middlewareLog "supermarket/internal/platform/web/middleware"
	"supermarket/internal/product/handler"
	"supermarket/internal/product/repository"
	"supermarket/internal/product/service"
	"supermarket/internal/product/storage"

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
	// - dependencies
	// -- authenticator
	au := auth.NewAuthTokenBasic(s.token)
	auMiddleware := middleware.NewAuthenticator(au)

	// -- logger
	lgMd := middlewareLog.NewLogger()

	// create Repository
	storage := storage.NewProductStorage(s.dbFile)
	repository := repository.NewProductRepository(storage)

	// create service and handler
	service := service.NewProductService(repository)
	handler := handler.NewProductHandler(service)

	// router
	router := chi.NewRouter()

	// - middlewares
	// -- logger
	router.Use(lgMd.Log)

	// - routes
	router.Get("/ping", handler.GetPingHandler)

	router.Route("/products", func(router chi.Router) {
		router.Get("/", handler.GetProductsHandler)
		router.Get("/{id}", handler.GetProductHandler)
		router.Get("/search", handler.SearchProductsByPriceHandler)
		router.Get("/consumer_price", handler.GetConsumerPriceHandler)

		// subrouter with auth middleware
		router.With(auMiddleware.Auth).Group(func(router chi.Router) {
			router.Post("/", handler.CreateProductHandler)
			router.Patch("/{id}", handler.UpdateProductHandler)
			router.Delete("/{id}", handler.DeleteProductHandler)
			router.Put("/{id}", handler.UpdateOrCreateProductHandler)
		})
	})

	// start server
	fmt.Printf("Server started on %s:%s\n", s.host, s.port)
	http.ListenAndServe(s.host+":"+s.port, router)
	return nil
}
