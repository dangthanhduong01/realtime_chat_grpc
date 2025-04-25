package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMessage(convID, senderID, content string, msgType MessageType) *Message {
	return &Message{
		ConversationID: convID,
		SenderID:       senderID,
		Content:        content,
		Type:           msgType,
		SentAt:         time.Now(),
		SeenBy:         []string{},
	}
}

func (m *Message) MarkAsSeen(userID string) {
	for _, id := range m.SeenBy {
		if id == userID {
			return
		}
	}

	now := time.Now()
	m.SeenBy = append(m.SeenBy, userID)
	if m.DeliveredAt == nil {
		m.DeliveredAt = &now
	}
	m.ReadAt = &now
}

func (c *Conversation) IsGroup() bool {
	return c.Type == ConversationTypeGroup
}

func (c *Conversation) AddParticipant(userID string) {
	for _, id := range c.Participants {
		if id == userID {
			return
		}
	}

	c.Participants = append(c.Participants, userID)
	c.UpdatedAt = time.Now()
}

func (e *Event) Channel() string {
	switch e.Type {
	case EventTypeNewMessage, EventTypeMessageUpdated, EventTypeMessageDeleted:
		return "conversation:" + e.ConversationID
	case EventTypeUserStatusChanged, EventTypeTypingIndicator:
		return "user:" + e.UserID
	default:
		return "events"
	}
}

func GenerateID() string {
	return primitive.NewObjectID().Hex()
}
