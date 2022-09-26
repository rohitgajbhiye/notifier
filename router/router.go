package router

import (
	"com.notifier/controller"
	"com.notifier/service"
	"github.com/gin-gonic/gin"
	"github.com/rohitgajbhiye/kafka_util/producer"
)

func InitRoute(engine *gin.Engine) {
	// We are passing nil Sarama config for now
	producer := producer.NewProducer([]string{"kafka:9092"}, nil)
	relayService := service.NewRelayService(producer)
	eventControler := controller.NewEventController(relayService)
	engine.POST("/event", eventControler.ListenForEvent)
}
