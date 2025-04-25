package chat

import (
	"context"
	pb "snowApp/gen"
)

type ChatServer struct {
	chatService
}

func (s *ChatServer) CreateConversation(ctx context.Context, req *pb.CreateConversationRequest) error {
}
