package main

import (
	_ "HL_online_shop/docs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Products API
// @version 1.0
// @description This is a products API.
// @host localhost:8082
// @BasePath /

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	r.HandleFunc("/search/products", SearchProducts).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8082",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Products service is running on port 8082")
	log.Fatal(srv.ListenAndServe())
}
