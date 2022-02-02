package main

import (
	"log"
	"net/http"
	"os"
	"poop/strawberry/monkey/handlers"
	"time"
)

func main() {
	// create a logger with logging to standard out, and the prefix is the service name, the flag will be standard flags
	// then inject this created log with the NewHello handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//http.ListenAndServe(":9090", sm)
}
