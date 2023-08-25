package config

import (
	"github.com/mukezhz/hexagonal-architecture/file/application"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"github.com/mukezhz/hexagonal-architecture/file/infrastructure"
	"go.uber.org/fx"
)

var FileModule = fx.Module("file", fx.Options(
	// for file
	fx.Provide(
		fx.Annotate(
			infrastructure.NewFileSystemAdapter,
			fx.As(new(domain.FileOutgoingPort))),
	),
	fx.Provide(
		fx.Annotate(
			application.NewFileUseCase,
			fx.As(new(domain.FilePort)),
		),
	),
	fx.Provide(
		fx.Annotate(
			infrastructure.NewMysqlRepository,
			//infrastructure.NewDynamoDBRepository,
			fx.As(new(domain.FileRepository)),
		),
	),
	fx.Provide(application.NewFileController),
))

var ExcelModule = fx.Module("excel", fx.Options(
	// for excel
	fx.Provide(
		fx.Annotate(
			//infrastructure.NewMysqlRepository,
			infrastructure.NewDynamoDBRepository,
			fx.As(new(domain.ExcelRepository)),
		),
	),
	fx.Provide(
		fx.Annotate(
			infrastructure.NewFileSystemAdapter,
			fx.As(new(domain.ExcelOutgoingPort)),
		),
	),
	fx.Provide(
		fx.Annotate(
			application.NewExcelUseCase,
			fx.As(new(domain.ExcelIncomingPort)),
		),
	),
))
