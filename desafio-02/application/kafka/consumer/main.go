package consumer

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaConsumer(producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaConsumer {
	return &KafkaConsumer{
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}
func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "desafio-2",
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"Teste"}
	c.SubscribeTopics(topics, nil)

	fmt.Println("Consumer iniciado...")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println("->", string(msg.Value))
		}
	}
}
