package model

import "errors"

var (
	ErrConversationNotFound  = errors.New("conversation not found")
	ErrMessageNotFound       = errors.New("message not found")
	ErrUserNotInConversation = errors.New("user not in conversation")
	ErrPermissionDenied      = errors.New("permission denied")
	ErrInvalidMessageType    = errors.New("invalid message type")
	ErrCallAlreadyEnded      = errors.New("call already ended")
	ErrDeviceNotRegistered   = errors.New("device not registered")
)
