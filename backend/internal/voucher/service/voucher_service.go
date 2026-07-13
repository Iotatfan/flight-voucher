package service

import (
	"fmt"
	"math/rand/v2"
	"time"

	aircraftModels "iotatfan.com/airline-voucher/internal/aircraft/models"
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
	layout, exists := aircraftModels.AircraftConfigs[aircraft]
	if !exists {
		return models.GenerateRandomSeatsResponse{}, models.ErrInvalidAircraftType
	}

	alreadyExists, err := s.voucherRepo.CheckFlight(flightNumber, date)
	if err != nil {
		return models.GenerateRandomSeatsResponse{}, err
	}
	if alreadyExists {
		return models.GenerateRandomSeatsResponse{}, models.ErrVoucherAlreadyExists
	}

	var availableSeats []string
	for row := layout.MinRow; row <= layout.MaxRow; row++ {
		for _, letter := range layout.SeatLetters {
			availableSeats = append(availableSeats, fmt.Sprintf("%d%s", row, letter))
		}
	}

	if len(availableSeats) < 3 {
		return models.GenerateRandomSeatsResponse{}, models.ErrFewerThan3SeatsAvailable
	}

	indices := rand.Perm(len(availableSeats))
	result := []string{
		availableSeats[indices[0]],
		availableSeats[indices[1]],
		availableSeats[indices[2]],
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

	err = s.voucherRepo.GenerateRandomSeats(voucher)
	if err != nil {
		return models.GenerateRandomSeatsResponse{Success: false, Seats: nil}, models.ErrVoucherAlreadyExists
	}

	return models.GenerateRandomSeatsResponse{
		Success: true,
		Seats:   result,
	}, nil
}
