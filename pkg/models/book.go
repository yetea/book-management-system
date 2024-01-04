package models

import (
	"github.com/jinzhu/gorm"

	"github.com/yetea/book-management-system/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `json:"name,omitempty"`
	Author      string `json:"author,omitempty"`
	Publication string `json:"publication,omitempty"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
