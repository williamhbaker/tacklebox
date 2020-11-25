package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wbaker85/tacklebox/pkg/models"
)

// HookModel comment
type HookModel struct {
	Col *mongo.Collection
	Ctx *context.Context
}

// InsertOne comment
func (m *HookModel) InsertOne(id *primitive.ObjectID, content string) error {
	doc := &models.HookDocument{
		ID:      id,
		Content: content,
	}
	m.Col.InsertOne(*m.Ctx, doc)
	return nil
}
