package service

import (
	"context"
	"snowApp/internal/model"
	"snowApp/internal/repository"
	"snowApp/pkg/logger"
	"time"

	pb "snowApp/gen"
)

type ChatService struct {
	messageRepo      repository.MessageRepository
	conversationRepo repository.ConversationRepository
	eventRepo        repository.EventRepository
	logger           logger.Logger
}

func NewChatService(
	mr repository.MessageRepository,
	cr repository.ConversationRepository,
	er repository.EventRepository,
	logger logger.Logger,
) *ChatService {
	return &ChatService{
		messageRepo:      mr,
		conversationRepo: cr,
		eventRepo:        er,
		logger:           logger,
	}
}

func (s *ChatService) CreateMessage(ctx context.Context, req *pb.SendMessageRequest) (*model.Message, error) {
	message := &model.Message{
		ConversationID: req.ConversationId,
		SenderID:       req.SenderId,
		Content:        req.Content,
		Type:           model.MessageType(req.Type),
		Metadata:       req.Metadata,
		SentAt:         time.Now(),
	}

	saveMsg, err := s.messageRepo.Save(ctx, message)
	if err != nil {
		return nil, err
	}

	if err := s.conversationRepo.UpdateLastMessage(ctx, req.ConversationId, saveMsg.ID); err != nil {
		return nil, err
	}
	return saveMsg, nil
}

func (s *ChatService) BroadcastMessage(ctx context.Context, msg *model.Message) error {
	event := &model.Event{
		Type:           model.EventTypeNewMessage,
		ConversationID: msg.ConversationID,
		MessageID:      msg.ID.String(),
		Data:           msg,
		Timestamp:      time.Now(),
	}
	return s.eventRepo.Publish(ctx, event)
}

func (s *ChatService) GetConversations(ctx context.Context, userID string) ([]*model.Conversation, error) {
	conversations, err := s.conversationRepo.FindByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}
func (s *ChatService) GetMessages(ctx context.Context, conversationID string, before time.Time, limit int) ([]*model.Message, error) {
	messages, err := s.messageRepo.FindByConversation(ctx, conversationID, before, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (s *ChatService) GetUserConversations(ctx context.Context, userID string, limit int, before time.Time) ([]*model.Conversation, error) {
	conversations, err := s.conversationRepo.GetUserConversations(ctx, userID, limit, before)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}
