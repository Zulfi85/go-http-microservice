package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product defines the structure for a API product
type Product struct {
	ID			int 	`json:"id"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Price		float32	 `json:"price"`
	SKU			string	`json:"sku"`
	CreatedOn	string	`json:"-"`
	UpdatedOn	string	`json:"-"`
	DeletedOn	string	`json:"-"`
}

type Products []*Product
func (p *Products) ConvToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product {
	{
		ID:				1,
		Name: 			"Latte",
		Description:	"Milky Cofee",
		Price:			2.50,
		SKU:			"late1",
		CreatedOn:		time.Now().UTC().String(),
		UpdatedOn:		time.Now().UTC().String(),

	},
	{
		ID:				2,
		Name: 			"Espresso",
		Description:	"Strong Coffee",
		Price:			1.5,
		SKU:			"Esp1",
		CreatedOn:		time.Now().UTC().String(),
		UpdatedOn:		time.Now().UTC().String(),

	},
}