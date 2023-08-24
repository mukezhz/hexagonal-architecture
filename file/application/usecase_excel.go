package application

import (
	"fmt"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
)

type ExcelUseCase struct {
	RouteStores       []domain.RouteStore
	ExcelOutgoingPort domain.ExcelOutgoingPort
}

func NewExcelUseCase(excelOutgoingPort domain.ExcelOutgoingPort) *ExcelUseCase {
	return &ExcelUseCase{ExcelOutgoingPort: excelOutgoingPort, RouteStores: make([]domain.RouteStore, 0)}
}

func (e *ExcelUseCase) Extract(filePath string) error {
	data, err := e.ExcelOutgoingPort.GetAll(filePath)
	if err != nil {
		return err
	}
	e.RouteStores = data
	return nil
}

func (e *ExcelUseCase) Print() {
	for _, d := range e.RouteStores {
		fmt.Println(d)
	}
}
