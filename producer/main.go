package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type CustomBalancer struct {
	// We use the builtin round robbin balancer
	baseBalancer kafka.RoundRobin
}

func (b *CustomBalancer) Balance(msg kafka.Message, partitions ...int) (partition int) {
	// only error logs are sent to first partition
	if string(msg.Key) == "error" {
		return 0
	}

    // everything else get round robined to the other partitions.
	return b.baseBalancer.Balance(msg, partitions[1:]...)
}

func main() {
	ctx := context.Background()
	topic := "logs"

	conn, err := kafka.DialContext(ctx, "tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err)
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
        Balancer: &CustomBalancer{}, // Use our custom balancer.
	})

	err = writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("error"),
			Value: []byte("some error happened in the application"),
		}, kafka.Message{
			Key:   []byte("warn"),
			Value: []byte("this is a warning"),
		}, kafka.Message{
			Key:   []byte("info"),
			Value: []byte("this is an info"),
		})
	if err != nil {
		panic(err)
	}

	log.Println("written messages successfully")
}
