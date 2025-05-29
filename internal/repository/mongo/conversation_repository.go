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

type ConversationRepository struct {
	collection *mongo.Collection
}

func NewConversationRepository(client *mongo.Client, dbName string) *ConversationRepository {
	return &ConversationRepository{
		collection: client.Database(dbName).Collection("conversations"),
	}
}

func (r *ConversationRepository) Create(ctx context.Context, conv *model.Conversation) (*model.Conversation, error) {
	conv.CreatedAt = time.Now()
	conv.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, conv)
	if err != nil {
		return nil, err
	}

	conv.ID = res.InsertedID.(primitive.ObjectID)
	return conv, nil
}

func (r *ConversationRepository) FindPrivateConversation(ctx context.Context, user1ID, user2ID string) (*model.Conversation, error) {
	filter := bson.M{
		"type": model.ConversationTypePrivate,
		"participants": bson.M{
			"$all": []string{user1ID, user2ID},
		},
	}

	var conv model.Conversation
	err := r.collection.FindOne(ctx, filter).Decode(&conv)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &conv, nil
}

func (r *ConversationRepository) AddParticipants(ctx context.Context, conversationID string, userIDs []string) error {
	id, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$addToSet": bson.M{
			"participants": bson.M{
				"$each": userIDs,
			},
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *ConversationRepository) UpdateLastMessage(ctx context.Context, conversationID string, messageID primitive.ObjectID) error {
	id, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"last_message_id": messageID,
			"updated_at":      time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *ConversationRepository) FindByUser(ctx context.Context, userID string) ([]*model.Conversation, error) {
	filter := bson.M{
		"participants": userID,
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var conversations []*model.Conversation
	if err := cursor.All(ctx, &conversations); err != nil {
		return nil, err
	}

	return conversations, nil
}

func (r *ConversationRepository) GetUserConversations(ctx context.Context, userID string, limit int, before time.Time) ([]*model.Conversation, error) {
	filter := bson.M{
		"participants": userID,
	}

	if !before.IsZero() {
		filter["updated_at"] = bson.M{"$lt": before}
	}

	opts := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}}).SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var conversations []*model.Conversation
	if err := cursor.All(ctx, &conversations); err != nil {
		return nil, err
	}

	return conversations, nil
}
