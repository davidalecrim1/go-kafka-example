# Go with Kafka Example

This repository demonstrates a Kafka setup using Go for producing and consuming messages. The project is structured to include separate directories for the producer, consumer, and a bootstrapper for setting up Kafka topics and configurations.

## Project Structure

### Root Directory

- `bootstrap` Directory

This directory contains the bootstrap code for setting up Kafka topics and initial configurations.

- `consumer` Directory

This directory contains the consumer code for consuming messages from Kafka topics.

- `producer` Directory

This directory contains the producer code for publishing messages to Kafka topics.

- `vendor` Directory (inside `bootstrap`, `consumer`, `producer`)

The `vendor` directories contain the dependencies required for each service. These are managed by Go modules and contain the necessary files from the Confluent Kafka Go client library. 

**Note:** You won't find the `vendor` folder in Github, it is not commited in the repository.

## Getting Started

To run the project, ensure you have **Docker** and **Docker Compose** installed. Follow these steps:

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/go-kafka-example.git
    cd go-kafka-example
    ```

2. **Build and run the services**:
    ```sh
    docker-compose up --build
    ```

This will set up Kafka, Zookeeper, and the Go services (producer, consumer, and bootstrap) based on the configuration in `docker-compose.yaml`.

## Usage
- The **bootstrap** service sets up Kafka topics.
- The **producer** service publishes messages to Kafka topics.
- The **consumer** service consumes messages from Kafka topics.

You can monitor the logs of each service to see the messages being produced and consumed.

## General Knowledge

### Definitions
**Zookeeper**
- **Definition**: Zookeeper is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services.
- **Purpose**: Zookeeper is used to manage and coordinate distributed systems. In Kafka, it helps manage brokers and topics, maintains metadata about the Kafka cluster, and helps with leader election for partitions. This ensures high availability, fault tolerance, and efficient load balancing in the Kafka ecosystem.

**Topics**
- **Definition**: A topic is a logical channel to which records are sent. It’s a category or feed name to which messages are published.
- **Purpose**: Topics allow you to categorize and segregate data streams. Consumers subscribe to topics to read messages.

**Partitions**
- **Definition**: A topic is divided into one or more partitions. Each partition is an ordered log of messages.
- **Purpose**: Partitions enable parallel processing and scalability. Each partition can be read by a single consumer in a consumer group, allowing for load balancing.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Licensec

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.