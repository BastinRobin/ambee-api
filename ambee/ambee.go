package ambee

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Ambee struct {
	BaseUrl string
	Header  map[string][]string
}

func (a *Ambee) Request(url string, method string) ([]byte, error) {

	// Create the full length URL
	url = fmt.Sprintf("%v%v", a.BaseUrl, url)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Reassign headers from the struct
	req.Header = a.Header

	// Create a client to perform http
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}
