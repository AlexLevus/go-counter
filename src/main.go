package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-martini/martini"
)

type Counter struct {
	Value     int                `bson:"value,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	collection := client.Database("CounterDB").Collection("Counter")

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var counters []Counter
	if err = cur.All(ctx, &counters); err != nil {
		panic(err)
	}

	counter := counters[0]

	fmt.Println(counter.UpdatedAt)

	m := martini.Classic()

	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Счетчик равен %v\n. Счетчик был обновлен %v\n", counter.Value, counter.UpdatedAt.Time())
	})

	m.Get("/stat", func(res http.ResponseWriter, req *http.Request) {
		id, _ := primitive.ObjectIDFromHex("63692f15b50ce6ea336f9139")
		filter := bson.D{{"_id", id}}

		updatedAt := time.Now()

		update := bson.M{
			"$set": bson.M{
				"value":     counter.Value + 1,
				"updatedAt": updatedAt,
			},
		}

		_, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			panic(err)
		}

		counter.Value = counter.Value + 1
		counter.UpdatedAt = primitive.NewDateTimeFromTime(updatedAt)
		fmt.Fprintf(res, "Счетчик равен %v\n", counter.Value)
	})

	m.Get("/about", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(res, "<h3> Hello, Александр Левусь</h3>")
	})

	m.RunOnAddr(":12345")
}
