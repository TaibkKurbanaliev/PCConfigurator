package items

type Memory struct {
	Item
	Type      string  `json:"type" bson:"type"`
	Timing    string  `json:"timing" bson:"timing"`
	Capacity  int     `json:"capacity" bson:"capacity"`
	Frequency int     `json:"frequency" bson:"frequency"`
	CAS       int     `json:"cas" bson:"cas"`
	Count     int     `json:"count" bson:"count"`
	Voltage   float32 `json:"voltage" bson:"voltage"`
}
