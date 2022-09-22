// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/dashboard/menu": {
            "get": {
                "security": [
                    {
                        "Auth0Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Get the menu for a store",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Menu"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/dashboard/orders": {
            "get": {
                "security": [
                    {
                        "Auth0Token": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Get all orders from a store",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GetOrdersResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "security": [
                    {
                        "FirebaseToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get all orders from a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GetOrdersResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "FirebaseToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create Order for a store",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/CreateOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/store/{id}": {
            "get": {
                "security": [
                    {
                        "FirebaseToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "Get menu for a given store",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the store",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetStoreResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/stores": {
            "get": {
                "security": [
                    {
                        "FirebaseToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stores"
                ],
                "summary": "Get all open stores",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GetStoresOverviewResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "FirebaseToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get User associated with the given Firebase Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "FirebaseToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create User associated with the given Firebase Token",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/whitelist": {
            "post": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "whitelist"
                ],
                "summary": "Check if identifier is whitelisted",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/WhitelistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateOrderRequest": {
            "type": "object",
            "properties": {
                "orderItems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/OrderItem"
                    }
                },
                "storeId": {
                    "type": "string"
                }
            }
        },
        "CreateOrderResponse": {
            "type": "object",
            "properties": {
                "stripeClientSecret": {
                    "type": "string"
                }
            }
        },
        "CreateUserRequest": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "language_code": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "GetOrdersResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "estimatedPickupTime": {
                    "type": "string"
                },
                "googleMapsLink": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isPaid": {
                    "type": "boolean"
                },
                "orderItems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/OrderItem"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "storeAddress": {
                    "type": "string"
                },
                "storeDescription": {
                    "type": "string"
                },
                "storeImageUrl": {
                    "type": "string"
                },
                "storeName": {
                    "type": "string"
                },
                "storePhoneNumber": {
                    "type": "string"
                }
            }
        },
        "GetStoreCategory": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GetStoreProduct"
                    }
                },
                "sortOrder": {
                    "type": "integer"
                }
            }
        },
        "GetStoreOpeningHour": {
            "type": "object",
            "properties": {
                "dayOfWeek": {
                    "type": "integer"
                },
                "endTimestamp": {
                    "type": "integer"
                },
                "startTimestamp": {
                    "type": "integer"
                }
            }
        },
        "GetStoreProduct": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "max": {
                    "type": "integer"
                },
                "min": {
                    "type": "integer"
                },
                "multiMax": {
                    "type": "integer"
                },
                "multiply": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "plu": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "productType": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GetStoreProduct"
                    }
                },
                "snoozed": {
                    "type": "boolean"
                },
                "sortOrder": {
                    "type": "integer"
                },
                "tax": {
                    "type": "integer"
                },
                "visible": {
                    "type": "boolean"
                }
            }
        },
        "GetStoreResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "averagePickupTime": {
                    "type": "integer"
                },
                "averageReview": {
                    "type": "number"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GetStoreCategory"
                    }
                },
                "description": {
                    "type": "string"
                },
                "googleMapsLink": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "openingHours": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GetStoreOpeningHour"
                    }
                },
                "phoneNumber": {
                    "type": "string"
                },
                "reviewCount": {
                    "type": "integer"
                }
            }
        },
        "GetStoresOverviewResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "averagePickupTime": {
                    "type": "integer"
                },
                "averageReview": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "googleMapsLink": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "isAvailable": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "reviewCount": {
                    "type": "integer"
                }
            }
        },
        "Menu": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GetStoreCategory"
                    }
                },
                "opening_hours": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GetStoreOpeningHour"
                    }
                }
            }
        },
        "OrderItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "plu": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "subItems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/OrderItem"
                    }
                }
            }
        },
        "User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "firebaseId": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "languageCode": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "WhitelistRequest": {
            "type": "object",
            "properties": {
                "identifier": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Auth0Token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "FirebaseToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
