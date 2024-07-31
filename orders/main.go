package main

import (
	_ "HL_online_shop/docs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Orders API
// @version 1.0
// @description This is an orders API.
// @host localhost:8083
// @BasePath /

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/orders", GetOrders).Methods("GET")
	r.HandleFunc("/orders", CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", GetOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", DeleteOrder).Methods("DELETE")
	r.HandleFunc("/search/orders", SearchOrders).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8083",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Orders service is running on port 8083")
	log.Fatal(srv.ListenAndServe())
}
