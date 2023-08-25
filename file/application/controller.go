package application

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"net/http"
	"path/filepath"
)

type FileController struct {
	fileUseCase  domain.FileIncomingPort
	excelUseCase domain.ExcelIncomingPort
}

func NewFileController(fileUseCase domain.FileIncomingPort, excelUseCase domain.ExcelIncomingPort) FileController {
	return FileController{fileUseCase: fileUseCase, excelUseCase: excelUseCase}
}

func (c *FileController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/", c.uploadFile)
}

func (c *FileController) uploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filePath := filepath.Join("uploads", file.Filename)
	if err := c.fileUseCase.Upload(file, filePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.fileUseCase.Save(domain.FileMetadata{
		Filename: file.Filename,
		UUID:     uuid.New().String(),
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// extract the xlxs file
	if err = c.excelUseCase.Extract(filePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.excelUseCase.SaveToDB(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "File uploaded successfully"})
}
