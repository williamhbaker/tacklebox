package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wbaker85/tacklebox/pkg/models"
)

// HookModel is a struct representing a hook document in the docdb
type HookModel struct {
	Col *mongo.Collection
	Ctx *context.Context
}

// Insert is for adding one hook document to the database
func (m *HookModel) Insert(content string, id *primitive.ObjectID) (string, error) {
	doc := &models.HookDocument{
		ID:      id,
		Content: content,
	}
	res, err := m.Col.InsertOne(*m.Ctx, doc)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// Destroy is for deleting a document from the database by ID (string)
func (m *HookModel) Destroy(id string) error {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.Col.DeleteOne(*m.Ctx, bson.M{"_id": pid})
	if err != nil {
		return err
	}

	return nil
}
