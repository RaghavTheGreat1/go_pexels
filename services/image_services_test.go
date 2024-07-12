package services_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/RaghavTheGreat1/go_pexels/models"

	"github.com/RaghavTheGreat1/go_pexels/services"
)

func TestGetPhoto(t *testing.T) {

	godotenv.Load(".env")
	apiKey := os.Getenv("API_KEY")

	client := models.PexelsClient{
		ApiKey: apiKey,
	}
	params := models.PhotoParams{
		Query: "Rocket",
	}

	result := services.SearchPhotos(&client, params)

	t.Log(result)

}
