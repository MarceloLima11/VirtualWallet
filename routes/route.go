package routes

import (
	"github.com/MarceloLima11/VirtualWallet/balance"
	"github.com/gin-gonic/gin"
)

const (
	path = "api/account"
)

func Init() error {
	r := gin.Default()
	balance.GetDatabaseInstance()

	v1 := r.Group(path)
	{
		v1.GET("/:id", balance.AccountBalance)
		v1.GET("/transfer/:debtorID/:beneficiaryID/:amount", balance.BalanceTransfer)
	}

	if err := r.Run(":8080"); err != nil {
		return err
	}

	return nil
}
