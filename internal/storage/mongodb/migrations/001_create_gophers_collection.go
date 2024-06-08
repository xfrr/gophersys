package mongomigrations

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateGophersCollection struct {
	db *mongo.Database
}

func NewCreateGophersCollection(db *mongo.Database) *CreateGophersCollection {
	return &CreateGophersCollection{db: db}
}

func (m *CreateGophersCollection) collname() string {
	return "gophers"
}

func (m *CreateGophersCollection) Name() string {
	return "001_create_gophers_collection"
}

func (m *CreateGophersCollection) Up(ctx context.Context) error {
	err := m.db.CreateCollection(ctx, m.collname(), m.getCreateCollectionOptions())
	if err != nil {
		return err
	}
	return m.createIndexes(ctx, m.db.Collection(m.collname()))
}

func (m *CreateGophersCollection) Down(ctx context.Context) error {
	return m.db.Collection("gophers").Drop(ctx)
}

func (m *CreateGophersCollection) getCreateCollectionOptions() *options.CreateCollectionOptions {
	return &options.CreateCollectionOptions{
		Validator: m.getValidator(),
	}
}

func (m *CreateGophersCollection) createIndexes(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{
				"username": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	})
	return err
}

func (m *CreateGophersCollection) getValidator() bson.M {
	return bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"username", "status", "created_at", "updated_at"},
			"properties": bson.M{
				"username": bson.M{
					"bsonType":    "string",
					"description": "username of the gopher",
				},
				"name": bson.M{
					"bsonType":    "string",
					"description": "name of the gopher",
				},
				"status": bson.M{
					"bsonType":    "string",
					"description": "status of the gopher",
					"enum":        []string{"active", "inactive", "suspended", "deleted"},
				},
				"metadata": bson.M{
					"bsonType":    "object",
					"description": "metadata of the gopher",
				},
				"created_at": bson.M{
					"bsonType":    "date",
					"description": "creation date of the gopher",
				},
				"updated_at": bson.M{
					"bsonType":    "date",
					"description": "update date of the gopher",
				},
			},
		},
	}
}
