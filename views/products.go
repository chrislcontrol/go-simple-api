package views

import (
	"encoding/json"
	"github/chrislcontrol/go-simple-api/controllers"
	"net/http"
)



type RequestBody struct {
	Name string `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}


func Products(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody RequestBody
	
	err := decoder.Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product := controllers.CreateProduct(requestBody.Name, requestBody.Price)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(product)

	
}
