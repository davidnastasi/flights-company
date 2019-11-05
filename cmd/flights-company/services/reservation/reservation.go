package reservation

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/davidnastasi/flights-company/cmd/flights-company/models"
	"github.com/davidnastasi/flights-company/cmd/flights-company/repository"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/reservation/dto"
	"io/ioutil"
	"net/http"
	"time"
)


func GetAll() error {

	resp, err := http.Get("https://brubank-flights.herokuapp.com/flight-reservations")
	if err != nil {
		return  err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Can't get reserves from server. Status code: %v ", resp.StatusCode))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return  err
	}

	reserves, err :=  parseJsonToReserves(data)
	if err != nil {
		return err
	}

 	err = (&repository.PostgreSQL{}).Save(reserves)
	if err != nil {
		return err
	}

 	return nil

}

func parseJsonToReserves(data []byte) ([]*dto.ReservationDTO,  error) { //([] *models.Hotel, error) {
	var result []interface{}
	var reserves []*dto.ReservationDTO
	err := json.Unmarshal(data, &result)
	if err != nil {
		return reserves, err
	}
	for _, arr := range result {
		value := arr.(map[string] interface{})
		v, err := time.Parse("2006-01-02T15:04:05.999999", fmt.Sprint(value["date"]))
		if err == nil {
			reserve := dto.NewReservationDTO(
				v,
				fmt.Sprint(value["destination"]),
				fmt.Sprint(value["reservationId"]),
			)
			reserves = append(reserves, reserve)
		}else {
			fmt.Println(err)
			return reserves,err
		}
	}
	return reserves, nil
}
