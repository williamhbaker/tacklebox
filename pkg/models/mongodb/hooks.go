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

// GetMany hooks from a given slice of document IDs
func (m *HookModel) GetMany(docIDs []string) ([]*models.HookDocument, error) {
	oids := make([]primitive.ObjectID, len(docIDs))
	for idx := range docIDs {
		objID, err := primitive.ObjectIDFromHex(docIDs[idx])
		if err == nil {
			oids[idx] = objID
		}
	}

	query := bson.M{"_id": bson.M{"$in": oids}}

	cursor, err := m.Col.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	hooks := []*models.HookDocument{}
	for cursor.Next(context.TODO()) {
		d := &models.HookDocument{}
		err = cursor.Decode(d)
		if err != nil {
			return nil, err
		}

		hooks = append(hooks, d)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}

	return hooks, nil
}
