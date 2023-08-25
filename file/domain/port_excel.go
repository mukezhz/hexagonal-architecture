package domain

type ExcelIncomingPort interface {
	Extract(filePath string) error
	Print()
	SaveToDB() error
	Fetch(storeName string) ([]RouteStore, error)
	GetDifference(filePath string) ([]RouteStore, error)
}

type ExcelOutgoingPort interface {
	GetAll(filePath string) ([]RouteStore, error)
}

type ExcelRepository interface {
	CreateExcel(data []RouteStore) error
	GetAllExcel(store string) ([]RouteStore, error)
}
