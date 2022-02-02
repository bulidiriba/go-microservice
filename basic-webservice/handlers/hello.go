package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world!")
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
}
