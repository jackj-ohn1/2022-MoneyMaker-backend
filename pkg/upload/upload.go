package upload

import (
	"encoding/base64"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, id int, which string, infor string) bool {
	var way string

	file_str, err := base64.StdEncoding.DecodeString(infor)
	if err != nil {
		log.Println("解码失败")
		return false
	}

	//filename := header.Filename
	//fmt.Println(file, err, filename)
	//以goodsid作为名字
	way = "./goods/" + which + "/" + strconv.Itoa(id) + ".jpg"

	err = ioutil.WriteFile(way, file_str, 0777)
	if err != nil {
		log.Println(err, "存储失败")
		return false
	}

	//fmt.Println("存储成功")
	err = CompressImg(way)
	if err != nil {
		log.Println(err, "压缩失败")
		return false
	}
	return true
}

func CompressImg(source string) error {
	var err error
	var file *os.File
	if file, err = os.Open(source); err != nil {
		log.Println(err, "压缩1")
		return err
	}

	defer file.Close()
	var img image.Image
	if img, err = jpeg.Decode(file); err != nil {
		log.Println(err, "压缩2")
		return err
	}

	if outFile, err := os.Create(source); err != nil {
		log.Println(err, "压缩3")
		return err
	} else {
		defer outFile.Close()
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 20})
		if err != nil {
			log.Println(err, "压缩4")
			return err
		}
	}
	return nil
}
