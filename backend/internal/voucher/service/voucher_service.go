package service

import (
	"fmt"
	"math/rand/v2"
	"time"

	"iotatfan.com/airline-voucher/internal/voucher/models"
	"iotatfan.com/airline-voucher/internal/voucher/repository"
)

type VoucherService interface {
	CheckFlight(flightNumber string, date string) (models.CheckFlightResponse, error)
	GenerateRandomSeats(name string, id string, flightNumber string, date string, aircraft string) (models.GenerateRandomSeatsResponse, error)
}

type voucherService struct {
	voucherRepo repository.VoucherRepository
}

func NewVoucherService(voucherRepo repository.VoucherRepository) VoucherService {
	return &voucherService{
		voucherRepo: voucherRepo,
	}
}

func (s *voucherService) CheckFlight(flightNumber string, date string) (models.CheckFlightResponse, error) {
	result, err := s.voucherRepo.CheckFlight(flightNumber, date)
	if err != nil {
		return models.CheckFlightResponse{Exists: false}, err
	}

	return models.CheckFlightResponse{Exists: result}, nil
}

func (s *voucherService) GenerateRandomSeats(name string, id string, flightNumber string, date string, aircraft string) (models.GenerateRandomSeatsResponse, error) {
	layout, exists := models.AircraftConfigs[aircraft]
	if !exists {
		return models.GenerateRandomSeatsResponse{}, fmt.Errorf("invalid aircraft type: %s", aircraft)
	}

	generatedSeats := make(map[string]bool)
	var result []string

	for len(result) < 3 {
		row := rand.IntN(layout.MaxRow-layout.MinRow+1) + layout.MinRow

		letterIndex := rand.IntN(len(layout.SeatLetters))
		letter := layout.SeatLetters[letterIndex]

		seatCode := fmt.Sprintf("%d%s", row, letter)
		if !generatedSeats[seatCode] {
			exists, err := s.voucherRepo.CheckSeat(seatCode, date)
			if err != nil {
				return models.GenerateRandomSeatsResponse{}, err
			}
			if !exists {
				generatedSeats[seatCode] = true
				result = append(result, seatCode)
			}
		}
	}

	voucher := models.Voucher{
		CrewName:     name,
		CrewID:       id,
		FlightNumber: flightNumber,
		FlightDate:   date,
		AircraftType: aircraft,
		Seat1:        result[0],
		Seat2:        result[1],
		Seat3:        result[2],
		Created_at:   time.Now().Format(time.RFC3339),
	}

	err := s.voucherRepo.GenerateRandomSeats(voucher)
	if err != nil {
		return models.GenerateRandomSeatsResponse{Success: false, Seats: nil}, err
	}

	return models.GenerateRandomSeatsResponse{
		Success: true,
		Seats:   []string{voucher.Seat1, voucher.Seat2, voucher.Seat3},
	}, nil
}
