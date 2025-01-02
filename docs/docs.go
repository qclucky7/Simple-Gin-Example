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
        "/v1/hello-word": {
            "get": {
                "description": "Hello World",
                "tags": [
                    "API"
                ],
                "summary": "Hello World",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/models.Result-string"
                        }
                    },
                    "403": {
                        "description": "System error",
                        "schema": {
                            "$ref": "#/definitions/models.Result-string"
                        }
                    },
                    "500": {
                        "description": "System error",
                        "schema": {
                            "$ref": "#/definitions/models.Result-string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Result-string": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "data",
                    "type": "string"
                },
                "status": {
                    "description": "status",
                    "type": "integer"
                }
            }
        }
    },
    "tags": [
        {
            "description": "Web API",
            "name": "API"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "127.0.0.1:9000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Web Server Documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}