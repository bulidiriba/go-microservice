package main

import (
	"log"
	"net/http"
	"os"
	"poop/strawberry/monkey/handlers"
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

	http.ListenAndServe(":9090", sm)
}
