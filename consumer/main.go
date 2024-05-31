package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
    ctx := context.Background()
    topic := "logs"

    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic: topic,
        Partition: 0, // Read only the error logs
    })

    for {
        msg, err := reader.ReadMessage(ctx)
        if err != nil {
            panic(err)
        }

        log.Println(string(msg.Key))
        log.Println(string(msg.Value))
    }
}
