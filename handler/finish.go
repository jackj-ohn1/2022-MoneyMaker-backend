package handler

import (
	"fmt"
	"log"
	"miniproject/model"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type Two struct {
	Good  tables.Good
	Buyer []string
}

type ReturnType struct {
	Strbuy  []tables.Good `json:"buy"`
	Strsell []Two         `json:"sell"`
}

//@Summary "返回用户与卖家未完成的订单"
//@Description "返回订单,需要点完成的是‘my sell’->[]string 是与我做交易的人的id,因为一个商品可能被多个人购买，所以string切片的长度就是‘完成订单’的订单数,点评价的是‘my buy’"
//@Tags Trade
//@Accept application/json
//@Produce application/json
//@Success 200 {object} response.Resp "success"
//@Failure 500 {object} response.Resp "error happened"
//@Router /money/my/goods/unfinish [get]
func UnFinish(c *gin.Context) {
	var (
		goods   []tables.Good
		user    tables.User
		strbuy  []tables.Good
		strsell []Two
		two     Two
		max     int
		hasbuy  []string
	)

	id, exists := c.MustGet("id").(string)
	if !exists {
		//fmt.Println("1", exists, ok)
		response.SendResponse(c, "unauthorized", 401)
		return
	}
	//确认完成
	//获取已购买我发布的商品的用户的id,可能不止一个商品,mysell
	mysql.DB.Model(&tables.Good{}).Select("price", "title", "goods_id", "buyer", "id", "goodsin").Where("id=?", id).Find(&goods)
	a := len(goods)
	for i := a - 1; i >= 0; i-- {
		if goods[i].Goodsin == "no" {
			continue
		}
		buyer := strings.Split(goods[i].Buyer, ",")
		two.Buyer = buyer
		two.Good = goods[i]
		strsell = append(strsell, two)
		//fmt.Println(two)
	}

	//我买入的
	mysql.DB.Where("id=?", id).Find(&user)
	if user.Buygoods == "" {
		max = 0
	} else {
		hasbuy = strings.Split(user.Buygoods, ",")
		fmt.Println(hasbuy)
		max = len(hasbuy)
	}

	if max == 0 {
		c.JSON(200, response.Resp{
			Data: strsell,
			Msg:  "success",
			Code: 200,
		})
		return
	}
	for i := max - 1; i >= 0; i-- {
		num := easy.STI(hasbuy[i])
		if num == -1 {
			log.Println("for")
			response.SendResponse(c, "error", 500)
			return
		} else {
			var good tables.Good
			mysql.DB.Model(&tables.Good{}).Select("price", "title", "goods_id", "id", "goodsin").Where("goods_id=?", num).Find(&good)
			if good.Goodsin == "no" {
				continue
			}
			strbuy = append(strbuy, good)
		}
	}
	c.JSON(200, response.Resp{
		Data: ReturnType{
			Strbuy:  strbuy,
			Strsell: strsell,
		},
		Msg:  "success",
		Code: 200,
	})
}

//@Summary "商家完成订单"
//@Description "点击确认完成时的api"
//@Tags Trade
//@Accept application/json
//@Produce application/json
//@Param goodsid query string true "商品编号"
//@Success 200 {object} response.Resp "success"
//@Failure 500 {object} response.Resp "error happened in server"
//@Router /money/my/goods/finish [get]
func Finsh(c *gin.Context) {
	var (
		good tables.Good
		user tables.User
		re   string
	)

	//点击完成之后把这个购买者从buyer中删去，以及goodsid从uesr中删除,有多个的情况下，则只删除一个
	goodsid := c.Query("goodsid")
	id, exists := c.MustGet("id").(string)

	if !exists {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(exists)
		return
	}

	mysql.DB.Where("goods_id=?", goodsid).Find(&good)
	mysql.DB.Where("id=?", id).Find(&user)

	re = easy.Delete(good.Buyer, id)
	num := easy.STI(goodsid)

	if num == -1 {
		response.SendResponse(c, "error happened in server", 500)
		log.Println(num)
		return
	}
	//mysql.DB.Model(&tables.Good{}).Where("goods_id=?", goodsid).Update("buyer", re)
	model.UpdateGoodBuyer(num, re)

	re = ""
	re = easy.Delete(user.Buygoods, goodsid)

	//mysql.DB.Model(&tables.User{}).Where("id=?", id).Update("buygoods", re)
	model.UpdateBuygoods(id, re)

	c.JSON(200, response.Resp{
		Code: 200,
		Msg:  "successfully",
		Data: nil,
	})
}
