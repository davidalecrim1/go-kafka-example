package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")

	admClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		log.Fatalf("Failed to create Kafka admin client: %s", err)
	}

	defer admClient.Close()

	var newTopics []kafka.TopicSpecification
	topics := createTopics()

	for _, topic := range topics {
		newTopics = append(newTopics, kafka.TopicSpecification{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := admClient.CreateTopics(ctx, newTopics)

	if err != nil {
		log.Fatalf("Failed to create topics: %v", err)
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			log.Printf("Failed to create topic %s: %s\n", result.Topic, result.Error.String())
		} else {
			log.Printf("Topic %s created successfully\n", result.Topic)
		}
	}

}

func createTopics() []string {
	ranges := [4][2]rune{
		{'a', 'g'},
		{'h', 'n'},
		{'o', 's'},
		{'t', 'z'},
	}

	var topics []string

	for _, r := range ranges {
		topic := fmt.Sprintf("topic-letters-%v-to-%v", string(r[0]), string(r[1]))
		topics = append(topics, topic)
	}

	return topics
}
