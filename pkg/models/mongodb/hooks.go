package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wbaker85/tacklebox/pkg/models"
)

// HookModel is a struct representing a hook document in the docdb
type HookModel struct {
	Col *mongo.Collection
	Ctx *context.Context
}

// InsertOne is for adding one hook document to the database
func (m *HookModel) InsertOne(id *primitive.ObjectID, content string) error {
	doc := &models.HookDocument{
		ID:      id,
		Content: content,
	}
	_, err := m.Col.InsertOne(*m.Ctx, doc)
	if err != nil {
		return err
	}

	return nil
}
