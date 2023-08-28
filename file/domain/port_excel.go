package domain

type ExcelIncomingPort interface {
	Extract(dst, filePath string) error
	Print()
	SaveToDB() error
	Fetch(storeName string) ([]RouteStore, error)
	GetDifference(dst, filePath string) ([]RouteStore, error)
}

type ExcelOutgoingPort interface {
	GetAll(dst, filePath string) ([]RouteStore, error)
}

type ExcelRepository interface {
	CreateExcel(data []RouteStore) error
	GetAllExcel(store string) ([]RouteStore, error)
}
