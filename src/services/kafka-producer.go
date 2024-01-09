package services

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const TIPS_CREATED_TOPIC = "tips.created"

type KafkaMessage struct {
	Data string `json:"data"`
}

func emitToKafka(message KafkaMessage) {
	kafkaBroker := os.Getenv("KAFKA_BROKER")

	conn, _ := kafka.DialLeader(context.Background(), "tcp", kafkaBroker, TIPS_CREATED_TOPIC, 0)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	jsonMsg, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	_, err = conn.WriteMessages(
		kafka.Message{Value: jsonMsg},
	)

	if err != nil {
		panic(err)
	}
}
