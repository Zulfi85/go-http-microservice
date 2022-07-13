package handlers

import (
	"go-http-microservice/product-api/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

/*func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	*/

// Get or Fetch Request 
func (p *Products) GetProductshandle (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Get Product Request") //Get Product returns the proucts from data store.
	lp := data.GetProducts()
	err := lp.ConvToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
// Post or Create Requested 
func (p *Products) PostProductshandle (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Post Product Request") //Post Products handle add products to data store. 

	}
//Put or Update Requested
func (p *Products) UpdateProductshandle (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id:= vars["id"]
	p.l.Println("Update Product Request", id) //Update Products handle to update products to data store. 
	
	lp := &data.Product {}
	err := lp.ConvFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal Id", http.StatusBadRequest)
	}
	data.UpdateProductshandle(lp)
	}