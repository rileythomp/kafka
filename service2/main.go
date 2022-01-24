package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

var (
	topic   = "service2"
	brokers = []string{"localhost:9093", "localhost:9093", "localhost:9095"}
)

type Event struct {
	EntityType string
	Action     string
	Timestamp  int64
}

func main() {
	ctx := context.Background()
	produce(ctx)
}

func produce(ctx context.Context) {
	fmt.Println("starting service2...")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	for i := 0; ; i++ {
		event := Event{
			EntityType: "content",
			Action:     "created",
			Timestamp:  time.Now().Unix(),
		}
		bytes, _ := json.Marshal(event)

		w.WriteMessages(ctx, kafka.Message{
			Key:   []byte("service2-event"),
			Value: bytes,
		})
		fmt.Println("wrote a message")

		time.Sleep(5 * time.Second)
	}
}
