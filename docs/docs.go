// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/shorten": {
            "post": {
                "description": "Takes a URL and returns a shortened URL alias and QR code path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL Shortener"
                ],
                "summary": "Shorten a URL",
                "parameters": [
                    {
                        "description": "URL to be shortened",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ShortenURL.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.ShortenURLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.HttpError"
                        }
                    }
                }
            }
        },
        "/{alias}": {
            "get": {
                "description": "Redirects to the original full URL based on the shortened URL alias",
                "tags": [
                    "URL Redirect"
                ],
                "summary": "Redirect to the full URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shortened URL alias",
                        "name": "alias",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.HttpError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.ShortenURL.RequestBody": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "handlers.ShortenURLResponse": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "An application for shortening links.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
