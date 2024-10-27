package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"week5/src/service"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": 
		data, err := service.GetAllProduct()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	case "POST":
		err := service.CreateProduct(r.Body)
		if err != nil {
			if err.Error() == "Bad Request" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Product Created Successfully")
		return
	default:
		log.Default().Println(http.StatusMethodNotAllowed)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}