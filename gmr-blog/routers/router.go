package routers

import (
	"gmr/gmr-blog/pkg/setting"

	"github.com/gin-gonic/gin"
	"gmr/gmr-blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	appv1 := router.Group("/api/v1")
	{
		appv1.GET("/tags",v1.GetTags)
		appv1.POST("/tags",v1.AddTag)
		appv1.PUT("/tags/:id",v1.EditTag)
		appv1.DELETE("/tags/:id",v1.DeleteTag)
	}
	return router
}