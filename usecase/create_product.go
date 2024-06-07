package usecase

import "github/chrislcontrol/go-simple-api/entities"

func CreateProduct(name string, price float64) entities.Product {
	return entities.Product{
		Name:  name,
		Price: price,
	}
}