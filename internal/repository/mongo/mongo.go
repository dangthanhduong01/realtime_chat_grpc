package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func New(uri, dbName string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(nil, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	db := client.Database(dbName)

	return &MongoClient{
		Client: client,
		DB:     db,
	}, nil
}

func (m *MongoClient) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return m.Client.Disconnect(ctx)
}
func (m *MongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return m.Client.Ping(ctx, rp)
}
func (m *MongoClient) Database(dbName string) *mongo.Database {
	return m.Client.Database(dbName)
}
func (m *MongoClient) Collection(collectionName string) *mongo.Collection {
	return m.Client.Database("").Collection(collectionName)
}
func (m *MongoClient) CollectionWithDB(dbName, collectionName string) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName)
}
func (m *MongoClient) CollectionWithOptions(dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContext(ctx context.Context, dbName, collectionName string) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName)
}
func (m *MongoClient) CollectionWithContextAndOptions(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClient(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabase(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabaseAndCollection(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabaseAndCollectionAndClient(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabaseAndCollectionAndClientAndDatabase(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabaseAndCollectionAndClientAndDatabaseAndCollection(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabaseAndCollectionAndClientAndDatabaseAndCollectionAndClient(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabaseAndCollectionAndClientAndDatabaseAndCollectionAndClientAndDatabase(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
func (m *MongoClient) CollectionWithContextAndOptionsAndClientAndDatabaseAndCollectionAndClientAndDatabaseAndCollectionAndClientAndDatabaseAndCollection(ctx context.Context, dbName, collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return m.Client.Database(dbName).Collection(collectionName, opts...)
}
