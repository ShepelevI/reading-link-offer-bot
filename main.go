package main

import (
	"context"
	"flag"
	"log"
	tgClient "reading-link-offer-bot/clients/telegram"
	eventConsumer "reading-link-offer-bot/consumer/event-consumer"
	"reading-link-offer-bot/events/telegram"
	"reading-link-offer-bot/storage/sqlite"
)

const (
	tgBotHost         = "api..org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to  bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")

	}

	return *token
}
