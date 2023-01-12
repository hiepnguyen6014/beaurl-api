package models

import (
	"context"

	"beaurl.vn/api/src/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Shorten struct {
	ID     string `json:"_id,omitempty" bson:"_id,omitempty"`
	Path   string `json:"path" bson:"path, required"`
	Target string `json:"target" bson:"target, required"`
}

func Save(path string, target string) error {

	mongoose := db.GetInstance().Collection("shorten")

	_, err := mongoose.InsertOne(context.Background(), Shorten{
		Path:   path,
		Target: target,
	})

	return err
}

func FindByPath(path string) (Shorten, error) {

	mongoose := db.GetInstance().Collection("shorten")

	var shorten Shorten

	err := mongoose.FindOne(context.Background(), bson.D{{Key: "path", Value: path}}).Decode(&shorten)

	return shorten, err
}
