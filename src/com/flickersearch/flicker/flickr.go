package flicker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	flickerURL    = "https://api.flickr.com/services/rest/?"
	searchMethod  = "flickr.photos.search"
	thumbNailSize = "t"
	largeSize     = "b"
)

func buildURL(apiKey string, attr map[string]string) (string, error) {
	urlPath, err := url.Parse(flickerURL)
	if err != nil {
		return "", err
	}
	parameters := url.Values{}
	for k, v := range attr {
		parameters.Add(k, v)
	}
	urlPath.RawQuery = parameters.Encode()
	return urlPath.String(), nil

}

type imageType struct {
	ThumbNail string `json:"thumb"`
	Large     string `json:"large"`
}
type flickerImages struct {
	ImageURL []imageType `json:"images"`
	Pages    int         `json:"pages"`
	Page     int         `json:"page"`
}

// GetAllImages queries the flicker rest api based upon the search criteria and return images that was found
func GetAllImages(apiKey, search string, perPage uint64, page uint64) ([]byte, error) {
	attr := map[string]string{
		"method":   searchMethod,
		"api_key":  apiKey,
		"text":     search,
		"per_page": strconv.FormatUint(perPage, 10),
		"page":     strconv.FormatUint(page, 10),
		"format":   "json",
	}
	searchURL, err := buildURL(apiKey, attr)
	if err != nil {
		return nil, err
	}
	log.Printf("search query %s", searchURL)
	response, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	images, err := processResponse(body)
	if err != nil {
		return nil, err
	}

	return json.Marshal(images)

}

func processResponse(body []byte) (*flickerImages, error) {
	//need to remove non-json token
	bodyText := string(body)
	log.Printf("body response before removal %s", bodyText)
	bodyText = strings.Replace(bodyText, "jsonFlickrApi(", "", -1)
	bodyText = strings.Replace(bodyText, ")", "", -1)
	log.Printf("body response after removal %s", bodyText)
	decodedPhotoResponse, err := decodeSearchJSONData([]byte(bodyText))
	if err != nil {
		return nil, err
	}
	flickerImages := flickerImages{
		ImageURL: buildPhotoURL(decodedPhotoResponse.Photos.Photo),
		Page:     decodedPhotoResponse.Photos.Page,
		Pages:    decodedPhotoResponse.Photos.Pages,
	}
	return &flickerImages, nil
}

type flickerPhoto struct {
	ID       string `json:"id"`
	ServerID string `json:"server"`
	FarmID   int    `json:"farm"`
	Secret   string `json:"secret"`
}
type flickerPhotos struct {
	Photo []flickerPhoto `json:"photo"`
	Page  int            `json:"page"`
	Pages int            `json:"pages"`
}

type flickerSearchPhotoResponse struct {
	Photos flickerPhotos `json:"photos"`
	Status string        `json:"stat"`
}

func decodeSearchJSONData(data []byte) (*flickerSearchPhotoResponse, error) {
	response := &flickerSearchPhotoResponse{}
	err := json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}
	if response.Status != "ok" {
		return nil, errors.New(response.Status)
	}
	return response, nil

}

func buildPhotoURL(photos []flickerPhoto) []imageType {
	var imageURLs []imageType
	for _, p := range photos {
		//https://farm{farm-id}.staticflickr.com/{server-id}/{id}_{secret}_[mstzb].jpg
		thumb := fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s_%s.jpg",
			p.FarmID, p.ServerID, p.ID, p.Secret, thumbNailSize)
		large := fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s_%s.jpg",
			p.FarmID, p.ServerID, p.ID, p.Secret, largeSize)
		imageURLs = append(imageURLs, imageType{
			ThumbNail: thumb,
			Large:     large,
		})

	}
	return imageURLs

}
