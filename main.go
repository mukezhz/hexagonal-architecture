package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/hexagonal-architecture/config"
	fileApplication "github.com/mukezhz/hexagonal-architecture/file/application"
	fileConfig "github.com/mukezhz/hexagonal-architecture/file/config"
	"go.uber.org/fx"
)

var ADDRESS string = ":8080"

func main() {
	app := fx.New(
		fx.Provide(config.NewMysqlDB),
		fx.Options(fileConfig.FileModule),
		fx.Provide(createRouter),
		fx.Invoke(start),
	)
	app.Run()
}

func createRouter() *gin.Engine {
	return gin.Default()
}

func registerAllRoutes(router *gin.Engine, fileController fileApplication.FileController) {
	apiGroup := router.Group("/api")
	fileController.RegisterRoutes(apiGroup)
}

func start(router *gin.Engine, fileController fileApplication.FileController) {
	registerAllRoutes(router, fileController)
	if err := router.Run(ADDRESS); err != nil {
		return
	}
}
