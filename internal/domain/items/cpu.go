package items

type CPU struct {
	Item
	NumberOfCores     int
	NumberOfThreads   int
	TDP               int
	L1CacheSize       int
	L2CacheSize       int
	L3CacheSize       int
	TechnologyProcess int
	BaseFrequency     float32
	MaxFrequency      float32
	MemoryType        string
	Chipsets          []string
}
