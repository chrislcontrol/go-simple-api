package usecase

import (
	"github/chrislcontrol/go-simple-api/entity"
	"github/chrislcontrol/go-simple-api/repository"
)

func CreateProduct(name string, price float64) entity.Product {
	repo := repository.ProductRepository{}

	return repo.Create(name, price)
}
