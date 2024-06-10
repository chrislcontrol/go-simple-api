package controller

import (
	"github/chrislcontrol/go-simple-api/entity"
	"github/chrislcontrol/go-simple-api/usecase"
)

func CreateProduct(name string, price float64) entity.Product {
	return usecase.CreateProduct(name, price)
}