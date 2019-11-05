package repository

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/config"
	"github.com/davidnastasi/flights-company/cmd/flights-company/models"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/reservation/dto"
	"time"
)

type PostgreSQL struct {}

func NewPostgreSQL() *PostgreSQL {
	return &PostgreSQL{}
}

// Get obtiene todas las reservas
func (p *PostgreSQL) GetAll(location string) ([]*dto.ReservationDTO, error) {
	var reserves []*dto.ReservationDTO

	rows, err := config.Config.DB.Table("locations").
		Select("locations.name, reservations.reservation, reservations.date").
		Joins("join reservations on locations.id = reservations.location_id").
		Where("locations.name = ?", location).
		Order("reservations.date").
		Rows()
	if err != nil{
		return reserves, err
	}
	for rows.Next(){
		var date time.Time
		var destination string
		var reservationId string
		if err = rows.Scan(&destination, &reservationId, &date); err != nil{
			break
		}
		reservation := dto.ReservationDTO{
			Destination: destination,
			Date: date,
			Reservation:reservationId,
		}

		reserves = append(reserves, &reservation)
	}

	if err != nil {
		return nil, err
	}

	return reserves, err
}

// Save salva las reservas
func (p *PostgreSQL) Save(reservesDTO []*dto.ReservationDTO) error {
	var err error
	for _, r := range reservesDTO {
		loc := p.GetLocation(r.Destination)
		res := models.NewReservation(
			r.Date,
			loc.ID,
			loc,
			r.Reservation,
		)
		config.Config.DB.NewRecord(&res)
		if res := config.Config.DB.Create(res) ; res.Error != nil {
			err = res.Error
		}
	}
	return err
}

func (p *PostgreSQL) GetLocation(location string) *models.Location {
	var l models.Location
	err := config.Config.DB.Where("name = ?", location).
		First(&l).
		Error
	if err != nil {
		l.Name = location
		config.Config.DB.NewRecord(l)
		config.Config.DB.Create(&l)
	}

	return &l

}

