// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for the WhiteBoard app developed by Weberster Inc",
    "title": "WhiteBoard API",
    "contact": {
      "email": "webersterinc@gmail.com"
    },
    "version": "0.1.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/api/v1",
  "paths": {
    "/boards/{boardName}/fingerpaths": {
      "get": {
        "description": "Gets all fingerpaths associated with the given board.\n",
        "produces": [
          "application/json"
        ],
        "operationId": "fingerPathsGet",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the board",
            "name": "boardName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "fingerpaths associated with board",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/FingerPath"
              }
            }
          },
          "400": {
            "description": "no associated board",
            "schema": {}
          }
        }
      },
      "post": {
        "description": "Adds fingerpaths to the associated board.",
        "produces": [
          "application/json"
        ],
        "operationId": "fingerPathsPost",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the board",
            "name": "boardName",
            "in": "path",
            "required": true
          },
          {
            "description": "Inventory item to add",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/FingerPath"
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "fingerpaths added"
          },
          "400": {
            "description": "no associated board",
            "schema": {}
          }
        }
      }
    }
  },
  "definitions": {
    "FingerPath": {
      "type": "object",
      "title": "FingerPath",
      "required": [
        "pathColor",
        "boardColor",
        "dash",
        "blur",
        "clear",
        "fingerPoints",
        "strokeWidth"
      ],
      "properties": {
        "blur": {
          "type": "boolean"
        },
        "boardColor": {
          "type": "integer",
          "format": "uint32"
        },
        "clear": {
          "type": "boolean"
        },
        "dash": {
          "type": "boolean"
        },
        "fingerPoints": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/FingerPoint"
          }
        },
        "pathColor": {
          "type": "integer",
          "format": "uint32"
        },
        "pathId": {
          "type": "integer",
          "format": "int32"
        },
        "strokeWidth": {
          "type": "integer",
          "format": "int32"
        },
        "userId": {
          "type": "string"
        }
      }
    },
    "FingerPoint": {
      "type": "object",
      "title": "FingerPoint",
      "required": [
        "x",
        "y"
      ],
      "properties": {
        "x": {
          "type": "integer",
          "format": "int32"
        },
        "y": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for the WhiteBoard app developed by Weberster Inc",
    "title": "WhiteBoard API",
    "contact": {
      "email": "webersterinc@gmail.com"
    },
    "version": "0.1.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/api/v1",
  "paths": {
    "/boards/{boardName}/fingerpaths": {
      "get": {
        "description": "Gets all fingerpaths associated with the given board.\n",
        "produces": [
          "application/json"
        ],
        "operationId": "fingerPathsGet",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the board",
            "name": "boardName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "fingerpaths associated with board",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/FingerPath"
              }
            }
          },
          "400": {
            "description": "no associated board",
            "schema": {}
          }
        }
      },
      "post": {
        "description": "Adds fingerpaths to the associated board.",
        "produces": [
          "application/json"
        ],
        "operationId": "fingerPathsPost",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the board",
            "name": "boardName",
            "in": "path",
            "required": true
          },
          {
            "description": "Inventory item to add",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/FingerPath"
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "fingerpaths added"
          },
          "400": {
            "description": "no associated board",
            "schema": {}
          }
        }
      }
    }
  },
  "definitions": {
    "FingerPath": {
      "type": "object",
      "title": "FingerPath",
      "required": [
        "pathColor",
        "boardColor",
        "dash",
        "blur",
        "clear",
        "fingerPoints",
        "strokeWidth"
      ],
      "properties": {
        "blur": {
          "type": "boolean"
        },
        "boardColor": {
          "type": "integer",
          "format": "uint32"
        },
        "clear": {
          "type": "boolean"
        },
        "dash": {
          "type": "boolean"
        },
        "fingerPoints": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/FingerPoint"
          }
        },
        "pathColor": {
          "type": "integer",
          "format": "uint32"
        },
        "pathId": {
          "type": "integer",
          "format": "int32"
        },
        "strokeWidth": {
          "type": "integer",
          "format": "int32"
        },
        "userId": {
          "type": "string"
        }
      }
    },
    "FingerPoint": {
      "type": "object",
      "title": "FingerPoint",
      "required": [
        "x",
        "y"
      ],
      "properties": {
        "x": {
          "type": "integer",
          "format": "int32"
        },
        "y": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
}
