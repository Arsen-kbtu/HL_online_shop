// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Check the health of the service",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/payments": {
            "get": {
                "description": "Get all payments",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payments"
                ],
                "summary": "Get all payments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Payment"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new payment using API ePayment.kz",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payments"
                ],
                "summary": "Create a payment",
                "parameters": [
                    {
                        "description": "Create payment",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.PaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.Payment"
                        }
                    }
                }
            }
        },
        "/payments/{id}": {
            "get": {
                "responses": {}
            },
            "put": {
                "description": "Update a payment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payments"
                ],
                "summary": "Update a payment by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Payment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update payment",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Payment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Payment"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a payment by ID",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "payments"
                ],
                "summary": "Delete a payment by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Payment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/search/payments": {
            "get": {
                "description": "Search payments by user, order, or status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payments"
                ],
                "summary": "Search payments by user, order, or status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Payment Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Payment"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Payment": {
            "type": "object",
            "required": [
                "amount",
                "order_id",
                "status",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 100
                },
                "id": {
                    "type": "integer",
                    "readOnly": true,
                    "example": 1
                },
                "order_id": {
                    "type": "integer",
                    "example": 1
                },
                "payment_date": {
                    "type": "string",
                    "readOnly": true,
                    "example": "2023-07-20T15:04:05Z"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "successful",
                        "unsuccessful"
                    ],
                    "example": "successful"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "main.PaymentRequest": {
            "type": "object",
            "required": [
                "amount",
                "cvc",
                "expDate",
                "hpan",
                "order_id",
                "terminalId",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 100
                },
                "cvc": {
                    "type": "string",
                    "example": "636"
                },
                "expDate": {
                    "type": "string",
                    "example": "1022"
                },
                "hpan": {
                    "type": "string",
                    "example": "4003032704547597"
                },
                "order_id": {
                    "type": "integer",
                    "example": 1
                },
                "terminalId": {
                    "type": "string",
                    "example": "67e34d63-102f-4bd1-898e-370781d0074d"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8084",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Payments API",
	Description:      "This is a payments API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
