package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// @summary Health check
// @tags Health
// @produce plain
// @success 200 {string} string "OK"
// @router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// helper function to handle proxy requests
func proxyRequest(w http.ResponseWriter, r *http.Request, url string) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func handleUsers(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://user-service:8081/users")
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func handleUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://user-service:8081/users/"+id)
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "Create user"
// @Success 201 {object} User
// @Router /users [post]
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://user-service:8081/users")
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "Update user"
// @Success 200 {object} User
// @Router /users/{id} [put]
func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://user-service:8081/users/"+id)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "Deleted"
// @Router /users/{id} [delete]
func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://user-service:8081/users/"+id)
}

// SearchUsers godoc
// @Summary Search users by name or role
// @Description Search users by name or role
// @Tags users
// @Produce json
// @Param name query string false "Name"
// @Param role query string false "Role"
// @Success 200 {array} User
// @Router /search/users [get]
func handleSearchUsers(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	role := r.URL.Query().Get("role")
	proxyRequest(w, r, "http://user-service:8081/search/users?name="+name+"&role="+role)
}

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]
func handleProducts(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://product-service:8082/products")
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Router /products/{id} [get]
func handleProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://product-service:8082/products/"+id)
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body Product true "Create product"
// @Success 201 {object} Product
// @Router /products [post]
func handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://product-service:8082/products")
}

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Update a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body Product true "Update product"
// @Success 200 {object} Product
// @Router /products/{id} [put]
func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://product-service:8082/products/"+id)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a product by ID
// @Tags products
// @Produce plain
// @Param id path int true "Product ID"
// @Success 200 {string} string "Deleted"
// @Router /products/{id} [delete]
func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://product-service:8082/products/"+id)
}

// SearchProducts godoc
// @Summary Search products by name or category
// @Description Search products by name or category
// @Tags products
// @Produce json
// @Param name query string false "Product Name"
// @Param category query string false "Product Category"
// @Success 200 {array} Product
// @Router /search/products [get]
func handleSearchProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")
	proxyRequest(w, r, "http://product-service:8082/search/products?name="+name+"&category="+category)
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Produce json
// @Success 200 {array} Order
// @Router /orders [get]
func handleOrders(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://order-service:8083/orders")
}

// GetOrder godoc
// @Summary Get an order by ID
// @Description Get an order by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} Order
// @Router /orders/{id} [get]
func handleOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://order-service:8083/orders/"+id)
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
func handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://order-service:8083/orders")
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
func handleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://order-service:8083/orders/"+id)
}

// DeleteOrder godoc
// @Summary Delete an order by ID
// @Description Delete an order by ID
// @Tags orders
// @Produce plain
// @Param id path int true "Order ID"
// @Success 200 {string} string "Deleted"
// @Router /orders/{id} [delete]
func handleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://order-service:8083/orders/"+id)
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
func handleSearchOrders(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user")
	status := r.URL.Query().Get("status")
	proxyRequest(w, r, "http://order-service:8083/search/orders?user="+userId+"&status="+status)
}

// GetPayments godoc
// @Summary Get all payments
// @Description Get all payments
// @Tags payments
// @Produce json
// @Success 200 {array} Payment
// @Router /payments [get]
func handlePayments(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://payment-service:8084/payments")
}

// GetPayment godoc
// @Summary Get a payment by ID
// @Description Get a payment by ID
// @Tags payments
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} Payment
// @Router /payments/{id} [get]
func handlePaymentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://payment-service:8084/payments/"+id)
}

// CreatePayment godoc
// @Summary Create a payment
// @Description Create a new payment using API ePayment.kz
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body PaymentRequest true "Create payment"
// @Success 201 {object} Payment
// @Router /payments [post]
func handleCreatePayment(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, "http://payment-service:8084/payments")
}

// UpdatePayment godoc
// @Summary Update a payment by ID
// @Description Update a payment by ID
// @Tags payments
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Param payment body Payment true "Update payment"
// @Success 200 {object} Payment
// @Router /payments/{id} [put]
func handleUpdatePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://payment-service:8084/payments/"+id)
}

// DeletePayment godoc
// @Summary Delete a payment by ID
// @Description Delete a payment by ID
// @Tags payments
// @Produce plain
// @Param id path int true "Payment ID"
// @Success 200 {string} string "Deleted"
// @Router /payments/{id} [delete]
func handleDeletePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	proxyRequest(w, r, "http://payment-service:8084/payments/"+id)
}

// SearchPayments godoc
// @Summary Search payments by user, order, or status
// @Description Search payments by user, order, or status
// @Tags payments
// @Produce json
// @Param user query int false "User ID"
// @Param order query int false "Order ID"
// @Param status query string false "Payment Status"
// @Success 200 {array} Payment
// @Router /search/payments [get]
func handleSearchPayments(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user")
	orderId := r.URL.Query().Get("order")
	proxyRequest(w, r, "http://payment-service:8084/search/payments?user="+userId+"&order="+orderId)
}
