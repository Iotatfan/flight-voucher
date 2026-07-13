package models

type CheckFlightResponse struct {
	Exists bool `json:"exists"`
}

type GenerataRandomSeatsResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}
