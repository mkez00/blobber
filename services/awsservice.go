package services

import (
	"blobber/models"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AmazonS3 struct {
}

func (a AmazonS3) ListItems(config models.Config) []models.Item {

	sess, bucket := getSessionAndBucket(config)

	// Create S3 service client
	svc := s3.New(sess)
	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
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

	return items
}

func (a AmazonS3) PutItem(config models.Config, filename string) models.Item {
	item := models.Item{}

	sess, bucket := getSessionAndBucket(config)

	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	item.Name = filename

	return item
}

func (a AmazonS3) GetItem(config models.Config, item string) {
	sess, bucket := getSessionAndBucket(config)

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(item)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func (a AmazonS3) DeleteItem(config models.Config, obj string) {
	sess, bucket := getSessionAndBucket(config)
	// Create S3 service client
	svc := s3.New(sess)

	// Delete the item
	svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj)})

	svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})

	fmt.Printf("Object %q successfully deleted\n", obj)
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
