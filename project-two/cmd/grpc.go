package cmd

import (
	"log"
	"net"

	messageProto "project-two/cmd/proto/message"
	messageHandler "project-two/internal/handler/message"
	messageRepository "project-two/internal/repository/message"
	messageSvc "project-two/internal/service/message"

	"google.golang.org/grpc"
)

func ServeGRPC() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	messageRepository := messageRepository.NewRepository()
	messageService := messageSvc.NewService(messageRepository)
	messageServer := messageHandler.NewHandler(messageService)

	messageProto.RegisterMessageServiceServer(s, messageServer)

	log.Println("gRPC server running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
