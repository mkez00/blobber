package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	StorageService     string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsS3Bucket        string
	AwsRegion          string
}

func NewConfig() Config {
	config := Config{}
	return config
}

func LoadConfig() (Config, error) {
	configFilePath := GetConfigFilePath()

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		fmt.Println("Config file not created.  Run \"blobber init\" to create default template")
		return Config{}, err
	}

	var conf Config
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conf, nil
}

func GenerateDefaultConfigFile() {
	configFilePath := GetConfigFilePath()

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		CreateConfigFile(configFilePath)
	} else {
		fmt.Println("Default config already exists at ~/.blobber")
	}
}

func GetConfigFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.FromSlash(usr.HomeDir + "/.blobber")
}

func CreateConfigFile(filepath string) {

	str := "##########################################################\n" +
		"# GENERAL CONFIGURATION\n" +
		"##########################################################\n\n" +

		"StorageService = \"AmazonS3\"\n\n" +

		"##########################################################\n" +
		"# AWS CONFIGURATION\n" +
		"##########################################################\n\n" +

		"AwsAccessKeyId = \"\"\n" +
		"AwsSecretAccessKey = \"\"\n" +
		"AwsS3Bucket = \"\"\n" +
		"AwsRegion = \"\"\n"

	d1 := []byte(str)
	err := ioutil.WriteFile(filepath, d1, 0640)
	if err != nil {
		panic(err)
	}

	fmt.Println("IMPORTANT: Default configuration file generated at ~/.blobber. This file requires modification in order to use Blobber")
}
