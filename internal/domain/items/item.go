package items

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Item struct {
	ID                   bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name                 string        `json:"name" bson:"name"`
	Brand                string        `json:"brand" bson:"brand"`
	Model                string        `json:"model" bson:"model"`
	Series               string        `json:"series" bson:"series"`
	ReleaseYear          string        `json:"releaseYear" bson:"releaseYear"`
	ManufacturingCountry string        `json:"manufacturingCountry" bson:"manufacturingCountry"`
}
