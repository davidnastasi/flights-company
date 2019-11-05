package repository

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/reservation/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInMemory_SaveAndGet(t *testing.T){

	inMemory := NewInMemory()
	rs := make([]*dto.ReservationDTO,0)
	r := dto.ReservationDTO{
		Date:time.Now(),
		Destination:"Buenos Aires, Argentina",
		Reservation: "735eb398-8724-403d-b192-38a837815a3d",
	}

	rs = append(rs, &r)
	err := inMemory.Save(rs)

	rsa, err := inMemory.GetAll("Buenos Aires, Argentina")

	assert.Nil(t, err)
	assert.NotEmpty(t, rsa)

}


func TestInMemory_GetNoDestination(t *testing.T){

	inMemory := NewInMemory()
	rsa, err := inMemory.GetAll("Buenos Aires, Argentina")

	assert.Nil(t, err)
	assert.Empty(t, rsa)

}

