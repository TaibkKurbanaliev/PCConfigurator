package items

type CPUCooler struct {
	Item
	Type                   string `json:"type" bson:"type"`
	CPUSocketCompatibility string `json:"cpu_socket_compatibility" bson:"cpu_socket_compatibility"`
	FanSize                string `json:"fan_size" bson:"fan_size"`
	FanRPM                 string `json:"fan_rpm" bson:"fan_rpm"`
	FanAirFlow             string `json:"fan_air_flow" bson:"fan_air_flow"`
	FanNoise               string `json:"fan_noise" bson:"fan_noise"`
	FanConnector           string `json:"fan_connector" bson:"fan_connector"`
	LED                    string `json:"led" bson:"led"`
	MaxHeight              int    `json:"max_height" bson:"max_height"`
	Weight                 string `json:"weight" bson:"weight"`
}
