package handlers

import (
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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// to return the existing product
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	// to add new product
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// to replace the existing product

	// to update some field of the existing product

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get requests")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post requests")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	// if error found
	if err != nil {
		http.Error(rw, "Unable to decode the json value", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
