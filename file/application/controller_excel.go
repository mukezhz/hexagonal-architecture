package application

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"log"
	"net/http"
	"path/filepath"
)

type ExcelController struct {
	FileUseCase  domain.FileIncomingPort
	excelUseCase domain.ExcelIncomingPort
}

func NewExcelController(excelUseCase domain.ExcelIncomingPort) ExcelController {
	return ExcelController{excelUseCase: excelUseCase}
}

func (c *ExcelController) RegisterRoutesExcel(router *gin.RouterGroup) {
	router.GET("/:store_id", c.fetchExcel)
	router.POST("/", c.addExcel)
}

func (c *ExcelController) fetchExcel(ctx *gin.Context) {
	storeId := ctx.Param("store_id")
	log.Println(storeId)
	all, err := c.excelUseCase.Fetch(storeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Fetched From DynamoDb", "items": all})
}

func (c *ExcelController) addExcel(ctx *gin.Context) {
	var filePayload map[string]string
	if err := ctx.ShouldBindJSON(&filePayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filePath := filepath.Join("uploads", filePayload["file_name"])
	newData, err := c.excelUseCase.GetDifference(BUCKET_NAME, filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "New Data are stored", "items": newData})
}
