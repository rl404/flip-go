package main

import (
	"log"

	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	banks, code, err := f.GetBanks(flip.BankAceh)
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, banks)
}
