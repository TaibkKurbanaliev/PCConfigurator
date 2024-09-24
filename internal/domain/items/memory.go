package items

type Memory struct {
	Item
	Type      string
	Timing    string
	Capacity  int
	Frequency int
	CAS       int
	Count     int
	Voltage   float32
}
