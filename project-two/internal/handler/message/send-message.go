package message

import (
	"context"
	message "project-two/cmd/proto/message"
)

func (h *MessageHandler) SendMessage(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	result := h.MessageSvc.SendMessage(ctx, req.Content)

	return &message.MessageResponse{
		Status:  "Message received",
		Content: result,
	}, nil
}
