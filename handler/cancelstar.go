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

//@Summary "取消收藏"
//@Description "在购物车里取消收藏的api"
//@Tags Star
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {object} response.Resp "cancel successfully"
//@Failure 500 {object} response.Resp "error happened in server"
//@Failure 304 {object} response.Resp "error in database"
//@Router /money/my/cancellation [post]
func Cancelstar(c *gin.Context) {
	var cart tables.Cart
	//var str string
	stuid, exists := c.MustGet("id").(string)
	goodsid := c.Query("goodsid")

	//mysql.DB.Where("id=?", stuid).Find(&cart)
	cart, err := model.GetOrderCart(stuid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}
	re := easy.Delete(cart.Goodsid, goodsid)

	err = mysql.DB.Model(&tables.Cart{}).Where("id=?", stuid).Update("goodsid", re).Error
	if !exists || err != nil {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(err, exists)
	}

	goodsNum := easy.STI(goodsid)
	
	if goodsNum == -1{
		log.Println("STI错误")
		response.SendResponse(c, "error happened in server", 500)
		return
	}

	good,err := model.GetOrderGood(goodsNum)
	if err!=nil{
		log.Println("获取good信息失败")
		response.SendResponse(c, "error happened in database", 500)
		return
	}

	easy.Returnstar(stuid,good.ID,false)

	response.SendResponse(c, "cancel successfully", 200)
}
