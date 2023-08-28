package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAWSConfig),
	fx.Provide(NewDynamoDBClient),
	fx.Provide(NewS3Client),
	fx.Provide(NewS3Uploader),
	fx.Provide(NewS3PresignClient),
	fx.Provide(NewMysqlDB),
)
