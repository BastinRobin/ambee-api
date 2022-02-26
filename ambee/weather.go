package ambee

import (
	"encoding/json"
	"fmt"
	"log"
)

type Weather struct {
	Message string
	Data    struct {
		Temperature float64
		DewPoint    float64 `json:"dewPoint"`
		Humidity    float64 `json:"humidity"`
	}
}

func (a *Ambee) GetWeather(lat float64, lng float64) (Weather, error) {

	url := fmt.Sprintf("/weather/latest/by-lat-lng?lat=%v&lng=%v", lat, lng)

	response, err := a.Request(url, "GET")
	if err != nil {
		log.Println(err)
		return Weather{}, err
	}

	var weather Weather
	err = json.Unmarshal(response, &weather)
	if err != nil {
		log.Println(err)
		return Weather{}, err
	}

	return weather, nil
}
