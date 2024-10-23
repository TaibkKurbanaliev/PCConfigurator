package items

import "go.mongodb.org/mongo-driver/v2/bson"

type Memory struct {
	ID bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Item
	Type      string  `json:"type" bson:"type"`
	Timing    string  `json:"timing" bson:"timing"`
	Capacity  int     `json:"capacity" bson:"capacity"`
	Frequency int     `json:"frequency" bson:"frequency"`
	CAS       int     `json:"cas" bson:"cas"`
	Count     int     `json:"count" bson:"count"`
	Voltage   float32 `json:"voltage" bson:"voltage"`
}
