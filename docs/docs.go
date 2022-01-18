// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/createWallet": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "钱包"
                ],
                "summary": "创建钱包地址",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.CreateWalletReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/server.CreateWalletRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/delWallet": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "钱包"
                ],
                "summary": "删除钱包地址",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.DelWalletReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.Response"
                        }
                    }
                }
            }
        },
        "/api/getTransactionReceipt": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "钱包"
                ],
                "summary": "获取交易结果",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.TransactionReceiptReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/server.TransactionReceiptRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/withdraw": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "钱包"
                ],
                "summary": "提现",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.WithdrawReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/server.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/server.WithdrawRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.CreateWalletReq": {
            "type": "object",
            "required": [
                "coinName",
                "protocol"
            ],
            "properties": {
                "coinName": {
                    "description": "币种名称",
                    "type": "string"
                },
                "protocol": {
                    "description": "协议",
                    "type": "string"
                }
            }
        },
        "server.CreateWalletRes": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "生成的钱包地址",
                    "type": "string"
                }
            }
        },
        "server.DelWalletReq": {
            "type": "object",
            "required": [
                "address",
                "coinName",
                "protocol"
            ],
            "properties": {
                "address": {
                    "description": "地址",
                    "type": "string"
                },
                "coinName": {
                    "description": "币种名称",
                    "type": "string"
                },
                "protocol": {
                    "description": "协议",
                    "type": "string"
                }
            }
        },
        "server.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误code码",
                    "type": "integer"
                },
                "data": {
                    "description": "成功时返回的对象"
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                }
            }
        },
        "server.TransactionReceiptReq": {
            "type": "object",
            "required": [
                "coinName",
                "hash",
                "protocol"
            ],
            "properties": {
                "coinName": {
                    "description": "币种名称",
                    "type": "string"
                },
                "hash": {
                    "description": "交易哈希",
                    "type": "string"
                },
                "protocol": {
                    "description": "协议",
                    "type": "string"
                }
            }
        },
        "server.TransactionReceiptRes": {
            "type": "object",
            "properties": {
                "status": {
                    "description": "交易状态（0：未成功，1：已成功）",
                    "type": "integer"
                }
            }
        },
        "server.WithdrawReq": {
            "type": "object",
            "required": [
                "address",
                "coinName",
                "orderId",
                "protocol",
                "value"
            ],
            "properties": {
                "address": {
                    "description": "提现地址",
                    "type": "string"
                },
                "coinName": {
                    "description": "币种名称",
                    "type": "string"
                },
                "orderId": {
                    "description": "订单号",
                    "type": "string"
                },
                "protocol": {
                    "description": "协议",
                    "type": "string"
                },
                "value": {
                    "description": "金额",
                    "type": "integer"
                }
            }
        },
        "server.WithdrawRes": {
            "type": "object",
            "properties": {
                "hash": {
                    "description": "生成的交易hash",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
