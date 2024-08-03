package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	groupID := os.Getenv("KAFKA_GROUP_ID")
	topic := os.Getenv("KAFKA_TOPIC")

	consumer, err := NewKafkaConsumer(broker, groupID)

	if err != nil {
		log.Fatalf("failed to create kafka consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic})

	if err != nil {
		log.Fatalf("failed to subscribe to topic '%v': %v", topic, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go awaitGracefulShutdown(cancel)

	defer cancel() // Ensures the context is properly cleaned up when main exits
	err = consumer.ConsumeMessages(ctx)

	if err != nil {
		log.Fatalf("failed to consume messages: %v", err)
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
