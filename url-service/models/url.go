package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ShortCode string `gorm:"unique" json:"short_code"`
	Original  string `json:"original_url"`
}
