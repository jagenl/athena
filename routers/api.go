package routers

import (
	"athena/app/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InitApiRouter(r *gin.Engine) *gin.Engine {
	api :=r.Group("api")
	{
		middlelibs := &controller.Middlelibs{};
		api.POST("/data",middlelibs.Save)
		api.GET("/content",middlelibs.Read)
	}

	// 处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "路由未找到 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	return r;
}