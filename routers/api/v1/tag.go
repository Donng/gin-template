package tag

import (
	"gin-blog/pkg/e"
	"gin-template/models"
	"gin-template/pkg/err"
	"gin-template/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询标签，分页查询，条件为某人的标签或者状态为 1，0 的标签
func Get(c *gin.Context) {
	maps := make(map[string]interface{})

	if name := c.Query("name"); name != "" {
		maps["name"] = name
	}
	if state := c.Query("state"); state != "" {
		maps["state"] = state
	}
	page := com.StrTo(c.Query("page")).MustInt()
	size := setting.App.PageSize
	offset := (page - 1) * size

	code := e.SUCCESS
	data := models.GetTags(maps, offset, size)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err.GetMsg(code),
		"data": data,
	})
}

func Create(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}
