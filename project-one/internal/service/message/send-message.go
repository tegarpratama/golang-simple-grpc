package message

import (
	"context"
	messageProto "project-one/cmd/proto/message"
	"project-one/internal/model/message"
	"project-one/pkg/connectgrpc"
)

func (s *service) SendMessage(ctx context.Context, req message.MessageReq) (*message.MessageRes, error) {
	address := "localhost:50051"
	connGRPC, err := connectgrpc.ConnectGRPCServer(address)
	if err != nil {
		return nil, err
	}

	defer connGRPC.Close()

	client := messageProto.NewMessageServiceClient(connGRPC)
	resp, err := client.SendMessage(context.Background(), &messageProto.MessageRequest{
		Content: req.Message,
	})

	if err != nil {
		return nil, err
	}

	result := &message.MessageRes{
		Status:  resp.Status,
		Message: resp.Content,
	}

	return result, nil
}
