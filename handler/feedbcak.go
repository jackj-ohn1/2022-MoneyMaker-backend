package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

//举报
//@Summary "接收举报"
//@Description "举报的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Param reasonNum formData string true "只需上传用户勾选的个数 内容不需要"
//@Success 200 {object} response.Resp "举报成功!"
//@Failure 500 {object} response.Resp "error happened in server"
//@Failure 304 {object} response.Resp "error in database"
//@Router /money/goods/feedback [post]
func Feedback(c *gin.Context) {
	var good tables.Good
	goodsstr := c.Query("goodsid")
	goodsid := easy.STI(goodsstr)

	resonNum := c.PostForm("reasonNum")
	num := easy.STI(resonNum)

	//mysql.DB.Select("feed_back").Where("goods_id=?", goodsid).Find(&good)
	good, err := model.GetOrderGood(goodsid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}

	if goodsid == -1 && num == -1 {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(goodsid, num)
		return
	}

	mysql.DB.Model(&tables.Good{}).Where("goods_id=?", goodsid).Update("feed_back", good.FeedBack+num)

	response.SendResponse(c, "举报成功", 200)
}
