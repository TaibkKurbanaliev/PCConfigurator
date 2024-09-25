package items

type Storage struct {
	Item
	Type             string `json:"type" bson:"type"`
	FormFactor       string `json:"formFactor" bson:"formFactor"`
	MemoryComponents string `json:"memoryComponents" bson:"memoryComponents"`
	Interface        string `json:"interface" bson:"interface"`
	Capacity         int    `json:"capacity" bson:"capacity"`
	MaxRead          int    `json:"maxRead" bson:"maxRead"`
	MaxWrite         int    `json:"maxWrite" bson:"maxWrite"`
	Width            int    `json:"width" bson:"width"`
	Height           int    `json:"height" bson:"height"`
	Length           int    `json:"length" bson:"length"`
}
