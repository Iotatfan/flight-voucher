package models

type CheckFlightResponse struct {
	Exists bool `json:"exists"`
}

type GenerateRandomSeatsResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}
