// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Health checking for the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Health Check",
                "operationId": "HealthCheckHandler",
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
<<<<<<< HEAD
        "/merchant-billpayment-qr": {
=======
        "/qrgen": {
>>>>>>> e62476a (QRGen wiht Golang)
            "post": {
                "description": "For Generate TQR",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
<<<<<<< HEAD
                "summary": "TQR Billpayment generate",
                "operationId": "TQR Merchant Billpayment",
                "parameters": [
                    {
                        "description": "Body for Generate TQR",
                        "name": "request.MerchantBillpaymentRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MerchantBillpaymentRequest"
                        }
                    }
                ],
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
        "/merchant-promptpay-qr": {
            "post": {
                "description": "For Generate TQR",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "TQR Promptpay generate",
                "operationId": "TQR Promptpay Promptpay",
                "parameters": [
                    {
                        "description": "Body for Generate TQR",
                        "name": "request.MerchantPromptpayRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MerchantPromptpayRequest"
=======
                "summary": "TQR generate",
                "operationId": "TQRGen",
                "parameters": [
                    {
                        "description": "Body for Generate TQR",
                        "name": "models.QRGenerateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.QRGenerateRequest"
>>>>>>> e62476a (QRGen wiht Golang)
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
<<<<<<< HEAD
        "request.MerchantBillpaymentRequest": {
=======
        "models.QRGenerateRequest": {
>>>>>>> e62476a (QRGen wiht Golang)
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "reference1": {
                    "type": "string"
                },
                "reference2": {
                    "type": "string"
                }
            }
<<<<<<< HEAD
        },
        "request.MerchantPromptpayRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "onetime": {
                    "type": "boolean"
                },
                "recieveId": {
                    "type": "string"
                }
            }
=======
>>>>>>> e62476a (QRGen wiht Golang)
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
