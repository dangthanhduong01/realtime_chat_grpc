package controller

import (
	"context"
	pb "snowApp/gen"
	"snowApp/internal/service"
	"snowApp/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatController struct {
	chatService *service.ChatService
	userService *service.AuthService
	logger      *logger.Logger
}

func NewChatController(cs *service.ChatService, us *service.AuthService, log *logger.Logger) *ChatController {
	return &ChatController{
		chatService: cs,
		userService: us,
		logger:      log,
	}
}
func (c *ChatController) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	// validate the request
	if req.SenderId == "" || req.ConversationId == "" {
		return nil, status.Error(codes.InvalidArgument, "SenderId and ConversationId are required")
	}
	msg, err := c.chatService.CreateMessage(ctx, req)
	if err != nil {
		c.logger.Error("Failed to send message:", err)
		return nil, err
	}

	if err := c.chatService.BroadcastMessage(ctx, msg); err != nil {
		c.logger.Error("Failed to broadcast message:", err)
		return nil, err
	}

	c.logger.Info("Message sent successfully:", msg.ID)
	return &pb.SendMessageResponse{
		MessageId: msg.ID.Hex(),
		SentAt:    timestamppb.New(msg.SentAt),
	}, nil
}

// func (c *ChatController) GetConversations(ctx context.Context, req *pb.GetConversationsRequest) (*pb.GetConversationsResponse, error) {
// 	// validate the request
// 	if req.UserId == "" {
// 		return nil, status.Error(codes.InvalidArgument, "UserId is required")
// 	}

// 	conversations, err := c.chatService.GetConversations(ctx, req.UserId)
// 	if err != nil {
// 		c.logger.Error("Failed to get conversations:", err)
// 		return nil, err
// 	}

// 	var pbConversations []*pb.Conversation
// 	for _, conv := range conversations {
// 		pbConv := &pb.Conversation{
// 			Id:          conv.ID.Hex(),
// 			Name:        conv.Name,
// 			LastMessage: conv.LastMessage,
// 			CreatedAt:   timestamppb.New(conv.CreatedAt),
// 			UpdatedAt:   timestamppb.New(conv.UpdatedAt),
// 		}
// 		pbConversations = append(pbConversations, pbConv)
// 	}

// 	return &pb.GetConversationsResponse{
// 		Conversations: pbConversations,
// 	}, nil
// }
