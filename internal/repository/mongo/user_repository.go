package mongo

import (
	"context"
	"fmt"
	"snowApp/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName string) *UserRepository {
	return &UserRepository{
		collection: client.Database(dbName).Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid.Hex()
	} else {
		return nil, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}
	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	oid, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Username,
			"email": user.Email,
			// Add other fields as necessary
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error retrieving user by email: %w", err)
	}
	return &user, nil
}
