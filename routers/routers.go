package routers

import (
	_ "gin-template/docs"
	"gin-template/middleware/jwt"
	"gin-template/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	// 获得Gin的实例Engine
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", v1.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.CreateTag)
		apiv1.PUT("/tags/:id", v1.UpdateTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.POST("/articles", v1.CreateArticle)
		apiv1.PUT("/articles", v1.UpdateArticle)
		apiv1.DELETE("/articles", v1.DeleteArticle)
	}

	return r
}
