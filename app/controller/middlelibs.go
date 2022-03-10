package controller

import (
	"athena/app/model"
	http2 "athena/app/service/http"
	"athena/pkg/database"
	"athena/pkg/helpers"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)
//type LoginByPhoneRequest struct {
//	Phone      string `json:"phone,omitempty" valid:"phone"`
//	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
//}
var(
)
type Middlelibs struct{
}

type MiddlelibsRequest struct {
	From string `json:"from" form:"from" binding:"required"`
	For  string `json:"for" form:"for" binding:"required"`
	IsQRcode bool `json:"need_qrcode" form:"need_qrcode" binding:"required"`
	Data interface{} `json:"data" form:"data" binding:"required"`
}
func (u *Middlelibs) Save(c *gin.Context) {
	var obj MiddlelibsRequest

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":10010,
			"error": err.Error(),
		})
		return
	}
	sn := helpers.RandomString(10)
	var qrcodeUrl string
	if obj.IsQRcode == true {
		//todo blhttpservice
		//if body, err := ioutil.ReadAll(result.Body); err !=nil{
		//	dump.P(body)
		//	//qrcodeUrl = body.qrcode
		//}
		//var res result
		//_ = json.Unmarshal(body,&res)
		//fmt.Printf("%#v", res)
	}

	//写入数据
	content,err := json.Marshal(obj)
	if err !=nil{
		c.JSON(http.StatusOK, gin.H{
			"code":10010,
			"error": "出错了",
		})
	}
	MiddleModel := model.MiddleLib{
		Sn:      sn,
		Source:  obj.For,
		Content: content,
		CodeUrl: qrcodeUrl,
	}
	database.DB.Create(&MiddleModel) // 通过数据的指针来创建


}

type ReadRequest struct {
	Sn string `json:"sn" form:"sn" uri:"sn" binding:"required"`
}

func (u *Middlelibs) Read(c *gin.Context) {
	var obj ReadRequest
	if err := c.BindQuery(&obj); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":10010,
			"error": err.Error(),
		})
		return
	}
	var middle model.MiddleLib
	database.DB.Where("sn = ?", obj.Sn).Find(&middle)
	if middle.Id == 0  {
		c.JSON(http.StatusOK,gin.H{
			"code":10010,
			"error": "数据不存在",
		})
		return
	}
	//middle 中的 Content 我定义的是一个string 其实他是一个json
	//取出这条数据后 我想针对content 进行json string转json对象操作
	//然后再返回给第三方 middle的数据

	c.JSON(http.StatusOK, gin.H{
		"code":0,
		"message":"操作成功",
		"data": middle,
	})

}

func (u *Middlelibs) Update(c *gin.Context) {

}