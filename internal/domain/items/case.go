package items

import "go.mongodb.org/mongo-driver/v2/bson"

type Case struct {
	ID bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Item
	Type                     string `json:"type" bson:"type"`
	Color                    string `json:"color" bson:"color"`
	CaseMaterial             string `json:"case_material" bson:"case_material"`
	PowerSupplyMounted       string `json:"power_supply_mounted" bson:"power_supply_mounted"`
	MotherboardCompatibility string `json:"motherboard_compatibility" bson:"motherboard_compatibility"`
	FrontPorts               string `json:"front_ports" bson:"front_ports"`
	MaxGPULength             int    `json:"max_gpu_length" bson:"max_gpu_length"`
	MaxCPUCoolerHeight       int    `json:"max_cpu_cooler_height" bson:"max_cpu_cooler_height"`
	MaxPSULength             int    `json:"max_psu_length" bson:"max_psu_length"`
	Weight                   int    `json:"weight" bson:"weight"`
	WithPowerSupply          bool   `json:"with_power_supply" bson:"with_power_supply"`
	SidePanelWindow          bool   `json:"side_panel_window" bson:"side_panel_window"`
	DustFilters              bool   `json:"dust_filters" bson:"dust_filters"`
}
