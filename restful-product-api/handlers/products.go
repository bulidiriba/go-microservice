package handlers

import (
	"log"
	"net/http"
	"strconv"

	"example.com/m/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get requests")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
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

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put requests")
	vars := mux.Vars(r)
	id, idErr := strconv.Atoi(vars["id"])

	if idErr != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
	}

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	// if error found
	if err != nil {
		http.Error(rw, "Unable to decode the json value", http.StatusBadRequest)
	}

	er := data.UpdateProduct(id, prod)

	if er == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if er != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
