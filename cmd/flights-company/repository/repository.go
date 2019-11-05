package repository

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/models"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/reservation/dto"
)

type Repository interface {
	GetAll(location string) ([]*dto.ReservationDTO, error)
	Save(reservesDTO []*dto.ReservationDTO) error
	GetLocation(location string) *models.Location
}