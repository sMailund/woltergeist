package domain

import (
	"fmt"
	"github.com/sMailund/woltergeist/external/wolt-api-bindings/entities"
	"http"
	"io"
	"json"
)

func GetRestaurants() (*entities.Restaurants, error) {
	var restaurants entities.Restaurants
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &restaurants); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return &restaurants, nil
}
