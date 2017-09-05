package dlframework

const (
	dlframework_swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "dlframework.proto",
    "version": "1.0.0",
    "contact": {
      "name": "Abdul Dakkak, Cheng Li",
      "url": "https://github.com/rai-project/carml"
    }
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/predict/dataset": {
      "post": {
        "summary": "Dataset method receives a single dataset and runs\nthe predictor on all elements of the dataset.",
        "description": "The result is a prediction feature stream.",
        "operationId": "Dataset",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictionFeatureResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictDatasetRequest"
            }
          }
        ],
        "tags": [
          "Predictor"
        ]
      }
    },
    "/v1/predict/images": {
      "post": {
        "summary": "Image method receives a list base64 encoded images and runs\nthe predictor on all the images.",
        "description": "The result is a prediction feature stream for each image.",
        "operationId": "Images",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictionFeatureResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictImagesRequest"
            }
          }
        ],
        "tags": [
          "Predictor"
        ]
      }
    },
    "/v1/predict/urls": {
      "post": {
        "summary": "Image method receives a stream of urls and runs\nthe predictor on all the urls. The",
        "description": "The result is a prediction feature stream for each url.",
        "operationId": "URLs",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictionFeatureResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictURLsRequest"
            }
          }
        ],
        "tags": [
          "Predictor"
        ]
      }
    },
    "/v1/registry/frameworks/agent": {
      "get": {
        "operationId": "FrameworkAgents",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkAgents"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "framework_version",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/registry/frameworks/manifest": {
      "get": {
        "operationId": "FrameworkManifests",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkFrameworkManifestsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "framework_version",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/registry/models/agent": {
      "get": {
        "operationId": "ModelAgents",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkAgents"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "framework_version",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "model_name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "model_version",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    },
    "/v1/registry/models/manifest": {
      "get": {
        "operationId": "ModelManifests",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/dlframeworkModelManifestsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "framework_name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "framework_version",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "model_name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "model_version",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Registry"
        ]
      }
    }
  },
  "definitions": {
    "ModelManifestModel": {
      "type": "object",
      "properties": {
        "base_url": {
          "type": "string"
        },
        "weights_path": {
          "type": "string"
        },
        "graph_path": {
          "type": "string"
        },
        "is_archive": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "PredictDatasetRequestDataset": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "PredictImagesRequestImage": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "An id used to identify the output feature: maps to input_id for output"
        },
        "image": {
          "type": "string",
          "format": "byte",
          "title": "The image is base64 encoded"
        },
        "preprocessed": {
          "type": "boolean",
          "format": "boolean",
          "title": "Preprocessed is set to true to disable preprocessing.\nIf enabled then the image is assumed to be rescaled and\nencoded as an array of float32 values"
        }
      }
    },
    "PredictURLsRequestURL": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "An id used to identify the output feature: maps to input_id for output"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "TypeParameter": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    },
    "dlframeworkAgent": {
      "type": "object",
      "properties": {
        "host": {
          "type": "string"
        },
        "port": {
          "type": "string"
        }
      }
    },
    "dlframeworkAgents": {
      "type": "object",
      "properties": {
        "agents": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkAgent"
          }
        }
      }
    },
    "dlframeworkContainerHardware": {
      "type": "object",
      "properties": {
        "gpu": {
          "type": "string"
        },
        "cpu": {
          "type": "string"
        }
      }
    },
    "dlframeworkFrameworkManifest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "container": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/dlframeworkContainerHardware"
          }
        }
      }
    },
    "dlframeworkFrameworkManifestsResponse": {
      "type": "object",
      "properties": {
        "manifests": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkFrameworkManifest"
          }
        }
      }
    },
    "dlframeworkModelManifest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "framework": {
          "$ref": "#/definitions/dlframeworkFrameworkManifest"
        },
        "container": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/dlframeworkContainerHardware"
          }
        },
        "description": {
          "type": "string"
        },
        "reference": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "license": {
          "type": "string"
        },
        "inputs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkModelManifestType"
          }
        },
        "output": {
          "$ref": "#/definitions/dlframeworkModelManifestType"
        },
        "before_preprocess": {
          "type": "string"
        },
        "preprocess": {
          "type": "string"
        },
        "after_preprocess": {
          "type": "string"
        },
        "before_postprocess": {
          "type": "string"
        },
        "postprocess": {
          "type": "string"
        },
        "after_postprocess": {
          "type": "string"
        },
        "model": {
          "$ref": "#/definitions/ModelManifestModel"
        },
        "attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "hidden": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "dlframeworkModelManifestType": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "parameters": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/TypeParameter"
          }
        }
      }
    },
    "dlframeworkModelManifestsResponse": {
      "type": "object",
      "properties": {
        "manifests": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkModelManifest"
          }
        }
      }
    },
    "dlframeworkPredictDatasetRequest": {
      "type": "object",
      "properties": {
        "dataset": {
          "$ref": "#/definitions/PredictDatasetRequestDataset"
        },
        "options": {
          "$ref": "#/definitions/dlframeworkPredictionOptions"
        }
      }
    },
    "dlframeworkPredictImagesRequest": {
      "type": "object",
      "properties": {
        "images": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PredictImagesRequestImage"
          },
          "title": "A list of Base64 encoded images"
        },
        "options": {
          "$ref": "#/definitions/dlframeworkPredictionOptions"
        }
      }
    },
    "dlframeworkPredictURLsRequest": {
      "type": "object",
      "properties": {
        "urls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PredictURLsRequestURL"
          }
        },
        "options": {
          "$ref": "#/definitions/dlframeworkPredictionOptions"
        }
      }
    },
    "dlframeworkPredictionFeature": {
      "type": "object",
      "properties": {
        "index": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "probability": {
          "type": "number",
          "format": "float"
        }
      }
    },
    "dlframeworkPredictionFeatureResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        },
        "input_id": {
          "type": "string"
        },
        "feature": {
          "$ref": "#/definitions/dlframeworkPredictionFeature"
        }
      }
    },
    "dlframeworkPredictionOptions": {
      "type": "object",
      "properties": {
        "request_id": {
          "type": "string"
        },
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        },
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        },
        "feature_limit": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  },
  "host": "localhost",
  "externalDocs": {
    "url": "https://rai-project.github.io/carml"
  }
}
`
	swagger_info = `{
	"info": {
		"version": "1.0.0",
		"contact": {
			"name": "Abdul Dakkak, Cheng Li",
			"url": "https://github.com/rai-project/carml"
		}
	},
	"host": "localhost",
	"externalDocs": {
		"url": "https://rai-project.github.io/carml"
	}
}
`
)
