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

// GetMany hooks from a given slice of document IDs
func (m *HookModel) GetMany(docIDs []string) ([]*models.HookDocument, error) {
	oids := docIDsToOIDS(docIDs)

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

// DestroyMany deletes many hooks, given a slice of document IDs
func (m *HookModel) DestroyMany(docIDs []string) (int, error) {
	oids := docIDsToOIDS(docIDs)

	query := bson.M{"_id": bson.M{"$in": oids}}

	res, err := m.Col.DeleteMany(context.TODO(), query)
	if err != nil {
		return 0, err
	}

	return int(res.DeletedCount), nil
}

func docIDsToOIDS(docIDs []string) []primitive.ObjectID {
	oids := make([]primitive.ObjectID, len(docIDs))
	for idx := range docIDs {
		objID, err := primitive.ObjectIDFromHex(docIDs[idx])
		if err == nil {
			oids[idx] = objID
		}
	}
	return oids
}
