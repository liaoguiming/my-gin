# 开箱即用的Gin开发流程

## 开发流程

- clone代码
	git clone https://github.com/liaoguiming/my-gin.git

- 初始化go模块
	go list

- 在文件夹model下新建数据库对应的model

- 在文件夹api下新建对应的业务逻辑并写好swagger需要的注释

- 在文件夹router下新建api对应的路由并关联api--func

- 在文件夹initialize下router.go引入router

- swag init生成api文档

- go run main.go

## 开发环境

- GoLang Gin Mysql Redis Swagger