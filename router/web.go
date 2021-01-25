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
		ApiRouter.GET("/mainTypeList", v1.MainTypeList)
	}
}
