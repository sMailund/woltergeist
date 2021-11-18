package domain

import "github.com/sMailund/woltergeist/external/wolt-api-bindings/entities"

func GetRestaurants() (*entities.Restaurants, error) {
	r := new(entities.Restaurants)

	return r, nil
}
