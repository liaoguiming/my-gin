package router

import (
	"github.com/gin-gonic/gin"
	"liao/api/v1"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.GET("/", v1.Province)
		ApiRouter.GET("/getCateList/:selectType", v1.GetCateList)
		ApiRouter.GET("/getMainCategory", v1.GetMainCategory)
		ApiRouter.POST("/getArticleList", v1.GetArticleList)
		ApiRouter.POST("/getArticleDetail", v1.GetArticleDetail)
		ApiRouter.POST("/getIndexArticleList", v1.GetIndexArticleList)
		ApiRouter.POST("/addPublish", v1.AddPublish)
		ApiRouter.POST("/publishList", v1.PublishList)
		ApiRouter.POST("/upload", v1.Upload)
	}
}
