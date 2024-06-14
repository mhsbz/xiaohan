// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "xiao han services",
    "title": "XiaohanService",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/xiaohan/v1",
  "paths": {
    "/action": {
      "get": {
        "description": "Action",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Action",
        "operationId": "action",
        "parameters": [
          {
            "type": "string",
            "description": "Member ID",
            "name": "member_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Action",
            "name": "action",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "description": "Health check",
        "summary": "Health check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "status": {
                  "type": "string",
                  "example": "OK"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "xiao han services",
    "title": "XiaohanService",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/xiaohan/v1",
  "paths": {
    "/action": {
      "get": {
        "description": "Action",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Action",
        "operationId": "action",
        "parameters": [
          {
            "type": "string",
            "description": "Member ID",
            "name": "member_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Action",
            "name": "action",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "description": "Health check",
        "summary": "Health check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "status": {
                  "type": "string",
                  "example": "OK"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
}