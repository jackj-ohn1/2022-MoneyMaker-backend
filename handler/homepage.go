package handler

import (
	"log"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//@Summary "主页内容"
//@Description "主页的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param page query string true "页码"
//@Success 200 {object} response.Resp "success"
//@Success 500 {object} response.Resp "error happened in server"
//@Router /money/homepage [get]
func Homepage(c *gin.Context) {
	var goods []tables.Good

	err := mysql.DB.Order("feed_back desc").Order("goods_id desc").Where("goodsin=?", "yes").Find(&goods).Error
	// mysql.DB.Limit(10).Offset((num-1)*10).Order("feed_back desc").Order("scores desc").Where("goodsin=? AND goods_id>? ", "yes","1").Find(&goods).Error
	if err != nil {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err)
		return
	}

	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}

	c.JSON(200, response.Resp{
		Code: 200,
		Msg:  "successfully",
		Data: goods,
	})
}
