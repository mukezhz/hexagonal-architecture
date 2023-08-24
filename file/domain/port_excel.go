package domain

type ExcelIncomingPort interface {
	Extract(filePath string) error
	Print()
	SaveToDB() error
}

type ExcelOutgoingPort interface {
	GetAll(filePath string) ([]RouteStore, error)
}

type ExcelRepository interface {
	CreateExcel(data []RouteStore) error
}
