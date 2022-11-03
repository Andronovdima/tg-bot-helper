package main

import (
	"log"

	"github.com/Andronovdima/tg-bot-helper/internal/anci_bot/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
