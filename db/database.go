package db

import "github/chrislcontrol/go-simple-api/entity"

var products []entity.Product

func InsertProduct(name string, price float64) entity.Product {
	var id int
	if len(products) == 0 {
		id = 1
	} else {
		id = int(products[len(products) - 1].Id) + 1
	}


	product := entity.Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
	products = append(products, product)

	return product
}
