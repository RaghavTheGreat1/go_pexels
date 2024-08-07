package main

import (
	"fmt"
	"os"

	"github.com/RaghavTheGreat1/go_pexels/services"

	"github.com/RaghavTheGreat1/go_pexels/models"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	apiKey := os.Getenv("API_KEY")

	client := models.PexelsClient{
		ApiKey: apiKey,
	}
	params := models.PhotoParams{
		Query: "Rocket",
	}
	result := services.SearchPhotos(&client, params)

	fmt.Print(result)
}
