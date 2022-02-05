package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if r.Method == http.MethodPut {
		// expect the id in the URL
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI 1", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI 2", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to number", idString)
			http.Error(rw, "Invalid URI 3", http.StatusBadRequest)
			return
		}
		p.l.Println("got id", id)
		p.updateProducts(id, rw, r)
	}

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

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put requests")

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

	if er == nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
