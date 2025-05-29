package redis

import (
	"context"
	"encoding/json"
	"snowApp/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type EventRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewEventRepository(client *redis.Client) *EventRepository {
	return &EventRepository{
		client: client,
		ttl:    24 * time.Hour, // Set a default TTL of 24 hours
	}
}

func (r *EventRepository) Publish(ctx context.Context, event *model.Event) error {
	event.Timestamp = time.Now()

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	convChannel := "conversation:" + event.ConversationID
	if err := r.client.Publish(ctx, convChannel, string(data)).Err(); err != nil {
		return err
	}

	// Also publish to user channels for multi-device sync
	userChannel := "user:" + event.UserID
	return r.client.Publish(ctx, userChannel, data).Err()

	// Luu vao stram de dam bao durability
	// _, err = r.client.Client.XAdd(ctx, &redis.XAddArgs{
	// 	Stream: "events_stream",
	// 	Values: map[string]interface{}{
	// 		"type":    event.Type,
	// 		"payload": data,
	// 	},
	// }).Result()

	// return err
}

func (r *EventRepository) Subscribe(ctx context.Context, userID string) (<-chan *model.Event, error) {
	pubsub := r.client.Subscribe(ctx, "user:"+userID)
	eventChan := make(chan *model.Event)

	go func() {
		defer pubsub.Close()
		ch := pubsub.Channel()

		for {
			select {
			case msg := <-ch:
				var event model.Event
				if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
					continue
				}
				eventChan <- &event
			case <-ctx.Done():
				close(eventChan)
				return
			}
		}
	}()

	return eventChan, nil
}

func (r *EventRepository) Unsubscribe(ctx context.Context, userID string) error {
	pubsub := r.client.Subscribe(ctx, "user:"+userID)
	if err := pubsub.Close(); err != nil {
		return err
	}
	return nil
}

// func (r *EventRepository) Subscribe(ctx context.Context, channel string, eventChan chan<- *model.Event) error {
// 	pubsub := r.client.Subscribe(ctx, channel)

// 	// Goroutine để nhận messages
// 	go func() {
// 		defer pubsub.Close()

// 		ch := pubsub.Channel()
// 		for {
// 			select {
// 			case msg := <-ch:
// 				var event model.Event
// 				if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
// 					continue
// 				}
// 				eventChan <- &event
// 			case <-ctx.Done():
// 				return
// 			}
// 		}
// 	}()

// 	return nil
// }
