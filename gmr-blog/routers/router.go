package routers

import (
	"gmr/gmr-blog/pkg/setting"
	"gmr/gmr-blog/routers/api"
	"gmr/gmr-blog/routers/api/v1"
	"gmr/gmr-blog/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	router.GET("/auth",api.GetAuth)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//添加标签列表
		apiv1.POST("/tags", v1.AddTag)
		//编辑标签列表
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除标签列表
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	apiv2 := router.Group("/api/v2")
	apiv2.Use(jwt.JWT())
	{
		//获取文章列表
		apiv2.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv2.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv2.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv2.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv2.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return router
}
