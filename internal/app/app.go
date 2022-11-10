package app

import (
	"github.com/AlexLevus/go-counter/internal/repository"
	"github.com/AlexLevus/go-counter/internal/server"
)

func Run() error {
	counterRepository, _ := repository.NewRepository()
	httpServer := server.NewServer(counterRepository)

	err := httpServer.Run("12345")

	if err != nil {
		return err
	}

	return nil
}
