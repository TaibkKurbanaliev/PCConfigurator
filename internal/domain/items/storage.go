package items

type Storage struct {
	Item
	Type             string
	FormFactor       string
	MemoryComponents string
	Interface        string
	Capacity         int
	MaxRead          int
	MaxWrite         int
	Width            int
	Height           int
	Length           int
}
