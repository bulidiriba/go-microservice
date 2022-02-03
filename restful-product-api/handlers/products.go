package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/m/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	listProducts := data.GetProducts()
	res, err := json.Marshal(listProducts)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(res)

}
