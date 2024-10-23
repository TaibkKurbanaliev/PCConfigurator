package items

import "go.mongodb.org/mongo-driver/v2/bson"

type CPU struct {
	ID bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Item
	NumberOfCores     int      `json:"number_of_cores" bson:"number_of_cores"`
	NumberOfThreads   int      `json:"number_of_threads" bson:"number_of_threads"`
	TDP               int      `json:"tdp" bson:"tdp"`
	L1CacheSize       int      `json:"l1_cache_size" bson:"l1_cache_size"`
	L2CacheSize       int      `json:"l2_cache_size" bson:"l2_cache_size"`
	L3CacheSize       int      `json:"l3_cache_size" bson:"l3_cache_size"`
	TechnologyProcess int      `json:"technology_process" bson:"technology_process"`
	BaseFrequency     float32  `json:"base_frequency" bson:"base_frequency"`
	MaxFrequency      float32  `json:"max_frequency" bson:"max_frequency"`
	MemoryType        string   `json:"memory_type" bson:"memory_type"`
	Chipsets          []string `json:"chipsets" bson:"chipsets"`
}
