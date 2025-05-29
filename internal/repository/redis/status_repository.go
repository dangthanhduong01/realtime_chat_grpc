package redis

import (
	"context"
	"encoding/json"
	"snowApp/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type StatusRepository struct {
	client *RedisClient
}

func NewStatusRepository(client *RedisClient) *StatusRepository {
	return &StatusRepository{
		client: client,
	}
}

func (r *StatusRepository) Update(ctx context.Context, status *model.UserStatus, expiresAt time.Time) error {
	status.LastUpdated = time.Now()

	data, err := json.Marshal(status)
	if err != nil {
		return err
	}

	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		ttl = 24 * time.Hour // Default TTL
	}

	return r.client.Client.Set(ctx, r.key(status.UserID), data, ttl).Err()
}

func (r *StatusRepository) Get(ctx context.Context, userID string) (*model.UserStatus, error) {
	data, err := r.client.Client.Get(ctx, r.key(userID)).Result()
	if err != nil {
		if err == redis.Nil {
			return &model.UserStatus{
				UserID:      userID,
				Status:      model.StatusOffline,
				LastUpdated: time.Now(),
			}, nil
		}
		return nil, err
	}

	var status model.UserStatus
	if err := json.Unmarshal([]byte(data), &status); err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *StatusRepository) BatchGet(ctx context.Context, userIDs []string) (map[string]*model.UserStatus, error) {
	keys := make([]string, len(userIDs))
	for i, id := range userIDs {
		keys[i] = r.key(id)
	}

	values, err := r.client.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	result := make(map[string]*model.UserStatus)
	for i, val := range values {
		if val == nil {
			result[userIDs[i]] = &model.UserStatus{
				UserID:      userIDs[i],
				Status:      model.StatusOffline,
				LastUpdated: time.Now(),
			}
			continue
		}

		var status model.UserStatus
		if err := json.Unmarshal([]byte(val.(string)), &status); err != nil {
			return nil, err
		}
		result[userIDs[i]] = &status
	}

	return result, nil
}

func (r *StatusRepository) key(userID string) string {
	return "user_status:" + userID
}
