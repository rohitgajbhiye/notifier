package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"com.notifier/service"
	"github.com/gin-gonic/gin"
)

type EventController interface {
	ListenForEvent(ctx *gin.Context)
}

type eventController struct {
	relay service.RelayService
}

func NewEventController(relay service.RelayService) EventController {
	return &eventController{
		relay: relay,
	}
}

func (ctrl *eventController) ListenForEvent(ctx *gin.Context) {
	var request EventRequest
	validationErr := ctx.BindJSON(&request)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"ErrorMessage": validationErr.Error(),
			})
		return
	}
	log.Println("Received notification event :" + request.AccountID)
	jsonBytes, _ := json.Marshal(request)
	broadcastErr := ctrl.relay.Boradcast(ctx, "ACCOUNT_CREATED", jsonBytes)
	if broadcastErr != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"ErrorMessage": "error boradcasting message",
			})
		return
	}
	ctx.JSON(http.StatusCreated,
		gin.H{
			"status": "OK",
		},
	)
	log.Println("processed notification event :" + request.AccountID)
	return
}
