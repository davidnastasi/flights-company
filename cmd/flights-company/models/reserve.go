package models

import "time"


// Reservation estructura para la reserva
type Reservation struct {
	Date time.Time `gorm:"column:date" json:"date"`
	LocationId int64 `gorm:"column:location_id"`
	Location *Location `json:"location"`
	ReservationId string `gorm:"column:reservation" json:"reservation"`
}

// NewReservation construye un nuevo reserva
func NewReservation(date time.Time, locationId int64, location *Location, reservationId string) *Reservation {
	return &Reservation{Date: date, LocationId: locationId, Location: location, ReservationId: reservationId}
}



