package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

func GetProducts() []*Product {

	return productList
}

var productList = []*Product{

	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          3,
		Name:        "Cappuccino",
		Description: "Espresso with steamed milk and foamed milk",
		Price:       2.99,
		SKU:         "a3sdf4",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
