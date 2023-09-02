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

func GetFeaturedCollections(c *models.PexelsClient, params models.CollectionParams) models.CollectionResponse {

	query := utils.BuildQueryParams(params)

	url := fmt.Sprintf("%s/featured?%s", constants.CollectionBaseURL, query.Encode())

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	featuredCollectionResponse := models.CollectionResponse{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &featuredCollectionResponse)

	return featuredCollectionResponse
}

func GetMyCollections(c *models.PexelsClient, params models.CollectionParams) models.CollectionResponse {

	query := utils.BuildQueryParams(params)

	url := fmt.Sprintf("%s?%s", constants.CollectionBaseURL, query.Encode())

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	myCollectionResponse := models.CollectionResponse{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &myCollectionResponse)

	return myCollectionResponse
}

func GetCollectionMedia(c *models.PexelsClient, id string, params models.CollectionParams) models.CollectionMediaResponse {

	query := utils.BuildQueryParams(params)

	url := fmt.Sprintf("%s/%s?%s", constants.CollectionBaseURL, id, query.Encode())

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	collectionMediaResponse := models.CollectionMediaResponse{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &collectionMediaResponse)

	return collectionMediaResponse
}
