package connection

import "github.com/jinzhu/gorm"

type connectionInterface interface {
	InitDB()
	GetDB() *gorm.DB
}
