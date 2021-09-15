package main

import (
	"log"

	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	cities, err := f.GetCities()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(cities)
}
