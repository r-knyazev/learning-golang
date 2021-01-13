package connection

import (
	"github.com/jinzhu/gorm"
)

var Connection connectionInterface

type connection struct {
	DB *gorm.DB
}

func init() {
	Connection = newConnection()
	Connection.InitDB()
}

func newConnection() connectionInterface {
	return &connection{}
}