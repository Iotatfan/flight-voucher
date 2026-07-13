package models

type Voucher struct {
	ID           int64  `gorm:"primaryKey;autoIncrement;column:id"`
	CrewName     string `gorm:"type:text;not null;column:crew_name"`
	CrewID       string `gorm:"type:text;not null;column:crew_id"`
	FlightNumber string `gorm:"type:text;uniqueIndex:idx_flight_date;not null;column:flight_number"`
	FlightDate   string `gorm:"type:text;uniqueIndex:idx_flight_date;not null;column:flight_date"`
	AircraftType string `gorm:"type:text;not null;column:aircraft_type"`
	Seat1        string `gorm:"type:text;not null;column:seat1"`
	Seat2        string `gorm:"type:text;not null;column:seat2"`
	Seat3        string `gorm:"type:text;not null;column:seat3"`
	Created_at   string `gorm:"type:timestamp;autoCreateTime;not null;column:created_at"`
}

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
