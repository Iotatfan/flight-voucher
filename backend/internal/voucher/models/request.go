package models

type CheckFlightRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

type GenerateRandomSeatsRequest struct {
	Name         string `json:"name" binding:"required"`
	ID           string `json:"id" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
	Aircraft     string `json:"aircraft" binding:"required"`
}
