package v1

import (
	"gin-template/models"
	"gin-template/pkg/e"
	"gin-template/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAuth(c *gin.Context) {
	data := make(map[string]interface{})

	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")

	code := e.SUCCESS
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, err := range valid.Errors {
			log.Printf("Invalid params key: %s, message: %s", err.Key, err.Message)
		}
	} else {
		if models.CheckAuth(username, password) {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
				log.Printf("Generate token error: %s", err)
			} else {
				data["token"] = token
			}
		} else {
			code = e.ERROR_AUTH
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}
