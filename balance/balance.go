package balance

import (
	"database/sql"

	"github.com/MarceloLima11/VirtualWallet/postgres"
	"github.com/shopspring/decimal"
)

var db *sql.DB

type (
	Account struct {
		Id      int64
		Balance decimal.Decimal
	}
)

func GetDatabaseInstance() {
	db = postgres.GetPostgreSQL()
}
