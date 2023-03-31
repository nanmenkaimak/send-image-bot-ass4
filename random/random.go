package random

import (
	"encoding/json"
	"fmt"
	"github.com/nanmenkaimak/send-image-bot-ass4/photo"
	"io/ioutil"
	"net/http"
)

func RandomPhoto(accessKey string) (photo.Photo, error) {
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s", accessKey)

	resp, err := http.Get(url)
	if err != nil {
		return photo.Photo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return photo.Photo{}, err
	}

	var photos photo.Photo
	err = json.Unmarshal(body, &photos)
	if err != nil {
		return photo.Photo{}, err
	}

	return photos, nil
}
