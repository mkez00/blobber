package main

import (
	"fmt"
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
		fmt.Println(config)
	} else if action == "version" {
		fmt.Println(getVersion())
	}
}

func processGetItem(service services.Base, config models.Config, args []string) {
	if len(args) > 2 {
		service.GetItem(config, args[2])
	} else {
		fmt.Println("\"blobber get\" requires at least 1 argument")
	}
}

func processDeleteItem(service services.Base, config models.Config, args []string) {
	if len(args) > 2 {
		service.DeleteItem(config, args[2])
	} else {
		fmt.Println("\"blobber delete\" requires at least 1 argument")
	}
}

func processPutItem(service services.Base, config models.Config, args []string) {
	if len(args) > 2 {
		item := service.PutItem(config, args[2])
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

func getImplementation(config models.Config) services.Base {
	if config.StorageService == "AmazonS3" {
		return services.AmazonS3{}
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
