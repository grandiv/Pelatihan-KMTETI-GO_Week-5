package main

import (
	"fmt"
	"net/http"

	handler "Pelatihan-KMTETI-GO_Week-5/api"
)

func main() {
	h := http.NewServeMux()

	s := &http.Server{
		Addr: ":8080",
		Handler: h,
	}

	h.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	h.HandleFunc("/api/products", handler.ProductHandler)

	fmt.Println("HTTP Server running on port 8080")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}