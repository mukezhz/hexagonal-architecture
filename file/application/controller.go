package application

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"net/http"
	"path/filepath"
)

type FileController struct {
	fileUseCase domain.FilePort
}

func NewProductController(fileUC domain.FilePort) FileController {
	return FileController{
		fileUseCase: fileUC,
	}
}

func (controller *FileController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/file", controller.uploadFile)
}

func (controller *FileController) uploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filePath := filepath.Join("uploads", file.Filename)
	if err := controller.fileUseCase.Upload(file, filePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "File uploaded successfully"})
}
