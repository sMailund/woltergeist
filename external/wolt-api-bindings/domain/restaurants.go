package domain

import (
	"net/http"
	"io"
	"encoding/json"
	"github.com/sMailund/woltergeist/external/wolt-api-bindings/entities"
)

func GetRestaurants() (*entities.Restaurants, error) {
	var restaurants entities.Restaurants
	resp, err := http.Get("https://restaurant-api.wolt.com/v1/pages/restaurants?lat=59.916538&lon=10.738423")
	if err != nil {
		return &restaurants, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &restaurants); err != nil { // Parse []byte to the go struct pointer
		return &restaurants, err
	}
	return &restaurants, nil
}
