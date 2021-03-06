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
            "email": "942801422@qq.com"
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
        "/article/info/{id}": {
            "get": {
                "description": "查询单个文章接口",
                "tags": [
                    "文章接口"
                ],
                "summary": "查询单个文章",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "文章编号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询文章成功",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseUser"
                        }
                    },
                    "400": {
                        "description": "查询文章失败",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            }
        },
        "/user/add": {
            "post": {
                "description": "新增用户接口",
                "tags": [
                    "用户接口"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "新增用户成功",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseUser"
                        }
                    },
                    "400": {
                        "description": "新增用户失败",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.ResponseError": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "Error"
                },
                "status": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "v1.ResponseUser": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "OK"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "v1.UserInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "1233456"
                },
                "role": {
                    "description": "权限码",
                    "type": "integer",
                    "example": 2
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "lin"
                }
            }
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
	Host:        "localhost:2333",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Gin Blog API v1",
	Description: "Gin博客接口文档",
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
