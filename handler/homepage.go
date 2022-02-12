package handler

import (
	"log"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "主页内容"
//@Description "主页的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param page query string true "页码"
//@Success 200 {string} json{"msg":"success","infor":[]tables.Good}
//@Success 500 {string} json{"msg":"error happened in server"}
//@Router /money/homepage [get]
func Homepage(c *gin.Context) {
	var goods []tables.Good

	page := c.DefaultQuery("page", "1")
	num := easy.STI(page)
	err := mysql.DB.Order("feed_back desc").Order("scores desc").Where("goodsin=?", "yes").Find(&goods).Error

	if err != nil || num == -1 {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err, num)
		return
	}

	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}

	if len(goods) < 10 {
		c.JSON(200, gin.H{
			"msg":   "success",
			"goods": goods,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":   "success",
			"goods": goods[:num*10],
		})
	}
}
