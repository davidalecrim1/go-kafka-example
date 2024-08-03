package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	events   chan kafka.Event
}

func NewKafkaProducer(broker string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	kp := &KafkaProducer{
		producer: p,
		events:   p.Events(),
	}

	go kp.handleEvents()
	return kp, nil
}

func (kp *KafkaProducer) ProduceMessage(topic string, message []byte, partition int32) error {
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
		Value:          message,
	}, nil)
}

func (kp *KafkaProducer) Close() {
	kp.producer.Flush(15 * 1000)
	kp.producer.Close()
}

func (kp *KafkaProducer) handleEvents() {
	// goroutine is blocked while there are not events in the beginning of for.
	// only ends when the channel is closed.
	for e := range kp.events {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Printf("failed to deliver message: %v\n", ev.TopicPartition)
			} else {
				log.Printf("produced message '%v' to topic '%s' in partition '%d' @ offset %v\n", string(ev.Value),
					*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
			}
		}
	}
}
