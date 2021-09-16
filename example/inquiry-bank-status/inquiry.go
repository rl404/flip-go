package main

import (
	"log"

	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	bank, code, err := f.InquiryBankAccount(flip.InquiryBankAccountRequest{
		AccountNumber: "5465327020",
		BankCode:      "bca",
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, bank)
}
