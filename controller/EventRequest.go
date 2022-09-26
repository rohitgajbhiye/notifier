package controller

type EventRequest struct {
	AccountID string `json:"accountId" binding:"required"`
	EventType string `json:"eventType" binding:"required"`
}
