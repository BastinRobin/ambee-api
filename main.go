package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bastinrobin/ambee-api/ambee"
)

func main() {
	// We have created a reusable ambee client
	ambee := ambee.Ambee{
		BaseUrl: "https://api.ambeedata.com",
		Header: http.Header{
			"x-api-key":    []string{"7c84823be075ece19092f8719ff6bf1168547cf6abea491982b65993a550eeca"},
			"Content-Type": []string{"application/json"},
		},
	}
	fire, err := ambee.GetFire(12.06552, 77.05711)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(fire.Data)

}
