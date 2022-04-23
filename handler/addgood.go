package handler

import (
	"log"
	"miniproject/model"
	"miniproject/model/mysql"
	"miniproject/model/tables"
	easy "miniproject/pkg/easygo"
	"miniproject/pkg/response"
	"miniproject/pkg/upload"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tmp struct {
	Title   string `json:"title" binding:"required"`
	Zone    string `json:"zone" binding:"required"`
	Summary string `json:"summary"`
	Price   string `json:"price" binding:"required"`
	Avatar  string `json:"avatar" binding:"required"`
	Way     string `json:"way" binding:"required"`
}

//@Summary "上架商品"
//@Description "新增一个商品时的api"
//@Tags Good
//@Produce application/json
//@Produce application/json
//@Param infor body Tmp true "detail"
//@Success 200 {object} response.Resp "upload successfully"
//@Failure 500 {object} response.Resp "error happened"
//@Router /money/goods/addition [post]
func Addgood(c *gin.Context) {
	//新增一个商品,goodsid不需要去获取，设置了自增就可以，只要管其他的字段
	var (
		good1 tables.Good
		good2 tables.Good
		msga  string
		msgw  string
		input Tmp
	)

	stuid, exists := c.MustGet("id").(string)
	if !exists {
		response.SendResponse(c, "error happened in server", 500)
		log.Println("get错误")
		return
	}

	//err := mysql.DB.Model(&tables.Good{}).Last(&good1).Error
	good1, err := model.GetLastRecord()

	if err != nil {
		log.Println("database")
		response.SendResponse(c, "database error", 500)
		return
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		log.Println("bind err", err)
		response.SendResponse(c, "bind error", 500)
		return
	}

	price := easy.STI(input.Price)
	if price == -1 {
		response.SendResponse(c, "STI", 500)
		return
	}

	good2.ID = stuid
	//获取到当前的最大值之后再加一即可
	good2.GoodsID = good1.GoodsID + 1
	//fmt.Println(good1.GoodsID, good2.GoodsID)

	//存放图片到本地
	oka := upload.Upload(c, good2.GoodsID, "avatar", input.Avatar)
	if !oka {
		msga = "the avatar failed to upload!"
	}

	//存放图片到本地
	okw := upload.Upload(c, good2.GoodsID, "way", input.Way)
	if !okw {
		msgw = "the way failed to upload!"
	}
	//good2.Way = wayW

	good2.Price = price
	good2.Goodszone = input.Zone
	if input.Summary == "" {
		good2.Summary = "该商品暂无其他描述"
	} else {
		good2.Summary = input.Summary
	}

	good2.Title = input.Title

	//直接存url
	good2.Avatar = "http://124.221.246.5:8080/images/avatar/" + strconv.Itoa(good2.GoodsID) + ".jpg"
	good2.Way = "http://124.221.246.5:8080/images/way/" + strconv.Itoa(good2.GoodsID) + ".jpg"
	good2.Goodsin = "yes"

	mysql.DB.Model(&tables.Good{}).Create(&good2)

	response.SendResponse(c, "upload successfully!"+msga+","+msgw, 200)
}
