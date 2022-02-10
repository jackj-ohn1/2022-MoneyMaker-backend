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
        "termsOfService": "http://swagger.io/terrms/",
        "contact": {
            "name": "yyj",
            "email": "2105753640@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/money/entrance": {
            "post": {
                "description": "\"登录的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "\"用户登录\"",
                "parameters": [
                    {
                        "description": "id 学号 password 密码进行base64加密后的字符串",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.user"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":   \"登录成功\",\"token\": token,\"tips\": \"请保留token并将其放在之后的请求头中\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "msg\":\"unauthorization\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"token生成错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/goods/addition": {
            "post": {
                "description": "\"新增一个商品时的api\"",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Good"
                ],
                "summary": "\"上架商品\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "分区",
                        "name": "zone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "价格",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "详情",
                        "name": "summary",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "商品图二进制文件",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "联系方式二进制文件",
                        "name": "way",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"upload successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/goods/comment": {
            "post": {
                "description": "\"用户做出评价，点击评价时的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "\"用户对某个商品的评论\"",
                "parameters": [
                    {
                        "description": "评论",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Description"
                        }
                    },
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"give successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/goods/comments": {
            "get": {
                "description": "\"商品详情页点击评价时的api \"scores\":所有分值情况, \"infor\":\"评论信息以及学号\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "\"获取某个商品的所有评论\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "infor\":[]tables.Comment,\"score\":All}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"err\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/goods/deletion": {
            "delete": {
                "description": "\"下架商品的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Good"
                ],
                "summary": "\"商家下架商品\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"delete successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/goods/feedback": {
            "post": {
                "description": "\"举报的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Good"
                ],
                "summary": "\"接收举报\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "原因 只需上传用户勾选的个数 内容不需要",
                        "name": "reasonNum",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"举报成功!\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/goods/scanning": {
            "get": {
                "description": "\"点击进入商品详情页的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Good"
                ],
                "summary": "\"查询某个商品的详细信息\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"success\",\"infor\":tables.User,\"user\":tables.User,\"tips\":\"如果goodsin是no则代表已经下架，此时则不显示开启交易按钮\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/goods/shopping": {
            "get": {
                "description": "\"点击购买时的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Good"
                ],
                "summary": "\"用户进行购买\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"success\",\"way\":\"联系方式对应的url\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/homepage": {
            "get": {
                "description": "\"order=1-\u003e返回前十个商品的内容，summary不需要展示出来，是在商品详情页里，主页的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Good"
                ],
                "summary": "\"主页内容\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"success\",\"infor\":[]tables.Good}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/message": {
            "get": {
                "description": "\"信息已根据时间按升序排列，id越大越新，消息通知的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "\"返回用户购买、收藏后的信息\"",
                "responses": {
                    "200": {
                        "description": "msg\":\"success\",\"infor\":[]tables.Message}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error happened"
                    }
                }
            }
        },
        "/money/my/cancellation": {
            "post": {
                "description": "\"在购物车里取消收藏的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Star"
                ],
                "summary": "\"取消收藏\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"cancel successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/my/cart": {
            "get": {
                "description": "\"购物车\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "My"
                ],
                "summary": "\"查看我的购物车\"",
                "responses": {
                    "200": {
                        "description": "msg\":\"check successfully\",\"infot\":[]tables.Good}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "204": {
                        "description": "msg\":\"check successfully\",\"infot\":\"nothing\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/my/goods": {
            "get": {
                "description": "\"橱窗\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "My"
                ],
                "summary": "\"查看我的橱窗\"",
                "responses": {
                    "200": {
                        "description": "msg\":\"check successfully\",\"infot\":[]tables.Good}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/my/goods/finish": {
            "get": {
                "description": "\"点击确认完成时的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trade"
                ],
                "summary": "\"商家完成订单\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error happened"
                    }
                }
            }
        },
        "/money/my/goods/unfish": {
            "get": {
                "description": "\"返回订单,需要点完成的是‘my sell’-\u003e[]string 是与我做交易的人的id,因为一个商品可能被多个人购买，所以string切片的长度就是‘完成订单’的订单数,点评价的是‘my buy’\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trade"
                ],
                "summary": "\"返回用户与卖家未完成的订单\"",
                "responses": {
                    "200": {
                        "description": "msg\":\"success\",\"my buy\":[]tables.Good,\"my sell\":[]Two}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/my/message": {
            "get": {
                "description": "\"我的个人信息的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "My"
                ],
                "summary": "\"返回我的信息\"",
                "responses": {
                    "200": {
                        "description": "msg\":\"avatar 是头像对应的url\",\"infor\":tables.User}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\",\"infor\":tables.User}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/new_star": {
            "patch": {
                "description": "\"添加至购物车时的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Star"
                ],
                "summary": "\"新增收藏\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品的编号",
                        "name": "goodsid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"add successfully\" \"msg\":\"你已经收藏过该商品了\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "msg\":\"error happened\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/money/search": {
            "post": {
                "description": "\"order=1-\u003e返回前十个商品的内容，summary不需要展示出来-》在商品详情页里,搜索的api\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Good"
                ],
                "summary": "\"搜索并返回已排序的商品信息\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索框输入的内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg\":\"search successfully\",\"infor\":[]tables.Good}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "204": {
                        "description": "msg\":\"find nothing\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Description": {
            "type": "object",
            "required": [
                "comment",
                "score"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "handler.user": {
            "type": "object",
            "required": [
                "id",
                "password"
            ],
            "properties": {
                "id": {
                    "description": "一定要输入的加上了required",
                    "type": "string"
                },
                "password": {
                    "type": "string"
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
	Version:     "1.0.0",
	Host:        "localhost:8080",
	BasePath:    "/api/vi",
	Schemes:     []string{"http"},
	Title:       "miniproject",
	Description: "\"赚圈圈API 返回的goods如果其中的goodsin为yes代表它可以进行交易，即：未下架\"",
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
