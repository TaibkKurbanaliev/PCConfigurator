package items

import "go.mongodb.org/mongo-driver/v2/bson"

type Motherboard struct {
	ID bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Item
	SocketType          string   `json:"socket_type" bson:"socket_type"`
	CPUSeries           string   `json:"cpu_series" bson:"cpu_series"`
	Chipset             string   `json:"chipset" bson:"chipset"`
	AudioChipset        string   `json:"audio_chipset" bson:"audio_chipset"`
	LANChipset          string   `json:"lan_chipset" bson:"lan_chipset"`
	MemorySupports      []string `json:"memory_supports" bson:"memory_supports"`
	NumberOfMemorySlots int      `json:"number_of_memory_slots" bson:"number_of_memory_slots"`
	MaxMemory           int      `json:"max_memory" bson:"max_memory"`
	MaxLanSpeed         int      `json:"max_lan_speed" bson:"max_lan_speed"`
	NumberOfPCIE5X16    int      `json:"number_of_pc_ie5_x16" bson:"number_of_pc_ie5_x16"`
	NumberOfPCIE4X16    int      `json:"number_of_pc_ie4_x16" bson:"number_of_pc_ie4_x16"`
	NumberOfPCIE3X16    int      `json:"number_of_pc_ie3_x16" bson:"number_of_pc_ie3_x16"`
	NumberOfPCIE4X8     int      `json:"number_of_pc_ie4_x8" bson:"number_of_pc_ie4_x8"`
	NumberOfPCIE3X8     int      `json:"number_of_pc_ie3_x8" bson:"number_of_pc_ie3_x8"`
	NumberOfSATA        int      `json:"number_of_sata" bson:"number_of_sata"`
	BluetoothVersion    float32  `json:"bluetooth_version" bson:"bluetooth_version"`
}
