package main

import (
	"context"
	"fmt"
	"sync"

	kafka "github.com/segmentio/kafka-go"
)

var (
	topics  = []string{"service1", "service2"}
	brokers = []string{"localhost:9093", "localhost:9093", "localhost:9095"}
)

func main() {
	fmt.Println("starting audit service...")
	ctx := context.Background()
	consume(ctx, brokers, topics)
}

func consume(ctx context.Context, brokers []string, topics []string) {
	var wg sync.WaitGroup

	for _, topic := range topics {
		topicReader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:     brokers,
			Topic:       topic,
			GroupID:     "audit-group",
			StartOffset: kafka.LastOffset,
		})

		wg.Add(1)
		go func(r *kafka.Reader) {
			defer wg.Done()
			for {
				// ReadMessage blocks until next event is received
				msg, err := r.ReadMessage(ctx)
				if err != nil {
					panic("could not read message " + err.Error())
				}
				fmt.Printf("received message from %s: %s\n", msg.Topic, string(msg.Value))
			}
		}(topicReader)
	}

	fmt.Printf("listening for messages from topics: %v\n", topics)
	wg.Wait()
}
