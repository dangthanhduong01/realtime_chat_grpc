package repository

import (
	"context"
	"snowApp/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository interface {
	Save(ctx context.Context, message *model.Message) (*model.Message, error)
	// FindByID(ctx context.Context, id string) (*model.Message, error)
	FindByConversation(ctx context.Context, conversationID string, before time.Time, limit int) ([]*model.Message, error)
	// MarkAsSeen(ctx context.Context, messageID string) error
	// Delete(ctx context.Context, messageID string) error
}

type ConversationRepository interface {
	Create(ctx context.Context, conv *model.Conversation) (*model.Conversation, error)
	// FindByID(ctx context.Context, id string) (*model.Conversation, error)
	FindPrivateConversation(ctx context.Context, userID string, otherUserID string) (*model.Conversation, error)
	// FindUserConversations(ctx context.Context, userID string, limit int, before time.Time) ([]*model.Conversation, error)
	AddParticipants(ctx context.Context, conversationID string, userIDs []string) error
	UpdateLastMessage(ctx context.Context, conversationID string, messageID primitive.ObjectID) error
	FindByUser(ctx context.Context, userID string) ([]*model.Conversation, error)
	GetUserConversations(ctx context.Context, userID string, limit int, before time.Time) ([]*model.Conversation, error)
}

type EventRepository interface {
	Publish(ctx context.Context, event *model.Event) error
	Subscribe(ctx context.Context, userID string) (<-chan *model.Event, error)
	Unsubscribe(ctx context.Context, channel string) error
}

// StatusRepository xử lý trạng thái người dùng
type StatusRepository interface {
	Update(ctx context.Context, status *model.UserStatus, expiresAt time.Time) error
	Get(ctx context.Context, userID string) (*model.UserStatus, error)
	BatchGet(ctx context.Context, userIDs []string) (map[string]*model.UserStatus, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
