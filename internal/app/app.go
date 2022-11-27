package app

import (
	"github.com/AlexLevus/go-counter/internal/repository"
	"github.com/AlexLevus/go-counter/internal/server"
	"os"
)

func Run() error {
	counterRepository, _ := repository.NewRepository()
	httpServer := server.NewServer(counterRepository)

	err := httpServer.Run(os.Getenv("PORT"))

	if err != nil {
		return err
	}

	return nil
}
