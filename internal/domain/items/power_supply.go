package items

import "go.mongodb.org/mongo-driver/v2/bson"

type PowerSupply struct {
	ID bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Item
	Color               string `json:"color" bson:"color"`
	Type                string `json:"type" bson:"type"`
	MaximumPower        int    `json:"maximum_power" bson:"maximum_power"`
	MainConnector       string `json:"main_connector" bson:"main_connector"`
	PCIExpressConnector string `json:"pci_express_connector" bson:"pci_express_connector"`
	SATAPowerConnector  int    `json:"sata_power_connector" bson:"sata_power_connector"`
	MaxPSULength        int    `json:"max_psu_length" bson:"max_psu_length"`
	Modular             string `json:"modular" bson:"modular"`
	EnergyEfficient     string `json:"energy_efficient" bson:"energy_efficient"`
	Connectors          string `json:"connectors" bson:"connectors"`
	CableSpec           string `json:"cable_spec" bson:"cable_spec"`
	Weight              string `json:"weight" bson:"weight"`
}
