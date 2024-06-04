package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/Controller"
	"main.go/docs"
)

// @title Documenting API (Hello World)
// @version 1.0.0
// @Description A three tier architecture that outputs hello world
// @host localhost:8080
// @BasePath /
func main() {

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/HelloWorld", Controller.HelloWorld)
	router.Run(":8080")

}
