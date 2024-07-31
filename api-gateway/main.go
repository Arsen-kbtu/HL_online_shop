package main

import (
	_ "HL_online_shop/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func main() {

	//Init()
	r := mux.NewRouter()

	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Пример маршрутов для пользователей
	r.HandleFunc("/users", handleUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handleUserByID).Methods("GET")
	r.HandleFunc("/users", handleCreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handleUpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handleDeleteUser).Methods("DELETE")
	r.HandleFunc("/search/users", handleSearchUsers).Methods("GET")

	// Пример маршрутов для товаров
	r.HandleFunc("/products", handleProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handleProductByID).Methods("GET")
	r.HandleFunc("/products", handleCreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handleUpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handleDeleteProduct).Methods("DELETE")
	r.HandleFunc("/search/products", handleSearchProducts).Methods("GET")

	r.HandleFunc("/orders", handleOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", handleOrderByID).Methods("GET")
	r.HandleFunc("/orders", handleCreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", handleUpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", handleDeleteOrder).Methods("DELETE")
	r.HandleFunc("/search/orders", handleSearchOrders).Methods("GET")

	r.HandleFunc("/payments", handlePayments).Methods("GET")
	r.HandleFunc("/payments/{id}", handlePaymentByID).Methods("GET")
	r.HandleFunc("/payments", handleCreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", handleUpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", handleDeletePayment).Methods("DELETE")
	r.HandleFunc("/search/payments", handleSearchPayments).Methods("GET")

	// Запуск сервера
	log.Println("API Gateway is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
