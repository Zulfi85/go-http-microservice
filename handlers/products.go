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
		p.getProductshandle (w, r)
		return
	}
	
	if r.Method == http.MethodPost {
		p.postProductshandle (w, r)
		return
	}
	
	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
	}
	


func (p *Products) getProductshandle (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Get Product Request") //Get Product returns the proucts from data store.
	lp := data.GetProducts()
	err := lp.ConvToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) postProductshandle (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Post Product Request") //Post Products handle add products to data store. 

	}
