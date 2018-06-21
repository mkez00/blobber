package services

import "github.com/mkez00/blobber/models"

type Base interface {
	ListItems(config models.Config) []models.Item
	GetItem(config models.Config, item string)
	PutItem(config models.Config, filename string) models.Item
	DeleteItem(config models.Config, obj string)
}
