package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mkez00/blobber/models"
	"github.com/mkez00/blobber/services"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// if no arguments passed print help
	if len(os.Args) <= 1 {
		printHelp()
		return
	}

	action := os.Args[1]
	if action == "init" {
		models.GenerateDefaultConfigFile()
		return
	}

	config, err := models.LoadConfig()
	if err != nil {
		return
	}
	service := getImplementation(config)

	switch strings.ToLower(action) {
	case "list":
		processListItems(service, &config)
	case "put":
		processPutItem(service, &config, os.Args)
	case "get":
		processGetItem(service, &config, os.Args)
	case "delete":
		processDeleteItem(service, &config, os.Args)
	case "help":
		printHelp()
	case "config":
		fmt.Printf("%+v\n", config)
	case "version":
		fmt.Println(getVersion())
	default:
		fmt.Println(action + " is not a valid argument")
	}
}

func processListItems(service services.BaseBlobService, config *models.Config) {
	items, err := service.ListItems(config)
	if err != nil {
		log.Fatalf("Failed to retrieve items from bucket: %v", err)
		return
	}
	printItems(items)
}

func processGetItem(service services.BaseBlobService, config *models.Config, args []string) {
	if len(args) > 2 {
		item, err := service.GetItem(config, args[2])
		if err != nil {
			log.Fatalf("Failed to get item from bucket: %v", err)
			return
		}
		fmt.Println("Downloaded", item.Name, item.FileSize, "bytes")
	} else {
		fmt.Println("\"blobber get\" requires at least 1 argument")
	}
}

func processDeleteItem(service services.BaseBlobService, config *models.Config, args []string) {
	if len(args) > 2 {
		msg, err := service.DeleteItem(config, args[2])
		if err != nil {
			log.Fatalf("Failed to delete item from bucket: %v", err)
			return
		}
		fmt.Printf("Object %q successfully deleted\n", msg)
	} else {
		fmt.Println("\"blobber delete\" requires at least 1 argument")
	}
}

func processPutItem(service services.BaseBlobService, config *models.Config, args []string) {
	if len(args) > 2 {
		item, err := service.PutItem(config, args[2])
		if err != nil {
			log.Fatalf("Failed to put item into bucket: %v", err)
			return
		}
		fmt.Println(item.Name + " successfully put")
	} else {
		fmt.Println("\"blobber put\" requires at least 1 argument")
	}
}

func printItems(items []models.Item) {
	for _, item := range items {
		fmt.Println(item.Name)
	}
}

func getImplementation(config models.Config) services.BaseBlobService {
	switch strings.ToLower(config.StorageService) {
	case "amazons3":
		return services.AmazonS3{}
	case "googlecloudstorage":
		return services.GoogleCloudStorage{}
	default:
		return services.AmazonS3{}
	}
}

func printHelp() {
	fmt.Println("Usage:  blobber COMMAND")
	fmt.Println()
	fmt.Println("Wrapper for cloud storage cli to allow for accessing blob storage anywhere")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  config\tOutput config file (~/.blobber)")
	fmt.Println("  delete\tDelete item from storage")
	fmt.Println("  get\t\tRetrieve item from storage")
	fmt.Println("  list\t\tList all items in defined bucket")
	fmt.Println("  put\t\tPut item in storage")
	fmt.Println("  version\tGet Blobber version")

}

func getVersion() string {
	return "0.0.1"
}
