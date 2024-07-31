package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

type PaymentRequest struct {
	Amount     float64 `json:"amount" validate:"required" example:"100.00"`
	OrderID    int     `json:"order_id" validate:"required" example:"1"`
	UserID     int     `json:"user_id" validate:"required" example:"1"`
	HPAN       string  `json:"hpan" validate:"required" example:"4003032704547597"`
	ExpDate    string  `json:"expDate" validate:"required" example:"1022"`
	CVC        string  `json:"cvc" validate:"required" example:"636"`
	TerminalID string  `json:"terminalId" validate:"required" example:"67e34d63-102f-4bd1-898e-370781d0074d"`
}

type Payment struct {
	ID          int       `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	UserID      int       `json:"user_id" validate:"required" example:"1"`
	OrderID     int       `json:"order_id" validate:"required" example:"1"`
	Amount      float64   `json:"amount" validate:"required,gt=0" example:"100"`
	PaymentDate time.Time `json:"payment_date" readonly:"true" example:"2023-07-20T15:04:05Z"`
	Status      string    `json:"status" validate:"required,oneof=successful unsuccessful" example:"successful"`
}

func (Payment) TableName() string {
	return "payments_shop"
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func getToken() (string, error) {
	tokenURL := "https://testoauth.homebank.kz/epay2/oauth2/token"

	// Создаем буфер для тела запроса и writer для multipart формы
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Добавляем поля формы
	writer.WriteField("grant_type", "client_credentials")
	writer.WriteField("scope", "webapi usermanagement email_send verification statement statistics payment")
	writer.WriteField("client_id", "test")
	writer.WriteField("client_secret", "yF587AV9Ms94qN2QShFzVR3vFnWkhjbAK3sG")
	writer.WriteField("invoiceId", "000000001")
	writer.WriteField("amount", "100")
	writer.WriteField("currency", "KZT")
	writer.WriteField("terminalId", "67e34d63-102f-4bd1-898e-370781d0074d")

	// Закрываем writer чтобы отправить все данные
	err := writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", tokenURL, body)
	if err != nil {
		return "", err
	}

	// Устанавливаем заголовок Content-Type включая boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status: %s, body: %s ", resp.Status, respBody)
	}

	var token TokenResponse
	err = json.Unmarshal(respBody, &token)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

type PaymentRequestMake struct {
	Amount          int    `json:"amount"`
	Currency        string `json:"currency"`
	Name            string `json:"name"`
	Cryptogram      string `json:"cryptogram"`
	InvoiceID       string `json:"invoiceId"`
	InvoiceIDAlt    string `json:"invoiceIdAlt,omitempty"`
	Description     string `json:"description"`
	AccountID       string `json:"accountId,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone,omitempty"`
	PostLink        string `json:"postLink"`
	FailurePostLink string `json:"failurePostLink,omitempty"`
	CardSave        bool   `json:"cardSave"`
	Data            string `json:"data,omitempty"`
}

type PaymentResponse struct {
	ID           string `json:"id"`
	Amount       int    `json:"amount"`
	Currency     string `json:"currency"`
	InvoiceID    string `json:"invoiceID"`
	AccountID    string `json:"accountID"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Description  string `json:"description"`
	Reference    string `json:"reference"`
	IntReference string `json:"intReference"`
	Secure3D     string `json:"secure3D"`
	CardID       string `json:"cardID"`
	Fee          int    `json:"fee"`
	Code         int    `json:"code"`
	Status       string `json:"status"`
}

func makePayment(token, cryptogram string) (*PaymentResponse, error) {
	url := "https://testepay.homebank.kz/api/payment/cryptopay"

	paymentRequestMake := PaymentRequestMake{
		Amount:          100,
		Currency:        "KZT",
		Name:            "JON JONSON",
		Cryptogram:      cryptogram,
		InvoiceID:       "000001",
		InvoiceIDAlt:    "8564546",
		Description:     "test payment",
		AccountID:       "uuid000001",
		Email:           "jj@example.com",
		Phone:           "77777777777",
		PostLink:        "https://testmerchant/order/1123",
		FailurePostLink: "https://testmerchant/order/1123/fail",
		CardSave:        true,
		Data:            "{\"statement\":{\"name\":\"Arman Ali\",\"invoiceID\":\"80000016\"}}",
	}

	jsonData, err := json.Marshal(paymentRequestMake)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payment request: %w", err)
	}

	//fmt.Println("Sending data to server:", string(jsonData)) // Отладка отправляемых данных

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	//if resp.StatusCode != http.StatusOK {
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	return nil, fmt.Errorf("server returned non-200 status: %d, body: %s", resp.StatusCode, string(body))
	//}

	var paymentResponse PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &paymentResponse, nil
}

func getRSAPublicKey() (*rsa.PublicKey, error) {
	resp, err := http.Get("https://testepay.homebank.kz/api/public.rsa")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pemData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}

	return publicKey, nil
}

// "hpan":"4003032704547597","expDate":"1022","cvc":"636","terminalId":"67e34d63-102f-4bd1-898e-370781d0074d"
type CardData struct {
	HPAN       string `json:"hpan" validate:"required" example:"4003032704547597"`
	ExpDate    string `json:"expDate" validate:"required" example:"1022"`
	CVC        string `json:"cvc" validate:"required" example:"636"`
	TerminalID string `json:"terminalId" validate:"required" example:"67e34d63-102f-4bd1-898e-370781d0074d"`
}

func createCryptogram(cardData CardData, publicKey *rsa.PublicKey) (string, error) {
	cardDataJSON, err := json.Marshal(cardData)
	if err != nil {
		return "", err
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, cardDataJSON)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedData), nil
}
