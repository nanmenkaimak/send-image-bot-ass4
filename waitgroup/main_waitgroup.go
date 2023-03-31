package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/nanmenkaimak/send-image-bot-ass4/keys"
	"github.com/nanmenkaimak/send-image-bot-ass4/random"
	"log"
	"os"
	"os/signal"
	"sync"
)

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

			photo, err := random.RandomPhoto(keys.AccessKey())
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
