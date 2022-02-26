package ambee

import (
	"encoding/json"
	"fmt"
	"log"
)

type Responses struct {
	Message  string
	Stations Stations
}

type Stations []Station

type Station struct {
	Id      string `json:"_id"`
	PlaceId string
	CO      float64
	NO2     float64
}

func (a *Ambee) GetAirQuality(country string) (Responses, error) {
	url := fmt.Sprintf("/latest/by-country-code?countryCode=%v", country)

	response, err := a.Request(url, "GET")

	if err != nil {
		log.Println(err)
		return Responses{}, err
	}

	var stations Responses
	err = json.Unmarshal(response, &stations)
	if err != nil {
		log.Println(err)
		return Responses{}, err
	}

	return stations, nil

}
