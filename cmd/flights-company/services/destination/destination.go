package destination

import (
	"errors"
	"github.com/davidnastasi/flights-company/cmd/flights-company/repository"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/destination/dto"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/hotels"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type DestinationService struct {
	Repository    repository.Repository
	cacheReservas *cache.Cache
	cacheHotels   *cache.Cache
}

func NewDestinationService(repository repository.Repository) *DestinationService {
	service := DestinationService{
		repository,
		cache.New(15*time.Minute, 30*time.Minute),
		cache.New(48*time.Hour, 72*time.Hour),
	}
	return &service
}

func (d *DestinationService) Get(destination string) (dto.DestinationDTO, error) {
	//var destinations dto.DestinationDTO
	var wg sync.WaitGroup
	var wgerr error
	var reservesDTO []dto.ReserveDTO
	var hotelsDTO []dto.HotelDTO

	wg.Add(1)

	go func() {
		r, err :=  d.getReserves(destination, &wg)
		if  err != nil {
			wgerr = err
		}else {
			reservesDTO = r
		}

	}()

	wg.Add(1)
	go func() {
		h, err :=  d.getHotels(destination, &wg)
		if  err != nil {
			wgerr = err
		}else {
			hotelsDTO = h
		}
	}()

	wg.Wait()

	if wgerr != nil {
		return dto.DestinationDTO{}, wgerr
	}

	if len(reservesDTO) == 0 {
		return  dto.DestinationDTO{}, errors.New("No reservertions found for destination")
	}

	return dto.DestinationDTO{
		reservesDTO,
		hotelsDTO,
	}, nil



}

func (d *DestinationService) getReserves(destination string, wg *sync.WaitGroup) ([]dto.ReserveDTO, error) {
	reservesDTO := make([]dto.ReserveDTO, 0)
	reservasCache, found := d.cacheReservas.Get(destination)
	if found {
		reservesDTO = reservasCache.([]dto.ReserveDTO)
	} else {
		reserves, err := d.Repository.GetAll(destination)
		if err != nil {
			wg.Done()
			return reservesDTO, err
		} else {
			for _, reserve := range reserves {
				reserveDTO := dto.ReserveDTO{
					Date:          reserve.Date,
					ReservationId: reserve.Reservation,
				}
				reservesDTO = append(reservesDTO, reserveDTO)
			}
			d.cacheReservas.Set(destination, reservesDTO, cache.DefaultExpiration)
		}
	}
	wg.Done()
	return reservesDTO, nil
}

func (d *DestinationService) getHotels(destination string, wg *sync.WaitGroup) ([]dto.HotelDTO, error) {
	hotelsDTO := make([]dto.HotelDTO, 0)
	hotelsCache, found := d.cacheHotels.Get(destination)
	if found {
		hotelsDTO = hotelsCache.([]dto.HotelDTO)
	} else {
		if hotels, err := hotels.GetHotels(destination); err != nil {
			wg.Done()
			return hotelsDTO, err
		} else {
			for _, hotel := range hotels {
				hotelDTO := dto.HotelDTO{
					Name:    hotel.Name,
					Address: hotel.Address,
				}
				hotelsDTO = append(hotelsDTO, hotelDTO)
			}
			d.cacheHotels.Set(destination, hotelsDTO, cache.DefaultExpiration)
		}
	}
	wg.Done()
	return hotelsDTO, nil
}
