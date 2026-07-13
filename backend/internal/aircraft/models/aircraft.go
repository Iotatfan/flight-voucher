package models

type AircraftLayout struct {
	MinRow      int
	MaxRow      int
	SeatLetters []string
}

const (
	AircraftATR       = "ATR"
	AircraftAirbus320 = "Airbus 320"
	AircraftBoeing737 = "Boeing 737 Max"
)

var AircraftConfigs = map[string]AircraftLayout{
	AircraftATR: {
		MinRow:      1,
		MaxRow:      18,
		SeatLetters: []string{"A", "C", "D", "F"}, // No B or E
	},
	AircraftAirbus320: {
		MinRow:      1,
		MaxRow:      32,
		SeatLetters: []string{"A", "B", "C", "D", "E", "F"},
	},
	AircraftBoeing737: {
		MinRow:      1,
		MaxRow:      32,
		SeatLetters: []string{"A", "B", "C", "D", "E", "F"},
	},
}
