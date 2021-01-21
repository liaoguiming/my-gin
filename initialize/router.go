package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "liao/docs"
	"liao/global"
	"liao/middleware"
	"liao/router"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	//具体的路由信息
	ApiGroup := Router.Group("")
	global.LOG.Debug("router register start")
	router.InitApiRouter(ApiGroup)
	//添加其他路由请在此处注册
	//TODO
	//生成文档信息
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Debug("swagger register success")
	return Router
}
