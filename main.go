package main

import (
	"log"

	"github.com/kaashmonee/signallm/internal/bot"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	b, err := bot.NewBot()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	b.Start()
}
