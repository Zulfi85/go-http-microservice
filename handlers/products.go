package handlers

import (
	"go-http-microservice/product-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProductshan (w, r)
		return
	}
	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
	}
	


func (p *Products) getProductshan (w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ConvToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}