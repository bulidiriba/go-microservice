package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// lets create an handler for / endpoint, then it accepts two parameter
	// handler function should be written before calling the serving method
	// the endpoint(pattern) and the function with response and request parameter

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			// There is two way to display an error, using WriteHeader of response writer
			// Or a go package called an error
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Oooops"))
			// 0r
			http.Error(rw, "Ooops", http.StatusBadRequest)
			// we need the return for both case bc they dont terminate
			return
		}
		log.Printf("Data %s", data)

		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World!")
	})

	// any endpoint that doesn't match the above defined patterns redirect to the root endpoint

	// 1st create webservices using http package
	// it takes two parameter binding port and the handler
	http.ListenAndServe(":9090", nil)
}
