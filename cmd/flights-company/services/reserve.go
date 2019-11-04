package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davidnastasi/flights-company/cmd/flights-company/models"
	_ "github.com/davidnastasi/flights-company/cmd/flights-company/models"
	"github.com/davidnastasi/flights-company/cmd/flights-company/repository"
	"io/ioutil"
	"net/http"
	"time"
)


func GetAll() error {

	resp, err := http.Get("https://flights-company-flights.herokuapp.com/flight-reservations")
	if err != nil {
		return  err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Can't get reserves from server")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return  err
	}

	reserves, err :=  parseJsonToReserves(data)
	if err != nil {
		return err
	}
	//log.Println( reserves)
 	err = repository.Save(reserves)
	if err != nil {
		return err
	}

 	return nil

}

func parseJsonToReserves(data []byte) ([]*models.Reserve,  error) { //([] *models.Hotel, error) {
	var result []interface{}
	var reserves []*models.Reserve
	err := json.Unmarshal(data, &result)
	if err != nil {
		return reserves, err
	}
	for _, arr := range result {
		value := arr.(map[string] interface{})
		v, err := time.Parse("2006-01-02T15:04:05.999999", fmt.Sprint(value["date"]))
		if err == nil {
			reserve := models.NewReserve(
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
