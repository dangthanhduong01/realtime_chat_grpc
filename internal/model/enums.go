package model

type ConversationType int

const (
	ConversationTypePrivate ConversationType = iota
	ConversationTypeGroup
	ConversationTypeChannel
)

type MessageType int

const (
	MessageTypeText MessageType = iota
	MessageTypeImage
	MessageTypeVideo
	MessageTypeAudio
	MessageTypeFile
	MessageTypeLocation
	MessageTypeSticker
)

type EventType int

const (
	EventTypeNewMessage EventType = iota
	EventTypeMessageUpdated
	EventTypeMessageDeleted
	EventTypeConversationUpdated
	EventTypeUserStatusChanged
	EventTypeTypingIndicator
	EventTypeCallStarted
	EventTypeCallEnded
)

type StatusType int

const (
	StatusOffline StatusType = iota
	StatusOnline
	StatusAway
	StatusDoNotDisturb
)

type CallType int

const (
	CallTypeVoice CallType = iota
	CallTypeVideo
)

type CallStatus int

const (
	CallStatusRinging CallStatus = iota
	CallStatusOngoing
	CallStatusEnded
	CallStatusMissed
	CallStatusDeclined
)
