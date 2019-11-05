package dto

import "time"

type ReservationDTO struct {
	Date time.Time `json:"date"`
	Destination string `json:"destination"`
	Reservation string `json:"reservation"`
}

// NewReserve construye un nuevo reserva
func NewReservationDTO(date time.Time, destination string, reservationId string) *ReservationDTO {
	return &ReservationDTO{Date: date, Destination: destination, Reservation: reservationId}
}