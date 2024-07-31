package main

import (
	_ "HL_online_shop/docs"
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
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
	//InitDB()

	var cfg Config

	url := os.Getenv("DATABASE_URL")
	flag.IntVar(&cfg.Port, "port", 8081, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.Db.Dsn, "db-dsn", url, "PostgreSQL DSN")
	flag.Parse()
	_, err := OpenDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	} else {
		log.Println("Connected to the database")
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/search/users", SearchUsers).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Разрешить все домены
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)(r)

	srv := &http.Server{
		Handler:      corsHandler,
		Addr:         "0.0.0.0:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Users service is running on port 8081")
	log.Fatal(srv.ListenAndServe())
}
