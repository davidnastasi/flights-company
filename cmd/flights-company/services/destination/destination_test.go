package destination

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/config"
	"github.com/davidnastasi/flights-company/cmd/flights-company/repository"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/reservation/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDestination_Get(t *testing.T){

	config.Config.FoursquareClientId= "HACVIHTUOMFKVK5HWQ0J0JCOKQAA2CSAVFS0LFQVN14EESS2"
	config.Config.FoursquareClientSecret="50ITRVSKRB1GH2YWOBBQWZS5BEDVEIWN3Z2YABJEI454V2JZ"

	inMemory:=repository.NewInMemory()
	rs := make([]*dto.ReservationDTO,0)
	r := dto.ReservationDTO{
		Date:time.Now(),
		Destination:"Buenos Aires, Argentina",
		Reservation: "735eb398-8724-403d-b192-38a837815a3d",
	}
	rs = append(rs, &r)
	err := inMemory.Save(rs)


	service := NewDestinationService(inMemory)
	destination , err := service.Get("Buenos Aires, Argentina")

	assert.Nil(t, err)
	assert.NotEmpty(t, destination)
}


func TestDestination_GetNoDestination(t *testing.T){

	config.Config.FoursquareClientId= "HACVIHTUOMFKVK5HWQ0J0JCOKQAA2CSAVFS0LFQVN14EESS2"
	config.Config.FoursquareClientSecret="50ITRVSKRB1GH2YWOBBQWZS5BEDVEIWN3Z2YABJEI454V2JZ"

	inMemory:=repository.NewInMemory()
	rs := make([]*dto.ReservationDTO,0)
	r := dto.ReservationDTO{
		Date:time.Now(),
		Destination:"Buenos Aires, Argentina",
		Reservation: "735eb398-8724-403d-b192-38a837815a3d",
	}
	rs = append(rs, &r)
	err := inMemory.Save(rs)


	service := NewDestinationService(inMemory)
	destination , err := service.Get("BueAires, Argentia")

	assert.NotNil(t, err)
	assert.Empty(t, destination)
}