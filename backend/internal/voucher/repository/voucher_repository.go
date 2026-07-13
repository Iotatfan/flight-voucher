package repository

import (
	"gorm.io/gorm"
	"iotatfan.com/airline-voucher/internal/voucher/models"
)

type VoucherRepository interface {
	CheckFlight(flightNumber string, date string) (bool, error)
	GenerateRandomSeats(models.Voucher) error
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{
		db: db,
	}
}

func (r *voucherRepository) CheckFlight(flightNumber string, date string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Voucher{}).Where("flight_number = ? AND flight_date = ?", flightNumber, date).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *voucherRepository) GenerateRandomSeats(v models.Voucher) error {
	return r.db.Create(&v).Error
}
