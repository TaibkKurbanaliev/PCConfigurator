package items

import "go.mongodb.org/mongo-driver/v2/bson"

type Item struct {
	ID                   bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name                 string        `json:"" bson:""`
	Brand                string        `json:"" bson:""`
	Model                string        `json:"" bson:""`
	Series               string        `json:"" bson:""`
	ReleaseYear          string        `json:"" bson:""`
	ManufacturingCountry string        `json:"" bson:""`
}
