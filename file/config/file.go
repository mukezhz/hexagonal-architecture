package config

import (
	"github.com/mukezhz/hexagonal-architecture/file/application"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"github.com/mukezhz/hexagonal-architecture/file/infrastructure"
	"go.uber.org/fx"
)

var FileModule = fx.Module("file", fx.Options(
	fx.Provide(
		fx.Annotate(
			infrastructure.NewFileSystemAdapter,
			fx.As(new(domain.FileRepository))),
	),
	fx.Provide(
		fx.Annotate(
			application.NewFileUseCase,
			fx.As(new(domain.FilePort)),
		),
	),
	fx.Provide(application.NewProductController),
))
