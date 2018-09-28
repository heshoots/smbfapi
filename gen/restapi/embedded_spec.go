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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple API",
    "title": "smbf API",
    "contact": {
      "email": "heshoots9999@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "brackets.superminerbattle.farm",
  "basePath": "/api/1.0.0",
  "paths": {
    "/events": {
      "get": {
        "description": "Get list of events\n",
        "produces": [
          "application/json"
        ],
        "summary": "get multiple events",
        "operationId": "searchEvents",
        "parameters": [
          {
            "type": "string",
            "description": "Facebook api key",
            "name": "api_key",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    },
    "/events/{id}": {
      "get": {
        "description": "search for event with challonge tournaments\n",
        "produces": [
          "application/json"
        ],
        "summary": "get specific event",
        "operationId": "getEvent",
        "parameters": [
          {
            "type": "string",
            "description": "Facebook api key",
            "name": "api_key",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "id of event",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "$ref": "#/definitions/Event"
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    },
    "/tournaments": {
      "get": {
        "description": "get tournaments from api\n",
        "produces": [
          "application/json"
        ],
        "summary": "get multiple tournament",
        "operationId": "searchTournaments",
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tournament"
              }
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    },
    "/tournaments/{id}": {
      "get": {
        "description": "Access values for a tournament by id\n",
        "produces": [
          "application/json"
        ],
        "summary": "get specific tournament",
        "operationId": "getTournament",
        "parameters": [
          {
            "type": "string",
            "description": "id of event",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returning tournament",
            "schema": {
              "$ref": "#/definitions/Tournament"
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    }
  },
  "definitions": {
    "Event": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "attending_count": {
          "type": "integer",
          "example": 12
        },
        "end_time": {
          "type": "string",
          "format": "int32",
          "example": "2016-08-29T09:12:33.001Z"
        },
        "id": {
          "type": "string",
          "format": "int32",
          "example": 476564289474918
        },
        "name": {
          "type": "string",
          "example": "Widget Adapter"
        },
        "start_time": {
          "type": "string",
          "format": "int32",
          "example": "2016-08-29T09:12:33.001Z"
        },
        "tournaments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tournament"
          }
        }
      }
    },
    "Tournament": {
      "required": [
        "id",
        "name",
        "full_challonge_url"
      ],
      "properties": {
        "created_at": {
          "type": "string",
          "format": "int32",
          "example": "2018-09-09T12:28:16.031Z"
        },
        "full_challonge_url": {
          "type": "string",
          "format": "url",
          "example": "https://www.acme-corp.com"
        },
        "game_name": {
          "type": "string",
          "example": "Street Fighter V"
        },
        "id": {
          "type": "integer",
          "format": "int32",
          "example": 8342251
        },
        "name": {
          "type": "string",
          "example": "ww345SFV"
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
    "description": "This is a simple API",
    "title": "smbf API",
    "contact": {
      "email": "heshoots9999@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "brackets.superminerbattle.farm",
  "basePath": "/api/1.0.0",
  "paths": {
    "/events": {
      "get": {
        "description": "Get list of events\n",
        "produces": [
          "application/json"
        ],
        "summary": "get multiple events",
        "operationId": "searchEvents",
        "parameters": [
          {
            "type": "string",
            "description": "Facebook api key",
            "name": "api_key",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    },
    "/events/{id}": {
      "get": {
        "description": "search for event with challonge tournaments\n",
        "produces": [
          "application/json"
        ],
        "summary": "get specific event",
        "operationId": "getEvent",
        "parameters": [
          {
            "type": "string",
            "description": "Facebook api key",
            "name": "api_key",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "id of event",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "$ref": "#/definitions/Event"
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    },
    "/tournaments": {
      "get": {
        "description": "get tournaments from api\n",
        "produces": [
          "application/json"
        ],
        "summary": "get multiple tournament",
        "operationId": "searchTournaments",
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tournament"
              }
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    },
    "/tournaments/{id}": {
      "get": {
        "description": "Access values for a tournament by id\n",
        "produces": [
          "application/json"
        ],
        "summary": "get specific tournament",
        "operationId": "getTournament",
        "parameters": [
          {
            "type": "string",
            "description": "id of event",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returning tournament",
            "schema": {
              "$ref": "#/definitions/Tournament"
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      }
    }
  },
  "definitions": {
    "Event": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "attending_count": {
          "type": "integer",
          "example": 12
        },
        "end_time": {
          "type": "string",
          "format": "int32",
          "example": "2016-08-29T09:12:33.001Z"
        },
        "id": {
          "type": "string",
          "format": "int32",
          "example": 476564289474918
        },
        "name": {
          "type": "string",
          "example": "Widget Adapter"
        },
        "start_time": {
          "type": "string",
          "format": "int32",
          "example": "2016-08-29T09:12:33.001Z"
        },
        "tournaments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tournament"
          }
        }
      }
    },
    "Tournament": {
      "required": [
        "id",
        "name",
        "full_challonge_url"
      ],
      "properties": {
        "created_at": {
          "type": "string",
          "format": "int32",
          "example": "2018-09-09T12:28:16.031Z"
        },
        "full_challonge_url": {
          "type": "string",
          "format": "url",
          "example": "https://www.acme-corp.com"
        },
        "game_name": {
          "type": "string",
          "example": "Street Fighter V"
        },
        "id": {
          "type": "integer",
          "format": "int32",
          "example": 8342251
        },
        "name": {
          "type": "string",
          "example": "ww345SFV"
        }
      }
    }
  }
}`))
}
