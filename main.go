package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	apiKey := os.Getenv("API_KEY")

	fmt.Println(apiKey)

}
