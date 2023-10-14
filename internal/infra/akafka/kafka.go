package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, server string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "lista-de-produtos",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(topics, nil)
	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err != nil {
			msgChan <- msg
		}
	}
}
