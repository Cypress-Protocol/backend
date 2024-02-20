package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}
