package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func storeCard(req PaymentRequest) error {
	_, err := db.Exec(`INSERT INTO payment_cards (tenant_id, name, card_number, expiry, cvv) VALUES ($1,$2,$3,$4,$5)`,
		req.TenantID, req.OrganizationName, req.CardNumber, req.Expiry, req.CVV)
	return err
}
