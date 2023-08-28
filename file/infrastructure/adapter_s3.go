package infrastructure

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"log"
	"mime/multipart"
	"time"
)

type S3Adapter struct {
	Client   *s3.Client
	uploader *manager.Uploader
	presign  *s3.PresignClient
}

func NewS3Adapter(client *s3.Client, uploader *manager.Uploader, presign *s3.PresignClient) *S3Adapter {
	return &S3Adapter{
		Client:   client,
		uploader: uploader,
		presign:  presign,
	}
}

func (fs3 *S3Adapter) Save(file *multipart.FileHeader, bucketName string) (string, error) {
	uploadFile, err := file.Open()
	if err != nil {
		return "", err
	}
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
		Body:   uploadFile,
	}
	result, err := fs3.uploader.Upload(context.TODO(), input)
	if err != nil {
		return "", err
	}
	return result.Location, nil
}

func (fs3 *S3Adapter) SavePublicly(file *multipart.FileHeader, bucketName string) (string, error) {
	uploadFile, err := file.Open()
	if err != nil {
		return "", err
	}
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
		Body:   uploadFile,
		ACL:    types.ObjectCannedACLPublicRead,
	}

	result, err := fs3.uploader.Upload(context.TODO(), input)
	if err != nil {
		return "", err
	}
	return result.Location, nil
}

// GetSignedURL get the signed url for file
func (fs3 *S3Adapter) GetSignedURL(file *multipart.FileHeader, bucketName string, expires *time.Time) (string, error) {
	if expires == nil {
		e := time.Now().Add(24 * time.Hour)
		expires = &e
	}
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
	}
	duration := time.Until(*expires)

	resp, err := fs3.presign.PresignGetObject(context.Background(), input, s3.WithPresignExpires(duration))
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}

func (fs3 *S3Adapter) GetAll(dst, filePath string) ([]domain.RouteStore, error) {
	routeStores := make([]domain.RouteStore, 0)

	object, err := fs3.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(dst),
		Key:    aws.String(filePath),
	})
	if err != nil {
		return routeStores, err
	}
	log.Println("OBJECT:", object)
	return routeStores, nil
}
