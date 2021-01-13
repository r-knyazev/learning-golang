package connection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (c *connection) InitDB() {

	//TODO вынести в конфиг
	host := "localhost"
	userName := "postgres"
	DBName := "products"
	password := "root"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, userName, DBName, password)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	//TODO сделать миграцию

	c.DB = conn
}

func (c *connection) GetDB() *gorm.DB {
	return c.DB
}