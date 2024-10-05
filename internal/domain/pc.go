package domain

import "go.mongodb.org/mongo-driver/v2/bson"

// This is needed to write a simple form of PC components
type SimpleItem struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Brand string `json:"brand" bson:"brand"`
}

type PC struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Rating      float32       `json:"rating" bson:"rating"`
	CPU         SimpleItem    `json:"cpu" bson:"cpu"`
	GPU         SimpleItem    `json:"gpu" bson:"gpu"`
	Motherboard SimpleItem    `json:"motherboard" bson:"motherboard"`
	Memory      SimpleItem    `json:"memory" bson:"memory"`
	Storage     SimpleItem    `json:"storage" bson:"storage"`
	PowerSupply SimpleItem    `json:"power_supply" bson:"power_supply"`
	CPUCooler   SimpleItem    `json:"cpu_cooler" bson:"cpu_cooler"`
	Case        SimpleItem    `json:"case" bson:"case"`
	OwnerName   string        `json:"owner_name" bson:"owner_name"`
	OwnerID     string        `json:"owner_id" bson:"owner_id"`
}
