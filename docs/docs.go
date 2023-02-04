// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "cars"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/car/{id}": {
            "get": {
                "description": "Reads a single car and returns it.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "read"
                ],
                "summary": "Get car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/constants.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/constants.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/constants.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cars": {
            "get": {
                "description": "Reads and returns all the cars.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "read"
                ],
                "summary": "GetCar all cars",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/constants.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/constants.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/create": {
            "post": {
                "description": "Creates a new car.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "write"
                ],
                "summary": "Creates car",
                "parameters": [
                    {
                        "description": "New car",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/constants.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/constants.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/constants.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "This endpoint will return a status to determine if the service is live or requires a restart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Check"
                ],
                "summary": "The liveness endpoint determines the LIVE status of the service",
                "operationId": "liveliness",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HealthResponse"
                        }
                    }
                }
            }
        },
        "/update": {
            "put": {
                "description": "Updates a new car.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "write"
                ],
                "summary": "Update car",
                "parameters": [
                    {
                        "description": "New car",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/constants.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/constants.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/constants.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "constants.ErrorResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "constants.UserResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Car": {
            "type": "object",
            "properties": {
                "Category": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "make": {
                    "type": "string"
                },
                "mileage": {
                    "type": "integer"
                },
                "model": {
                    "type": "string"
                },
                "package": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.HealthResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GetCars CarsService",
	Description:      "This is a Goland server that manages cars.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}