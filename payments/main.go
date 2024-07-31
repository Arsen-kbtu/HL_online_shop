package main

import (
	_ "HL_online_shop/docs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Payments API
// @version 1.0
// @description This is a payments API.
// @host localhost:8084
// @BasePath /

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/payments", GetPayments).Methods("GET")
	r.HandleFunc("/payments", CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", GetPayment).Methods("GET")
	r.HandleFunc("/payments/{id}", UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", DeletePayment).Methods("DELETE")
	r.HandleFunc("/search/payments", SearchPayments).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8084",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Payments service is running on port 8084")
	log.Fatal(srv.ListenAndServe())
}
