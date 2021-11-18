package main

import (
	"fmt"

	"github.com/sMailund/woltergeist/external/wolt-api-bindings/domain"
)

func main() {
	rest, err := domain.GetRestaurants()
	if err != nil {
		println(err)
	}
	response := rest
	fmt.Printf("%+v\n", response)
}
