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
        "/api/docs": {
            "post": {
                "description": "创建文档，创建成功之后，data=文档id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文档管理"
                ],
                "summary": "创建文档",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_api.DocCreateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/docs/{id}": {
            "get": {
                "description": "文档内容",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文档管理"
                ],
                "summary": "文档内容",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/image": {
            "post": {
                "description": "上传图片",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "上传图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件上传",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/images": {
            "delete": {
                "description": "删除图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "删除图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "用户登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/logout": {
            "get": {
                "description": "用户注销",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户注销",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/roles": {
            "get": {
                "description": "角色列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "角色列表",
                "parameters": [
                    {
                        "type": "string",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-role_api_RoleListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "更新角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "更新角色",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/role_api.RoleCreateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "创建角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/role_api.RoleCreateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/site": {
            "get": {
                "description": "站点配置查询",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "站点配置"
                ],
                "summary": "站点配置查询",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/config.Site"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "description": "用户列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "roleID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_UserModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "管理员更新用户信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "管理员更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "创建用户，只能管理员创建",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserCreateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除用户",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "idList",
                        "name": "IDList",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/users_info": {
            "get": {
                "description": "用户信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/user_api.UserInfoResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "用户更新自己的信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户更新自己的信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserUpdateInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/users_password": {
            "put": {
                "description": "用户修改密码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户修改密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserUpdatePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.Site": {
            "type": "object",
            "properties": {
                "abstract": {
                    "description": "简介",
                    "type": "string"
                },
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "footer": {
                    "description": "尾部信息",
                    "type": "string"
                },
                "href": {
                    "description": "点击go的跳转链接",
                    "type": "string"
                },
                "icon": {
                    "description": "首页的图标",
                    "type": "string"
                },
                "iconHref": {
                    "description": "图标链接",
                    "type": "string"
                },
                "title": {
                    "description": "网站名称",
                    "type": "string"
                }
            }
        },
        "doc_api.DocCreateRequest": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "parentID": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.RoleModel": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "添加时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "isSystem": {
                    "description": "是否是系统角色",
                    "type": "boolean"
                },
                "pwd": {
                    "description": "角色密码",
                    "type": "string"
                },
                "title": {
                    "description": "角色的名称",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "models.UserModel": {
            "type": "object",
            "properties": {
                "addr": {
                    "description": "地址",
                    "type": "string"
                },
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "createdAt": {
                    "description": "添加时间",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "ip": {
                    "description": "ip",
                    "type": "string"
                },
                "lastLogin": {
                    "description": "用户最后登录时间",
                    "type": "string"
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "roleID": {
                    "description": "用户对应的角色",
                    "type": "integer"
                },
                "roleModel": {
                    "description": "用户角色信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.RoleModel"
                        }
                    ]
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "res.Code": {
            "type": "integer",
            "enum": [
                0,
                7,
                9
            ],
            "x-enum-comments": {
                "ErrCode": "系统错误",
                "ValidCode": "校验错误"
            },
            "x-enum-varnames": [
                "SUCCESS",
                "ErrCode",
                "ValidCode"
            ]
        },
        "res.ListResponse-models_UserModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserModel"
                    }
                }
            }
        },
        "res.ListResponse-role_api_RoleListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/role_api.RoleListResponse"
                    }
                }
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/res.Code"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "role_api.RoleCreateRequest": {
            "type": "object",
            "required": [
                "Title"
            ],
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "Pwd": {
                    "type": "string"
                },
                "Title": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 2
                }
            }
        },
        "role_api.RoleListResponse": {
            "type": "object",
            "properties": {
                "DocCount": {
                    "description": "角色拥有的文档数",
                    "type": "integer"
                },
                "UserCount": {
                    "description": "角色下的用户数",
                    "type": "integer"
                },
                "createdAt": {
                    "description": "添加时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "isSystem": {
                    "description": "是否是系统角色",
                    "type": "boolean"
                },
                "pwd": {
                    "description": "角色密码",
                    "type": "string"
                },
                "title": {
                    "description": "角色的名称",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "user_api.UserCreateRequest": {
            "type": "object",
            "required": [
                "password",
                "roleID",
                "userName"
            ],
            "properties": {
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "roleID": {
                    "description": "角色id",
                    "type": "integer"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "user_api.UserInfoResponse": {
            "type": "object",
            "properties": {
                "addr": {
                    "description": "地址",
                    "type": "string"
                },
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "createdAt": {
                    "description": "添加时间",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "ip": {
                    "description": "ip",
                    "type": "string"
                },
                "lastLogin": {
                    "description": "用户最后登录时间",
                    "type": "string"
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "roleID": {
                    "description": "用户对应的角色",
                    "type": "integer"
                },
                "roleModel": {
                    "description": "用户角色信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.RoleModel"
                        }
                    ]
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "user_api.UserLoginRequest": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "user_api.UserUpdateInfoRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                }
            }
        },
        "user_api.UserUpdatePasswordRequest": {
            "type": "object",
            "required": [
                "oldPwd",
                "password"
            ],
            "properties": {
                "oldPwd": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user_api.UserUpdateRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "roleID": {
                    "description": "角色id",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8082",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "文档项目api文档",
	Description:      "API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
