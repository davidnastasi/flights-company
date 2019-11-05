package repository

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/models"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/reservation/dto"
)

type InMemory struct {
	data []*dto.ReservationDTO
}

func NewInMemory() *InMemory {
	return &InMemory{data: make([]*dto.ReservationDTO,0)}
}

func (i *InMemory) GetAll(location string) ([]*dto.ReservationDTO, error){
	var result []*dto.ReservationDTO
	for _, v := range i.data{
		if v.Destination == location {
			result = append(result, v)
		}
	}
	return result, nil
}

func (i *InMemory) Save(reservesDTO []*dto.ReservationDTO) error {
	for _, v := range reservesDTO {
			i.data = append(i.data, v)
	}
	return nil
}

func (i *InMemory) GetLocation(location string) *models.Location {
	return nil
}

