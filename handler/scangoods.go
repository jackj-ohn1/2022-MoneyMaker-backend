package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"

	"github.com/gin-gonic/gin"
)

type Userinfor struct {
	User tables.User   `json:"user"`
	Good tables.Good   `json:"good"`
	All  []tables.Good `json:"all"`
}

//@Summary "查询某个商品的详细信息"
//@Description "点击进入商品详情页的api"
//@Tags Good
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {object} response.Resp "success"
//@Failure 500 {object} response.Resp "error happened"
//@Failure 304 {object} response.Resp "error in database"
//@Router /money/goods/scanning [get]
func Scan(c *gin.Context) {
	var good tables.Good
	var user_goods []tables.Good
	var user tables.User
	var msg = ""

	/*id, exists := c.MustGet("id").(string)
	if !exists {
		response.SendResponse(c, "error happened", 500)
		log.Println(exists)
		return
	}*/

	goodsidstr := c.Query("goodsid")
	goodsid := easy.STI(goodsidstr)
	//err := mysql.DB.Where("goods_id=?", goodsid).Find(&good).Error

	good, err := model.GetOrderGood(goodsid)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}

	user_goods, _ = model.GetUserAllGoods(good.ID)
	for i := 0; i < len(user_goods); i++ {
		user_goods[i].Way = ""
	}

	if goodsid == -1 {
		response.SendResponse(c, "error happened!", 500)
		return
	}

	good.Way = ""
	good.Buyer = ""
	//mysql.DB.Model(&tables.User{}).Where("id=?", good.ID).Find(&user)

	user, err = model.GetOrderUser(good.ID)
	if err != nil {
		log.Println(err)
		response.SendResponse(c, "error in database", 304)
	}
	if good.FeedBack >= 100 {
		msg = "该商品被举报次数过多,请谨慎交易!"
	}

	user.Buygoods = ""

	c.JSON(200, response.Resp{
		Code: 200,
		Msg:  "success " + msg,
		Data: Userinfor{
			User: user,
			Good: good,
			All:  user_goods,
		},
	})
}
