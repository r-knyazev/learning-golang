package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init()  {
	goPath, _ := os.LookupEnv("GOPATH")

	err := godotenv.Load(goPath + "/learning/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}