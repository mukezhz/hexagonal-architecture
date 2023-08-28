package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

func NewAWSConfig() aws.Config {
	// get config from environment variables
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")

	// setup aws credential provider
	credProvider := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		awsAccessKey, awsSecretAccessKey, "",
	))

	conf, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credProvider),
	)

	if err != nil {
		panic(err)
	}
	return conf
}

func NewDynamoDBClient(sdkConfig aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(sdkConfig)
}

func NewS3Client(sdkConfig aws.Config) *s3.Client {
	return s3.NewFromConfig(sdkConfig)
}

func NewS3Uploader(client *s3.Client) *manager.Uploader {
	return manager.NewUploader(client)
}

func NewS3PresignClient(client *s3.Client) *s3.PresignClient {
	return s3.NewPresignClient(client)
}
