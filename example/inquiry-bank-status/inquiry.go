package main

import (
	"log"

	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	bank, err := f.InquiryBankAccount(flip.InquiryBankAccountRequest{
		AccountNumber: "5465327020",
		BankCode:      "bca",
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(bank)
}
