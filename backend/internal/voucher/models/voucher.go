package models

import "errors"

var ErrVoucherAlreadyExists = errors.New("vouchers already generated for this flight on this date")
var ErrInvalidAircraftType = errors.New("invalid aircraft type")
var ErrFewerThan3SeatsAvailable = errors.New("aircraft layout has fewer than 3 seats")

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
