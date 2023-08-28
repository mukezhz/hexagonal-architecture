package application

import (
	"fmt"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"strings"
)

type ExcelUseCase struct {
	RouteStores       []domain.RouteStore
	ExcelOutgoingPort domain.ExcelOutgoingPort
	ExcelRepository   domain.ExcelRepository
}

func NewExcelUseCase(excelOutgoingPort domain.ExcelOutgoingPort, excelRepository domain.ExcelRepository) *ExcelUseCase {
	return &ExcelUseCase{ExcelOutgoingPort: excelOutgoingPort, ExcelRepository: excelRepository}
}

func (e *ExcelUseCase) Extract(dst, filePath string) error {
	data, err := e.ExcelOutgoingPort.GetAll(dst, filePath)
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

func (e *ExcelUseCase) Fetch(storeName string) ([]domain.RouteStore, error) {
	stores, err := e.ExcelRepository.GetAllExcel(storeName)
	if err != nil {
		return make([]domain.RouteStore, 0), err
	}
	return stores, err
}

func (e *ExcelUseCase) GetDifference(dst, filePath string) ([]domain.RouteStore, error) {
	var allFromExcel []domain.RouteStore
	if err := e.Extract(dst, filePath); err != nil {
		return make([]domain.RouteStore, 0), err
	}
	allFromExcel = e.RouteStores
	uniqueRouteNames := make(map[string]int)
	uniqueKeysFromExcel := make(map[string]int)
	for _, e := range allFromExcel {
		if uniqueRouteNames[e.RouteName] == 0 {
			uniqueRouteNames[e.RouteName] += 1
		}
		uniqueKeysFromExcel[fmt.Sprintf("%v::%v", e.RouteName, e.Store)] = 1
	}
	var allFromDB []domain.RouteStore
	for k, _ := range uniqueRouteNames {
		routeStores, err := e.Fetch(k)
		if err != nil {
			return make([]domain.RouteStore, 0), err
		}
		allFromDB = append(allFromDB, routeStores...)
	}
	uniqueKeysFromDB := make(map[string]int)
	for _, e := range allFromDB {
		uniqueKeysFromDB[fmt.Sprintf("%v::%v", e.RouteName, e.Store)] = 1
	}

	data := make([]domain.RouteStore, 0)
	for k, _ := range uniqueKeysFromExcel {
		if uniqueKeysFromDB[k] == 1 {
			// item exits in both excel and DB no need to add in slice
			continue
		}
		split := strings.Split(k, "::")
		data = append(data, domain.RouteStore{
			RouteName: split[0],
			Store:     split[1],
		})
	}
	return data, nil
}
