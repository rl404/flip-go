package main

import (
	"log"

	"github.com/rl404/flip-id"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	balance, err := f.GetBalance()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(balance)
}
