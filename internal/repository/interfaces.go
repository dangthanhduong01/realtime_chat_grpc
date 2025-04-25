package repository

import (
	"context"
	"snowApp/internal/model"
	"time"
)

type MessageRepository interface {
	Save(ctx context.Context, message *model.Message) (*model.Message, error)
	FindByID(ctx context.Context, id string) (*model.Message, error)
	FindByConversation(ctx context.Context, conversationID string, before time.Time, limit int) ([]*model.Message, error)
	MarkAsSeen(ctx context.Context, messageID string) error
	Delete(ctx context.Context, messageID string) error
}
