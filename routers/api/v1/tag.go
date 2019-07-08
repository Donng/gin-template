package v1

import (
	"gin-template/models"
	"gin-template/pkg/e"
	"gin-template/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 查询标签，分页查询，条件为某人的标签或者状态为 1，0 的标签
func GetTags(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

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
	data["list"] = models.GetTags(maps, offset, size)
	data["total"] = models.GetTagsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func CreateTag(c *gin.Context) {
	name := c.PostForm("name")
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("标签名称不能为空")
	valid.MaxSize(name, 20, "name").Message("标签名称长度不能超过20个字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 15, "created_by").Message("创建人长度不能超过15个字符")

	code := e.SUCCESS
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, err := range valid.Errors {
			log.Printf("Requst err key: %s, message: %s", err.Key, err.Message)
		}
	} else {
		// 判断标签是否已存在
		if models.ExistTagByName(name) {
			code = e.ERROR_EXIST_TAG
		} else {
			models.CreateTag(name, createdBy)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

func UpdateTag(c *gin.Context) {
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	valid.Required(id, "id").Message("标签ID不能为空")
	valid.Min(id, 1, "id").Message("标签ID必须大于0")

	name := c.Query("name")
	if name != "" {
		valid.MaxSize(name, 20, "name").Message("标签名称长度不能超过20个字符")
	}

	modifiedBy := c.Query("modified_by")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 15, "modified_by").Message("修改人长度不能超过15个字符")

	code := e.SUCCESS
	if valid.HasErrors() {
		code = e.ERROR
		for _, err := range valid.Errors {
			log.Printf("Error params key: %s, message: %s", err.Key, err.Message)
		}
	} else {
		// 检查 id 是否存在
		// 检查是否有重名的 name
		// 更新
		if !models.ExistTagById(id) {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			if name != "" {
				maps["name"] = name
			}
			if modifiedBy != "" {
				maps["modified_by"] = modifiedBy
			}
			models.UpdateTag(id, maps)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id").Message("用户ID不能为空")

	code := e.SUCCESS
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, err := range valid.Errors {
			log.Printf("Error Params key: %s, message %s", err.Key, err.Message)
		}
	} else {
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": "",
	})
}
