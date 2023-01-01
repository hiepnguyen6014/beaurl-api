package models

import (
	"context"

	"beaurl.vn/api/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Shorten struct {
	ID     string `json:"_id,omitempty" bson:"_id,omitempty"`
	Path   string `json:"path" bson:"path, required"`
	Target string `json:"target" bson:"target, required"`
}

var mongoose *mongo.Collection = db.GetInstance().Collection("shorten")

func Save(path string, target string) error {

	_, err := mongoose.InsertOne(context.Background(), Shorten{
		Path:   path,
		Target: target,
	})

	return err
}

func FindByPath(path string) (Shorten, error) {

	var shorten Shorten

	err := mongoose.FindOne(context.Background(), bson.D{{"path", path}}).Decode(&shorten)

	return shorten, err
}
