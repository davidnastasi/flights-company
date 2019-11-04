package models

import "time"

// Reserve estructura para la reserva
type Reserve struct {
	Date time.Time `gorm:"column:date" json:"date"`
	Destination string `gorm:"column:destination" json:"destination"`
	ReservationId string `gorm:"column:reservationId" json:"reservationId"`
}

// NewReserve construye un nuevo reserva
func NewReserve(date time.Time, destination string, reservationId string) *Reserve {
	return &Reserve{Date: date, Destination: destination, ReservationId: reservationId}
}



