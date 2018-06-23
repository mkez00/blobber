package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mkez00/blobber/models"
	"github.com/mkez00/blobber/services"
)

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

	if action == "list" {
		items := service.ListItems(config)
		printItems(items)
	} else if action == "put" {
		processPutItem(service, config, os.Args)
	} else if action == "get" {
		processGetItem(service, config, os.Args)
	} else if action == "delete" {
		processDeleteItem(service, config, os.Args)
	} else if action == "help" {
		printHelp()
	} else if action == "config" {
		fmt.Printf("%+v\n", config)
	} else if action == "version" {
		fmt.Println(getVersion())
	} else {
		fmt.Println(action + " is not a valid argument")
	}
}

func processGetItem(service services.BaseBlobService, config models.Config, args []string) {
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

func processDeleteItem(service services.BaseBlobService, config models.Config, args []string) {
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

func processPutItem(service services.BaseBlobService, config models.Config, args []string) {
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
	if config.StorageService == "AmazonS3" {
		return services.AmazonS3{}
	} else if config.StorageService == "GoogleCloudStorage" {
		return services.GoogleCloudStorage{}
	} else {
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
