package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("PASSWORD")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	fmt.Println(string(hashedPassword))
}
