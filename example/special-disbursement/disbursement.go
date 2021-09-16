package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	tx, code, err := f.CreateSpecialDisbursement(flip.CreateSpecialDisbursementRequest{
		IdempotencyKey:       uuid.New().String(),
		AccountNumber:        "5465327020",
		BankCode:             "bca",
		Amount:               100000,
		Remark:               "remark example",
		RecipientCity:        391,
		SenderCountry:        100252,
		SenderPlaceOfBirth:   391,
		SenderDateOfBirth:    "1992-01-01",
		SenderIdentityType:   flip.IdentityNationalID,
		SenderName:           "naruto",
		SenderAddress:        "konohagakure",
		SenderIdentityNumber: "123456",
		SenderJob:            flip.JobFoundationBoard,
		Direction:            flip.DirectionDomestic,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, tx)
}
