package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mkez00/blobber/models"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AmazonS3 struct {
}

func (a AmazonS3) ListItems(config models.Config) ([]models.Item, error) {

	sess, bucket := getSessionAndBucket(config)

	// Create S3 service client
	svc := s3.New(sess)
	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		return nil, err
	}

	items := []models.Item{}
	for _, s3item := range resp.Contents {
		item := models.Item{}
		item.Name = *s3item.Key
		item.FileSize = *s3item.Size
		item.LastModified = *s3item.LastModified
		item.StorageClass = *s3item.StorageClass
		items = append(items, item)
	}

	return items, nil
}

func (a AmazonS3) PutItem(config models.Config, filename string) (models.Item, error) {

	item := models.Item{}

	sess, bucket := getSessionAndBucket(config)

	file, err := os.Open(filename)
	if err != nil {
		return item, err
	}

	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	filename = filepath.Base(filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return item, err
	}

	item.Name = filename

	return item, nil
}

func (a AmazonS3) GetItem(config models.Config, itemName string) (models.Item, error) {

	item := models.Item{}

	sess, bucket := getSessionAndBucket(config)

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(itemName)
	if err != nil {
		return item, err
	}

	defer file.Close()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(itemName),
		})
	if err != nil {
		return item, err
	}

	item.Name = file.Name()
	item.FileSize = numBytes
	return item, nil
}

func (a AmazonS3) DeleteItem(config models.Config, obj string) (string, error) {
	sess, bucket := getSessionAndBucket(config)

	// Create S3 service client
	svc := s3.New(sess)

	// Delete the item
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj)})
	if err != nil {
		return "", err
	}

	svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})

	return obj, nil
}

func getSessionAndBucket(config models.Config) (*session.Session, string) {
	bucket := config.AwsS3Bucket
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(config.AwsRegion),
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyId,
			config.AwsSecretAccessKey,
			"")},
	)
	return sess, bucket
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
