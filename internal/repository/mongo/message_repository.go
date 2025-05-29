package mongo

import (
	"context"
	"snowApp/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageRepository struct {
	collection *mongo.Collection
}

func NewMessageRepository(client *mongo.Client, dbName string) *MessageRepository {
	return &MessageRepository{
		collection: client.Database(dbName).Collection("messages"),
	}
}

func (r *MessageRepository) Save(ctx context.Context, message *model.Message) (*model.Message, error) {
	if message.ID.IsZero() {
		message.SentAt = time.Now()
		res, err := r.collection.InsertOne(ctx, message)
		if err != nil {
			return nil, err
		}
		message.ID = res.InsertedID.(primitive.ObjectID)
		return message, nil
	}
	message.UpdatedAt = time.Now()
	update := bson.M{
		"$set": message,
	}
	_, err := r.collection.UpdateByID(ctx, message.ID, update)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessageRepository) FindByConversation(ctx context.Context, conversationID string, before time.Time, limit int) ([]*model.Message, error) {
	filter := bson.M{
		"conversation_id": conversationID,
	}

	if !before.IsZero() {
		filter["sent_at"] = bson.M{"$lt": before}
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "sent_at", Value: -1}}).
		SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var messages []*model.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// func (r *MessageRepository) MarkAsSeen(ctx context.Context, messageIDs []string, userID string) error {
// 	ids := make([]primitive.ObjectID, len(messageIDs))
// 	for _, id := range messageIDs {
// 		old, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		ids = append(ids, old)
// 	}
// 	filter := bson.M{
// 		"_id": bson.M{"$in": ids},
// 	}

// 	update := bson.M{
// 		"$addToSet": bson.M{
// 			"seen_by": userID,
// 		},
// 		"$set": bson.M{
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	_, err := r.collection.UpdateMany(ctx, filter, update)
// 	return err
// }
