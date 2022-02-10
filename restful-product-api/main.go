package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/m/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	prodHandler := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	serverMux := mux.NewRouter()
	getRouter := serverMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", prodHandler.GetProducts)

	putRouter := serverMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", prodHandler.UpdateProducts)
	putRouter.Use(prodHandler.MiddlewareProductValidation)

	postRouter := serverMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/add", prodHandler.AddProduct)
	postRouter.Use(prodHandler.MiddlewareProductValidation)

	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	serve := &http.Server{
		Addr:         ":9091",
		Handler:      corsHandler(serverMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := serve.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	serve.Shutdown(tc)

}
