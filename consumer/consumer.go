package main

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(broker string, groupID string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	kc := &KafkaConsumer{
		consumer: c,
	}

	return kc, nil
}

func (kp *KafkaConsumer) Close() {
	kp.consumer.Close()
}

func (kp *KafkaConsumer) SubscribeTopics(topics []string) error {
	err := kp.consumer.SubscribeTopics(topics, nil)

	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	return nil
}

func (kp *KafkaConsumer) ConsumeMessages(ctx context.Context) error {
	log.Printf("starting to consume messages...")

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := kp.consumer.ReadMessage(-1)

			if err != nil {
				log.Printf("error reading message: %v\n", err)
				continue
			}

			log.Printf("consumed message '%v' from topic '%v' in partition '%v' @ offset %v\n",
				string(msg.Value), *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
		}
	}
}
