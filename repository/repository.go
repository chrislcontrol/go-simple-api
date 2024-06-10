package repository

import (
	"github/chrislcontrol/go-simple-api/db"
	"github/chrislcontrol/go-simple-api/entity"
)

type ProductRepository struct {

}

func (repo ProductRepository) Create (name string, price float64) entity.Product {
	return db.InsertProduct(name, price)
}
