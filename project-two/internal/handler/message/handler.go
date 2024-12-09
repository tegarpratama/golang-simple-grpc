package message

import (
	"context"

	messageProto "project-two/cmd/proto/message"
)

type service interface {
	SendMessage(ctx context.Context, message string) string
}

type MessageHandler struct {
	MessageSvc service
	messageProto.UnimplementedMessageServiceServer
}

func NewHandler(messageSvc service) *MessageHandler {
	return &MessageHandler{
		MessageSvc: messageSvc,
	}
}
