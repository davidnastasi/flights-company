package dto

import "time"

// ReserveDTO estructura dto para las reservas
type ReserveDTO struct {
	Date time.Time `json:"date"`
	ReservationId string `json:"reservationId"`
}

// HotelDTO estructura dto para los hotel
type HotelDTO struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

// DestinationDTO estructura dto para los destinos
type DestinationDTO struct {
	Reserves []ReserveDTO `json:"reserves"`
	Hotels []HotelDTO `json:"hotels"`
}

