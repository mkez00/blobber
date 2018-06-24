package services

import "github.com/mkez00/blobber/models"

type BaseBlobService interface {
	ListItems(config *models.Config) ([]models.Item, error)
	GetItem(config *models.Config, itemName string) (models.Item, error)
	PutItem(config *models.Config, filename string) (models.Item, error)
	DeleteItem(config *models.Config, obj string) (string, error)
}
