package connection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func (c *connection) InitDB() {
	host 		:= os.Getenv("PG_HOST")
	user		:= os.Getenv("PG_USER")
	database	:= os.Getenv("PG_DATABASE")
	password 	:= os.Getenv("PG_PASSWORD")

	dbUri := fmt.Sprintf(os.Getenv("DB_URI_FORMAT"), host, user, database, password)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	//TODO сделать миграцию

	c.DB = conn
}

func (c *connection) GetDB() *gorm.DB {
	return c.DB
}