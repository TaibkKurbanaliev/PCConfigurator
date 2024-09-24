package items

type Motherboard struct {
	Item
	SocketType          string
	CPUSeries           string
	Chipset             string
	AudioChipset        string
	LANChipset          string
	MemorySuports       []string
	NumberOfMemorySlots int
	MaxMemory           int
	MaxLanSpeed         int
	NumberOfPCIE5X16    int
	NumberOfPCIE4X16    int
	NumberOfPCIE3X16    int
	NumberOfPCIE4X8     int
	NumberOfPCIE3X8     int
	NumberOfSATA        int
	BluetoothVersion    float32
}
