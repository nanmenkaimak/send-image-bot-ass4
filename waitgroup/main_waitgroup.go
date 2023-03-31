package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/nanmenkaimak/send-image-bot-ass4/keys"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type Photo struct {
	ID   string `json:"id"`
	URLs struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := bot.New(keys.Token(), bot.WithDefaultHandler(handler))
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var wg sync.WaitGroup

	if update.Message.Text == "image" || update.Message.Text == "/image" {
		wg.Add(1)

		go func(chatID int64) {
			defer wg.Done()

			photo, err := randomPhoto(keys.AccessKey())
			if err != nil {
				log.Fatal(err)
			}

			params := &bot.SendPhotoParams{
				ChatID: update.Message.Chat.ID,
				Photo:  &models.InputFileString{Data: photo.URLs.Regular},
			}
			if _, err := b.SendPhoto(ctx, params); err != nil {
				log.Println(err)
			}
		}(update.Message.Chat.ID)
	}

	wg.Wait()
}

func randomPhoto(accessKey string) (Photo, error) {
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s", accessKey)

	resp, err := http.Get(url)
	if err != nil {
		return Photo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Photo{}, err
	}

	var photos Photo
	err = json.Unmarshal(body, &photos)
	if err != nil {
		return Photo{}, err
	}

	return photos, nil
}
