package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mukezhz/hexagonal-architecture/config"
	fileApplication "github.com/mukezhz/hexagonal-architecture/file/application"
	fileConfig "github.com/mukezhz/hexagonal-architecture/file/config"
	"go.uber.org/fx"
	"log"
)

var ADDRESS string = ":8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fx.New(
		config.Module,
		fx.Options(fileConfig.FileModule),
		fx.Options(fileConfig.ExcelModule),
		fx.Provide(createRouter),
		fx.Invoke(start),
	)
	app.Run()
}

func createRouter() *gin.Engine {
	return gin.Default()
}

func registerAllRoutes(router *gin.Engine, fileController fileApplication.FileController, excelController fileApplication.ExcelController) {
	apiGroup := router.Group("/api")
	fileController.RegisterRoutes(apiGroup.Group("/files"))
	excelController.RegisterRoutesExcel(apiGroup.Group("/sheets"))
}

func start(router *gin.Engine, fileController fileApplication.FileController, excelController fileApplication.ExcelController) {
	registerAllRoutes(router, fileController, excelController)
	if err := router.Run(ADDRESS); err != nil {
		return
	}
}
