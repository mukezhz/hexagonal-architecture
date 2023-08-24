package domain

type ExcelIncomingPort interface {
	Extract(filePath string) error
	Print()
}

type ExcelOutgoingPort interface {
	GetAll(filePath string) ([]RouteStore, error)
}
