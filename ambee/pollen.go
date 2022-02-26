package ambee

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pollen struct {
	Message string `json:"message"`
	Data    []struct {
		Count struct {
			GrassPollen int `json:"grass_pollen"`
			TreePollen  int `json:"tree_pollen"`
			WeedPollen  int `json:"weed_pollen"`
		} `json:"Count"`
		Risk struct {
			GrassPollen string `json:"grass_pollen"`
			TreePollen  string `json:"tree_pollen"`
			WeedPollen  string `json:"weed_pollen"`
		} `json:"Risk"`
	} `json:"data"`
}

func (a *Ambee) GetPollen(lat float64, lng float64) (Pollen, error) {

	url := fmt.Sprintf("/latest/pollen/by-lat-lng?lat=%v&lng=%v", lat, lng)

	response, err := a.Request(url, "GET")
	if err != nil {
		log.Println(err)
		return Pollen{}, err
	}

	var pollen Pollen
	err = json.Unmarshal(response, &pollen)
	if err != nil {
		log.Println(err)
		return Pollen{}, err
	}

	return pollen, nil
}
