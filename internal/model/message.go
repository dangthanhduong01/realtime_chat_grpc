package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID            primitive.ObjectID `bson: "_id, omitempty" json:"id"`
	Type          ConversationType   `bson: "type, omitempty" json:"type"`
	Title         string             `bson: "title, omitempty" json:"title"`
	Participants  []string           `bson: "participants, omitempty" json:"participants"`
	Admins        []string           `bson: "admins, omitempty" json:"admins"`
	AvatarURL     string             `bson: "avatar_url, omitempty" json:"avatar_url"`
	LastMessageID primitive.ObjectID `bson: "last_message_id, omitempty" json:"last_message_id"`
	CreatedAt     time.Time          `bson: "created_at, omitempty" json:"created_at"`
	UpdatedAt     time.Time          `bson: "updated_at, omitempty" json:"updated_at"`
}

type Message struct {
	ID             primitive.ObjectID `bson: "_id, omitempty" json:"id"`
	ConversationID string             `bson: "conversation_id, omitempty" json:"conversation_id"`
	SenderID       string             `bson: "sender_id, omitempty" json:"sender_id"`
	Type           MessageType        `bson: "type, omitempty" json:"type"`
	Content        string             `bson: "content, omitempty" json:"content"`
	Metadata       map[string]string  `bson: "metadata, omitempty" json:"metadata"`
	SeenBy         []string           `bson: "seen_by, omitempty" json:"seen_by"`
	SentAt         time.Time          `bson:"sent_at" json:"sent_at"`
	DeliveredAt    *time.Time         `bson: "delivered_at, omitempty" json:"delivered_at"`
	ReadAt         *time.Time         `bson: "read_at, omitempty" json:"read_at"`
	UpdatedAt      time.Time          `bson: "created_at, omitempty" json:"created_at"`
}

type Event struct {
	ID             string      `bson: "_id, omitempty" json:"id"`
	Type           EventType   `bson: "type, omitempty" json:"type"`
	ConversationID string      `bson: "conversation_id, omitempty" json:"conversation_id"`
	UserID         string      `bson: "user_id, omitempty" json:"user_id"`
	MessageID      string      `bson: "message_id, omitempty" json:"message_id"`
	Data           interface{} `bson: "data, omitempty" json:"data"`
	Timestamp      time.Time   `bson: "timestamp, omitempty" json:"timestamp"`
}

type UserStatus struct {
	UserID      string      `bson: "user_id, omitempty" json:"user_id"`
	Status      StatusType  `bson: "status, omitempty" json:"status"`
	StatusText  string      `bson: "status_text, omitempty" json:"status_text"`
	LastSeen    time.Time   `bson: "last_seen, omitempty" json:"last_seen"`
	LastUpdated time.Time   `bson: "updated_at, omitempty" json:"updated_at"`
	DeviceInfo  *DeviceInfo `bson: "device_info, omitempty" json:"device_info"`
}

type DeviceInfo struct {
	ID       string `bson: "id, omitempty" json:"id"`
	Platform string `bson: "platform, omitempty" json:"platform"`
	Version  string `bson: "version, omitempty" json:"version"`
}

type CallInfo struct {
	ID             string            `bson:"_id,omitempty" json:"id"`
	ConversationID string            `bson:"conversation_id" json:"conversation_id"`
	InitiatorID    string            `bson:"initiator_id" json:"initiator_id"`
	Participants   []CallParticipant `bson:"participants" json:"participants"`
	Type           CallType          `bson:"type" json:"type"`
	Status         CallStatus        `bson:"status" json:"status"`
	StartedAt      time.Time         `bson:"started_at" json:"started_at"`
	EndedAt        *time.Time        `bson:"ended_at,omitempty" json:"ended_at,omitempty"`
}

// CallParticipant thông tin người tham gia cuộc gọi
type CallParticipant struct {
	UserID     string      `bson:"user_id" json:"user_id"`
	JoinedAt   time.Time   `bson:"joined_at" json:"joined_at"`
	LeftAt     *time.Time  `bson:"left_at,omitempty" json:"left_at,omitempty"`
	StreamInfo *StreamInfo `bson:"stream_info,omitempty" json:"stream_info,omitempty"`
}

// StreamInfo thông tin media stream
type StreamInfo struct {
	AudioEnabled  bool `bson:"audio_enabled" json:"audio_enabled"`
	VideoEnabled  bool `bson:"video_enabled" json:"video_enabled"`
	ScreenSharing bool `bson:"screen_sharing" json:"screen_sharing"`
}
