package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/RaghavTheGreat1/go_pexels/utils"

	"github.com/RaghavTheGreat1/go_pexels/constants"

	"github.com/RaghavTheGreat1/go_pexels/models"
)

func SearchPhotos(c *models.PexelsClient, params models.PhotoParams) models.SearchPhotoResponse {

	query := utils.BuildQueryParams(params)

	url := fmt.Sprintf("%s/search?%s", constants.BaseURL, query.Encode())

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	photosResponse := models.SearchPhotoResponse{}
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &photosResponse)

	return photosResponse
}

func GetCuratedPhotos(c *models.PexelsClient, params models.PhotoParams) models.CuratedPhotosResponse {
	query := utils.BuildQueryParams(params)

	url := fmt.Sprintf("%s/curated?%s", constants.BaseURL, query.Encode())

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	curatedPhotosResponse := models.CuratedPhotosResponse{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &curatedPhotosResponse)

	return curatedPhotosResponse
}

func GetPhoto(c *models.PexelsClient, id string) models.PhotoResponse {

	url := fmt.Sprintf("%s/photos/%s", constants.BaseURL, id)

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	photoResponse := models.PhotoResponse{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &photoResponse)

	return photoResponse
}
