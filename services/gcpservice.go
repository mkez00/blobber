package services

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/mkez00/blobber/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type GoogleCloudStorage struct {
}

func getClient(ctx context.Context, config models.Config) (storage.Client, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(filepath.FromSlash(usr.HomeDir+"/"+config.GcpCredentialsFile)))
	if err != nil {
		return *client, err
	}

	return *client, nil
}

func (a GoogleCloudStorage) ListItems(config models.Config) []models.Item {

	ctx := context.Background()
	client, err := getClient(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
		return nil
	}

	// Creates a Bucket instance.
	bucket := client.Bucket(config.GcpBucketName)
	objectIterator := bucket.Objects(ctx, nil)
	items := []models.Item{}
	for {
		obj, err := objectIterator.Next()
		if err == iterator.Done {
			break
		}
		if obj == nil {
			break
		}
		item := models.Item{}
		item.Name = obj.Name
		items = append(items, item)
	}

	return items
}
func (a GoogleCloudStorage) PutItem(config models.Config, filename string) models.Item {
	item := models.Item{}

	ctx := context.Background()
	client, err := getClient(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
		return item
	}
	bucket := client.Bucket(config.GcpBucketName)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
		return item
	}
	defer file.Close()

	filename = filepath.Base(filename)
	writer := bucket.Object(filename).NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		log.Fatalf("Failed to write new object: %v", err)
		return item
	}
	if err := writer.Close(); err != nil {
		log.Fatalf("Failed to closer writer: %v", err)
		return item
	}

	item.Name = filename
	return item
}

func (a GoogleCloudStorage) GetItem(config models.Config, item string) {

	ctx := context.Background()
	client, err := getClient(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
		return
	}

	bucket := client.Bucket(config.GcpBucketName)
	reader, err := bucket.Object(item).NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to create new reader: %v", err)
		return
	}

	defer reader.Close()

	fileContent, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Failed to read data from bucket: %v", err)
		return
	}

	file, err := os.Create(item)
	if err != nil {
		log.Fatalf("Failed to create new file: %v", err)
		return
	}
	defer file.Close()

	numBytes, err := file.Write(fileContent)

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func (a GoogleCloudStorage) DeleteItem(config models.Config, obj string) {
	ctx := context.Background()
	client, err := getClient(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
		return
	}

	bucket := client.Bucket(config.GcpBucketName)
	err = bucket.Object(obj).Delete(ctx)
	if err != nil {
		fmt.Println("Error deleting object", err)
		return
	}

	fmt.Printf("Object %q successfully deleted\n", obj)
}
