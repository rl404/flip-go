package main

import (
	"log"

	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	balance, code, err := f.GetBalance()
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, balance)
}
