package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := newKafkaProducer()
	Publish("mensagem 6", "test", producer, []byte("Transferencia"), deliveryChan)
	go DeliveryReport(deliveryChan)
	// e := <-deliveryChan
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Erro ao enviar")
	// } else {
	// 	fmt.Println("Mensagem enviada", msg.TopicPartition)
	// }
	producer.Flush(5000)

}
func newKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "go-kafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		// Garantir que todas as mensagens chegaram, caso o contrário desparar um erro
		"acks": "all",
		// Não repetir uma mesma mensagem dentro de um tópico
		"enable.idempotence": "true",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

// Mostra se a mensagem foi entregue com sucesso
func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada", ev.TopicPartition)
				// Anotar que uma mensagem foi enviada no banco de dados
			}

		}
	}
}
