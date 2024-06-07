package entities

type Product struct {
	Name string `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
