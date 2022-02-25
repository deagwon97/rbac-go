// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "User 생성",
                "parameters": [
                    {
                        "description": "회원가입 정보",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddUserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResult"
                        }
                    }
                }
            }
        },
        "/account/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "로그인",
                "parameters": [
                    {
                        "description": "로그인 정보",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "access token \u0026 refresh token",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResult"
                        }
                    }
                }
            }
        },
        "/account/renew": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "access token 재발급",
                "parameters": [
                    {
                        "description": "Refresh Token",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.refreshToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "access token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/account/valid": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "access token 인증",
                "parameters": [
                    {
                        "description": "Access Token",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.accessToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "유효성 검증 결과",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/content": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Content 생성",
                "parameters": [
                    {
                        "description": "Content Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ContentData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Content"
                        }
                    }
                }
            }
        },
        "/content/list": {
            "get": {
                "description": "Content 목록 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Content 목록 조회",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ContentList"
                        }
                    }
                }
            }
        },
        "/content/{id}": {
            "get": {
                "description": "Content 상세 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Content 상세 조회",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Content id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Content"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Content 삭제",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Content id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Content"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Content 수정",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Content id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Content Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ContentData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Content"
                        }
                    }
                }
            }
        },
        "/rbac/role": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC"
                ],
                "summary": "Role 생성",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dblayer.RoleData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    }
                }
            }
        },
        "/rbac/role/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC"
                ],
                "summary": "Role 목록 조회",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dblayer.RolePage"
                        }
                    }
                }
            }
        },
        "/rbac/role/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC"
                ],
                "summary": "Role 삭제",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC"
                ],
                "summary": "Role Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update에 사용할 Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dblayer.RoleData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "삭제된 Role 데이터",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dblayer.RoleData": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dblayer.RolePage": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "next": {
                    "type": "string"
                },
                "previous": {
                    "type": "string"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Role"
                    }
                }
            }
        },
        "models.AddUserData": {
            "type": "object",
            "required": [
                "login_id",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "login_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Content": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "content_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "type": "integer"
                }
            }
        },
        "models.ContentData": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ContentItem": {
            "type": "object",
            "properties": {
                "content_id": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ContentList": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "next": {
                    "type": "string"
                },
                "previous": {
                    "type": "string"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ContentItem"
                    }
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "login_id",
                "password"
            ],
            "properties": {
                "login_id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.LoginResult": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Role": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "rest.accessToken": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "rest.refreshToken": {
            "type": "object",
            "properties": {
                "refresh_token": {
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
	Title:            "RBAC GO API",
	Description:      "This is a RBAC server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
