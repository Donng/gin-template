package routers

import (
	"gin-template/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 获得Gin的实例Engine
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/tags", tag.Get)
		v1.POST("/tags", tag.Create)
		v1.PUT("/tags/:id", tag.Update)
		v1.DELETE("/tags/:id", tag.Delete)
	}

	// Context是Gin最重要的部分
	r.GET("/ping", func(c *gin.Context) {
		// String将字符串写入响应体
		c.String(http.StatusOK, "ping %s", "pong")
	})

	return r
}
