package balance

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func BalanceTransfer(ctx *gin.Context) {
	debtorID := ctx.Param("debtorID")
	beneficiaryID := ctx.Param("beneficiaryID")
	amountParam := ctx.Param("amount")

	amount, err := decimal.NewFromString(amountParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "amount format invalid",
		})
	}

	var debtor Account
	debtorRow := db.QueryRow("SELECT id, balance FROM account WHERE id = $1", debtorID)
	if err := debtorRow.Scan(&debtor.Id, &debtor.Balance); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "debtor account not found",
		})
	}

	if debtor.Balance.LessThan(amount) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "insufficient funds",
		})
	}

	var beneficiary Account
	beneficiaryRow := db.QueryRow("SELECT id, balance FROM account WHERE id = $1", beneficiaryID)
	if err := beneficiaryRow.Scan(&beneficiary.Id, &beneficiary.Balance); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "beneficiary account not found",
		})
	}

	if err := addBalance(&beneficiary, amount); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error, please try again later",
		})
	}
}

func addBalance(account *Account, amount decimal.Decimal) error {
	account.Balance.Add(amount)
	balance, _ := account.Balance.Float64()

	_, err := db.Exec("UPDATE account SET balance = $1 WHERE = $2", balance, account.Id)
	if err != nil {
		return err
	}

	return nil
}
