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

func SearchVideos(c *models.PexelsClient, params models.VideoParams) models.SearchVideoResponse {

	query := utils.BuildQueryParams(params)

	url := fmt.Sprintf("%s/search?%s", constants.VideoBaseURL, query.Encode())

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	searchVideosResponse := models.SearchVideoResponse{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &searchVideosResponse)

	return searchVideosResponse
}

func GetPopularVideos(c *models.PexelsClient, params models.VideoParams) models.PopularVideoResponse {
	query := utils.BuildQueryParams(params)

	url := fmt.Sprintf("%s/popular?%s", constants.VideoBaseURL, query.Encode())

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	popularVideosResponse := models.PopularVideoResponse{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &popularVideosResponse)

	return popularVideosResponse
}

func GetVideo(c *models.PexelsClient, params models.VideoParams) models.Video {

	url := fmt.Sprintf("%s/videos/%s", params.Id)

	newReq, _ := http.NewRequest("GET", url, nil)

	newReq.Header.Add("Authorization", c.ApiKey)
	newReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(newReq)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	videoResponse := models.Video{}
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &videoResponse)

	return videoResponse
}
