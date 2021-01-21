package main

import (
	"liao/global"
	"liao/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.Gorm()
	initialize.Redis()
	db, _ := global.DB.DB()
	defer db.Close()
	// 默认启动的是 8080端口，也可以自己定义启动端口
	//gin框架内容
	router := initialize.Router()
	global.LOG.Debug("server run success on :8088")
	router.Run(":8088")
}
