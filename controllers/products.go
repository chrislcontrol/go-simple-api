package controllers

import (
	"github/chrislcontrol/go-simple-api/entities"
	"github/chrislcontrol/go-simple-api/usecase"
)

func CreateProduct(name string, price float64) entities.Product {
	return usecase.CreateProduct(name, price)
}