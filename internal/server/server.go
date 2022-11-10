package server

import (
	"github.com/AlexLevus/go-counter/internal/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *gin.Engine
}

func NewServer(rep *repository.Repository) *Server {
	counter, err := rep.GetCounter()
	if err != nil {
		log.Fatalf("Error when get counter")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Счетчик равен %v\n. Счетчик был обновлен %v\n", counter.Value, counter.UpdatedAt.Time())
	})

	r.GET("/stat", func(c *gin.Context) {
		err := rep.UpdateCounter(counter)
		if err != nil {
			log.Fatalf("Error when update counter")
		}

		updatedAt := time.Now()
		counter.Value = counter.Value + 1
		counter.UpdatedAt = primitive.NewDateTimeFromTime(updatedAt)
		c.String(http.StatusOK, "Счетчик равен %v\n", counter.Value)
	})

	r.GET("/about", func(c *gin.Context) {
		r.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "about.tmpl", gin.H{
			"name": "Александр Левусь",
		})
	})

	return &Server{httpServer: r}
}

func (s *Server) Run(port string) error {
	return s.httpServer.Run(":" + port)
}
