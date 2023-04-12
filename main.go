package main

import (
	"flag"
	"log"
)

func main() {
	// tgClient = telegram.New()
	token := mustToken()
	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-telegram-bot", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("bot token isn't specified")
	}
	return *token
}
