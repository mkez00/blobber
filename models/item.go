package models

import "time"

type Item struct {
	Name         string
	FileSize     int64
	LastModified time.Time
	StorageClass string
}
