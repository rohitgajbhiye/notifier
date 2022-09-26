package service

import (
	"context"
	"log"

	"github.com/rohitgajbhiye/kafka_util/producer"
)

type RelayService interface {
	Boradcast(context.Context, string, []byte) error
}

type relayService struct {
	producer producer.Producer
}

func NewRelayService(producer producer.Producer) RelayService {
	return &relayService{
		producer: producer,
	}
}

func (r *relayService) Boradcast(ctx context.Context, topic string, value []byte) error {
	_, err := r.producer.SendMessage(topic, value)
	if err != nil {
		log.Println("relayService: error while sending message,=", err)
		return err
	}
	return nil
}
