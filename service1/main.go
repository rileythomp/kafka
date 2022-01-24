package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

var (
	topic   = "service1"
	brokers = []string{"localhost:9093", "localhost:9093", "localhost:9095"}
)

type Event struct {
	EntityType string
	Action     string
	Timestamp  int64
}

func main() {
	fmt.Println("starting service1...")
	ctx := context.Background()
	produce(ctx)
}

func produce(ctx context.Context) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	for {
		event := Event{
			EntityType: "push",
			Action:     "sent",
			Timestamp:  time.Now().Unix(),
		}
		bytes, _ := json.Marshal(event)

		// key is used to decide which partition (and thus broker) the message gets published on
		// using the same key guarantees in order messages
		w.WriteMessages(ctx, kafka.Message{
			Key:   []byte("service1-event"),
			Value: bytes,
		})
		fmt.Println("wrote a message")

		time.Sleep(time.Second)
	}
}
