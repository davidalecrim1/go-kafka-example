package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")

	client, err := NewKafkaAdminClient(broker)
	if err != nil {
		log.Fatalf("failed to create Kafka admin client: %s", err)
	}
	defer client.Close()

	topic := NewKafkaTopic(os.Getenv("KAFKA_TOPIC"), 4, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = CreateTopicsInKafka(ctx, client, []kafka.TopicSpecification{topic})
	if err != nil {
		log.Fatalf("failed to create topics: %v", err)
	}

}

func NewKafkaAdminClient(broker string) (*kafka.AdminClient, error) {
	client, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateTopicsInKafka(ctx context.Context, client *kafka.AdminClient, topics []kafka.TopicSpecification) error {
	results, err := client.CreateTopics(ctx, topics)

	if err != nil {
		return err
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			log.Printf("failed to create topic %s: %s\n", result.Topic, result.Error.String())
		} else {
			log.Printf("topic '%s' created successfully\n", result.Topic)
		}
	}

	return nil
}

func NewKafkaTopic(name string, partitions int, replicationFactor int) kafka.TopicSpecification {
	return kafka.TopicSpecification{
		Topic:             name,
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	}
}
