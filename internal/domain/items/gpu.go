package items

type GPU struct {
	Item
	Interface           string
	ChipsetManufacturer string
	GPUSeries           string
	MemoryType          string
	Cooler              string
	FormFactor          string
	Length              int
	Height              int
	Width               int
	TDP                 int
	RecommendedPSU      int
	MemorySize          int
	MemoryInterface     int
	NubmerOfHDMI        int
}
