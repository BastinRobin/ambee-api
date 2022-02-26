package ambee

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
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

func (a *Ambee) GetAirQuality(country string) (Response, error) {
	url := fmt.Sprintf("/latest/by-country-code?countryCode=%v", country)

	response, err := a.Request(url, "GET")

	if err != nil {
		log.Println(err)
		return Response{}, err
	}

	var stations Response
	err = json.Unmarshal(response, &stations)
	if err != nil {
		log.Println(err)
		return Response{}, err
	}

	return stations, nil

}
