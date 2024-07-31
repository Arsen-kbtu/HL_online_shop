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

// GetPayments godoc
// @Summary Get all payments
// @Description Get all payments
// @Tags payments
// @Produce json
// @Success 200 {array} Payment
// @Router /payments [get]
func GetPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := GetAllPaymentsRepo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
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
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := validate.Struct(paymentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получение токена
	token, err := getToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	publicKey, err := getRSAPublicKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cardData := CardData{
		HPAN:       paymentRequest.HPAN,
		ExpDate:    paymentRequest.ExpDate,
		CVC:        paymentRequest.CVC,
		TerminalID: paymentRequest.TerminalID,
	}

	cryptogram, err := createCryptogram(cardData, publicKey)

	// Выполнение платежа
	paymentResponse, err := makePayment(token, cryptogram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Сохранение информации о платеже в базу данных
	payment := &Payment{
		Amount:      paymentRequest.Amount,
		OrderID:     paymentRequest.OrderID,
		Status:      paymentResponse.Status,
		UserID:      paymentRequest.UserID,
		PaymentDate: time.Now(),
	}
	if payment.Status == "" {
		payment.Status = "failure"

	}

	if err := CreatePaymentRepo(payment); err != nil {
		http.Error(w, "Failed to save payment", http.StatusInternalServerError)
		return
	}

	// Возврат успешного ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payment)
}

// GetPayment godoc
// @Summary Get a payment by ID
// @Description Get a payment by ID
// @Tags payments
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} Payment
// @Router /payments/{id} [get]
func GetPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	payment, err := GetPaymentByIDRepo(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Payment not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(payment)
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
func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validate.Struct(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payment.ID = id
	if err := UpdatePaymentRepo(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

// DeletePayment godoc
// @Summary Delete a payment by ID
// @Description Delete a payment by ID
// @Tags payments
// @Produce plain
// @Param id path int true "Payment ID"
// @Success 200 {string} string "Deleted"
// @Router /payments/{id} [delete]
func DeletePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	if err := DeletePaymentRepo(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
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
func SearchPayments(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user")
	orderIDStr := r.URL.Query().Get("order")
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

	var orderID uint
	if orderIDStr != "" {
		parsedOrderID, err := strconv.Atoi(orderIDStr)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}
		orderID = uint(parsedOrderID)
	}

	payments, err := SearchPaymentsRepo(userID, orderID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
	w.WriteHeader(http.StatusOK)
}
