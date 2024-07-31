package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

var validate = *validator.New()

// HealthCheck godoc
// @Summary Health Check
// @Description Check the health of the service
// @Tags health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Produce json
// @Success 200 {array} Order
// @Router /orders [get]
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := GetAllOrdersRepo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

// CreateOrder godoc
// @Summary Create an order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body Order true "Create order"
// @Success 201 {object} Order
// @Router /orders [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := validate.Struct(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if order.OrderDate.IsZero() {
		order.OrderDate = time.Now()

	}
	if err := CreateOrderRepo(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrder godoc
// @Summary Get an order by ID
// @Description Get an order by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} Order
// @Router /orders/{id} [get]
func GetOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := GetOrderByIDRepo(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(order)
}

// UpdateOrder godoc
// @Summary Update an order by ID
// @Description Update an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body Order true "Update order"
// @Success 200 {object} Order
// @Router /orders/{id} [put]
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validate.Struct(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order.ID = uint(id)
	if err := UpdateOrderRepo(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(order)
}

// DeleteOrder godoc
// @Summary Delete an order by ID
// @Description Delete an order by ID
// @Tags orders
// @Produce plain
// @Param id path int true "Order ID"
// @Success 200 {string} string "Deleted"
// @Router /orders/{id} [delete]
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	if err := DeleteOrderRepo(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

// SearchOrders godoc
// @Summary Search orders by user or status
// @Description Search orders by user or status
// @Tags orders
// @Produce json
// @Param user query int false "User ID"
// @Param status query string false "Order Status"
// @Success 200 {array} Order
// @Router /search/orders [get]
func SearchOrders(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user")
	status := r.URL.Query().Get("status")

	var userID uint
	if userIDStr != "" {
		parsedUserID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		userID = uint(parsedUserID)
	}

	orders, err := SearchOrdersRepo(userID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
	w.WriteHeader(http.StatusOK)
}
