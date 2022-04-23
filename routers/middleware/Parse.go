package middleware

import (
	"miniproject/config"
	"miniproject/pkg/response"
	"miniproject/pkg/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Parse(c *gin.Context) {
	//dvar stu controller.Json
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, response.Resp{
			Msg:  "请求头中的auth为空",
			Data: nil,
			Code: 401,
		})
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		c.JSON(http.StatusUnauthorized, response.Resp{
			Msg:  "请求头中的auth格式有误",
			Data: nil,
			Code: 401,
		})
		c.Abort()
		return
	}

	token, err := token.ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Resp{
			Msg:  "token无效",
			Data: nil,
			Code: 401,
		})
		c.Abort()
		return
	}

	issuer := token.Issuer
	//_, err := model.GetUserInfoFormOne()

	if issuer != config.Issuer {
		c.JSON(http.StatusUnauthorized, response.Resp{
			Msg:  "发布者错误",
			Data: nil,
			Code: 401,
		})
		c.Abort()
		return
	}

	id := token.ID
	c.Set("id", id)
}
