package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka_kafka_1:9092",
		// Identificador do consumer
		"client.id": "go-consumer",
		// Identificador de grupo
		"group.id": "go-group",
		// Pegar todas as mensagens enviadas para um tópico desde o início e reprocessar
		// "auto.offset.reset":"earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"test"}
	c.SubscribeTopics(topics, nil)
	// loop infinito
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
