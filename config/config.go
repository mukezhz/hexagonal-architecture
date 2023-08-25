package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAWSConfig),
	fx.Provide(NewDynamoDBClient),
	fx.Provide(NewMysqlDB),
)
