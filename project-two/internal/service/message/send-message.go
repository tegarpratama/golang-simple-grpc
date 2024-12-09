package message

import "context"

func (s *service) SendMessage(ctx context.Context, message string) string {
	return message
}
