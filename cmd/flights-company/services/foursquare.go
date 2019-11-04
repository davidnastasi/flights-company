package services

import (
	"encoding/json"
	"fmt"
	"github.com/davidnastasi/flights-company/cmd/flights-company/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const CLIENT_ID  = "HACVIHTUOMFKVK5HWQ0J0JCOKQAA2CSAVFS0LFQVN14EESS2"
const CLIENT_SECRET = "50ITRVSKRB1GH2YWOBBQWZS5BEDVEIWN3Z2YABJEI454V2JZ"

func GetHotels(destination string) ([]models.Hotel, error) {
	query, err := generateQuery(destination)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(query)
	if err != nil {
		return nil , err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil , err
	}

	return parseJsonToHotels(data)
}

func generateQuery(destination string) (string, error) {

	baseUrl, err := url.Parse("https://api.foursquare.com/v2/venues/search?")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return "", err
	}

	params := url.Values{}
	params.Add("near", destination)
	params.Add("intent", "browse")
	params.Add("query", "hotel")
	params.Add("client_id", CLIENT_ID)
	params.Add("client_secret", CLIENT_SECRET)
	params.Add("v", "20190709")

	baseUrl.RawQuery = params.Encode() // Escape Query Parameters
	log.Println(baseUrl.String())

	return baseUrl.String(),nil
}

func parseJsonToHotels(data []byte) ([]models.Hotel, error) {
	var result map[string]interface{}
	err :=json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	var hotels []models.Hotel
	//fmt.Println("%s", result)
	response := result["response"].(map[string]interface{})
	//fmt.Println("%s", response)
	venues := response["venues"].([] interface{})

	//fmt.Printf("%v", venue)
	//venues := response["venues"].(map[string] []interface{})
	//fmt.Println("%s", venues)
	for _, value := range venues {
		v := value.(map[string]interface{})
		//fmt.Printf("%s\n", v["name"])
		location := v["location"].(map[string] interface{})
		//fmt.Printf("%s\n", location["address"])
		hotel := models.NewHotel(fmt.Sprint(v["name"]),fmt.Sprint(location["address"]))
		hotels = append(hotels, *hotel)
	}
	return hotels,nil
}
