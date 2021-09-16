package main

import (
	"log"

	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	data, code, err := f.GetCitiesCountries()
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, data)
}
