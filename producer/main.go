package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	producer, err := NewKafkaProducer(broker)
	topic := os.Getenv("KAFKA_TOPIC")

	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}
	defer producer.Close()

	messages := []string{
		"apple", "beauty", "chocolate", "dream", "energy", "freedom", "gratitude",
		"harmony", "inspiration", "joy", "kindness", "light", "miracle", "nature",
		"opportunity", "peace", "quest", "respect", "serenity", "trust", "unity",
		"victory", "wisdom", "xylophone", "youth", "zeal",
	}

	ctx, cancel := context.WithCancel(context.Background())
	go awaitGracefulShutdown(cancel)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, message := range messages {
				producer.ProduceMessage(topic, []byte(message), getEventPartition(message))
			}
			time.Sleep(time.Second * 10)
		}
	}

}

func getEventPartition(message string) int32 {
	firstLetter := rune(message[0])

	switch {
	case firstLetter >= 'a' && firstLetter <= 'g':
		return 0
	case firstLetter >= 'h' && firstLetter <= 'n':
		return 1
	case firstLetter >= 'o' && firstLetter <= 's':
		return 2
	case firstLetter >= 't' && firstLetter <= 'z':
		return 3
	default:
		return kafka.PartitionAny
	}
}

func awaitGracefulShutdown(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Wait for an interrupt signal
	<-sigChan
	fmt.Println("\nreceived an interrupt, starting graceful shutting down...")
	cancel()
}
