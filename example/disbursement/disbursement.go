package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	tx, err := f.CreateDisbursement(flip.CreateDisbursementRequest{
		IdempotencyKey: uuid.New().String(),
		AccountNumber:  "5465327020",
		BankCode:       "bca",
		Amount:         100000,
		Remark:         "remark example",
		RecipientCity:  391,
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(tx)

	tx, err = f.GetDisbursement(tx.ID)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(tx)

	txs, err := f.GetDisbursements(flip.GetDisbursementsRequest{
		Pagination: 10,
		Page:       1,
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(txs)
}
