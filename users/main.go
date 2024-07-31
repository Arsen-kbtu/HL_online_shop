package main

import (
	_ "HL_online_shop/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"time"
)

// @title Users API
// @version 1.0
// @description This is a users API.
// @host localhost:8081
// @BasePath /

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/search/users", SearchUsers).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Users service is running on port 8081")
	log.Fatal(srv.ListenAndServe())
}
