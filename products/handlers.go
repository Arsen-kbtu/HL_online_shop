package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := GetAllProductsRepo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
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
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := validate.Struct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	if err := CreateProductRepo(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Router /products/{id} [get]
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := GetProductByIDRepo(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	json.NewEncoder(w).Encode(product)
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
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validate.Struct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	product.ID = uint(id)
	if err := UpdateProductRepo(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a product by ID
// @Tags products
// @Produce plain
// @Param id path int true "Product ID"
// @Success 200 {string} string "Deleted"
// @Router /products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := DeleteProductRepo(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
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
func SearchProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")

	products, err := SearchProductsRepo(name, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	json.NewEncoder(w).Encode(products)
	w.WriteHeader(http.StatusOK)
}
