package ambee

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Message string
	Data    Datas
}

type Datas []Data

type Data struct {
	Confidence string  `json:"_id"`
	Frp        float64 `json:"frp"`
}

func (a *Ambee) GetFire(lat float64, lng float64) (Response, error) {
	url := fmt.Sprintf("/latest/fire?lat=%v&lng=%v", lat, lng)

	response, err := a.Request(url, "GET")

	if err != nil {
		log.Println(err)
		return Response{}, err
	}

	var datas Response
	err = json.Unmarshal(response, &datas)
	if err != nil {
		log.Println(err)
		return Response{}, err
	}

	return datas, nil

}
