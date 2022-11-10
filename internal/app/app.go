package app

import (
	"github.com/AlexLevus/go-counter/internal/repository"
	"github.com/AlexLevus/go-counter/internal/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Counter struct {
	Value     int                `bson:"value,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
}

func Run() error {
	counterRepository, _ := repository.NewRepository()
	httpServer := server.NewServer(counterRepository)

	err := httpServer.Run("12345")

	if err != nil {
		return err
	}

	return nil
}
