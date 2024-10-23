package items

import "go.mongodb.org/mongo-driver/v2/bson"

type GPU struct {
	ID bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Item
	Interface           string `json:"interface" bson:"interface"`
	ChipsetManufacturer string `json:"chipset_manufacturer" bson:"chipset_manufacturer"`
	GPUSeries           string `json:"gpu_series" bson:"gpu_series"`
	MemoryType          string `json:"memory_type" bson:"memory_type"`
	Cooler              string `json:"cooler" bson:"cooler"`
	FormFactor          string `json:"form_factor" bson:"form_factor"`
	Length              int    `json:"length" bson:"length"`
	Height              int    `json:"height" bson:"height"`
	Width               int    `json:"width" bson:"width"`
	TDP                 int    `json:"tdp" bson:"tdp"`
	RecommendedPSU      int    `json:"recommended_psu" bson:"recommended_psu"`
	MemorySize          int    `json:"memory_size" bson:"memory_size"`
	MemoryInterface     int    `json:"memory_interface" bson:"memory_interface"`
	NumberOfHDMI        int    `json:"number_of_hdmi" bson:"number_of_hdmi"`
}
