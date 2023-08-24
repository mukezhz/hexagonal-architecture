package application

import (
	"fmt"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
)

type ExcelUseCase struct {
	RouteStores       []domain.RouteStore
	ExcelOutgoingPort domain.ExcelOutgoingPort
	ExcelRepository   domain.ExcelRepository
}

func NewExcelUseCase(excelOutgoingPort domain.ExcelOutgoingPort, excelRepository domain.ExcelRepository) *ExcelUseCase {
	return &ExcelUseCase{ExcelOutgoingPort: excelOutgoingPort, ExcelRepository: excelRepository}
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

func (e *ExcelUseCase) SaveToDB() error {
	if err := e.ExcelRepository.CreateExcel(e.RouteStores); err != nil {
		return err
	}
	return nil
}
