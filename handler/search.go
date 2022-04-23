package handler

import (
	"fmt"
	"log"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	"miniproject/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

//@Summary "搜索并返回已排序的商品信息"
//@Description "order=1->返回前十个商品的内容，summary不需要展示出来-》在商品详情页里,搜索的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param content query string true "搜索框输入的内容"
//@Param page query string true "页码"
//@Success 200 {object} response.Resp "search successfully"
//@Success 204 {object} response.Resp "find nothing"
//@Router /money/search [post]
func Search(c *gin.Context) {
	//根据举报次数与评分进行优先返回
	var goods []tables.Good

	content := c.Query("content")
	content = strings.Replace(content, " ", "", -1)
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "\t", "", -1)

	if content != "" {
		content_Slice := strings.Split(content, "")
		content = "%"
		for i := 0; i < len(content_Slice); i++ {
			content = content + content_Slice[i] + "%"
		}
	}

	fmt.Println(content)
	err := mysql.DB.Order("feed_back asc").Order("goods_id desc").Where(fmt.Sprintf(`title like "%s" AND goodsin=%s`, content, "'yes'")).Find(&goods).Error

	if err != nil {
		response.SendResponse(c, "error happened", 500)
		log.Println(err)
		return
	}
	for i := 0; i < len(goods); i++ {
		goods[i].Way = ""
	}
	//这里不需要返回图片的url

	/*if len(goods) < 10 {
		c.JSON(200, response.Resp{
			Code: 200,
			Msg:  "success",
			Data: goods,
		})
	} else {
		c.JSON(200, response.Resp{
			Code: 200,
			Msg:  "success",
			Data: goods[:num*10],
		})
	}*/
	c.JSON(200, response.Resp{
		Code: 200,
		Msg:  "success",
		Data: goods,
	})
}
