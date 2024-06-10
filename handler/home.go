package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string
}

func Home(w http.ResponseWriter, r *http.Request) {
	response := Response {
		Message: "Golang project, home page.",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}