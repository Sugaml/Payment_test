package models

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/paymentdb?sslmode=disable"
)

var TestDB *gorm.DB

func TestMain(m *testing.M) {
	conn, err := gorm.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	TestDB = conn
	os.Exit(m.Run())
}
