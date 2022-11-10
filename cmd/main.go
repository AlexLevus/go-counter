package main

import (
	"log"

	"github.com/AlexLevus/go-counter/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("app.Run: %v", err)
	}
}
