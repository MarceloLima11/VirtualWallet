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

	if debtorID == beneficiaryID {
		sendError(ctx, http.StatusBadRequest, "cannot transfer to the same account")
		return
	}

	amount, err := decimal.NewFromString(amountParam)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "amount format invalid")
		return
	}

	var debtor Account
	debtorRow := db.QueryRow("SELECT id, balance FROM account WHERE id = $1", debtorID)
	if err := debtorRow.Scan(&debtor.Id, &debtor.Balance); err != nil {
		sendError(ctx, http.StatusNotFound, "debtor account not found")
		return
	}

	if debtor.Balance.LessThan(amount) {
		sendError(ctx, http.StatusBadRequest, "insufficient funds")
		return
	}

	var beneficiary Account
	beneficiaryRow := db.QueryRow("SELECT id, balance FROM account WHERE id = $1", beneficiaryID)
	if err := beneficiaryRow.Scan(&beneficiary.Id, &beneficiary.Balance); err != nil {
		sendError(ctx, http.StatusNotFound, "beneficiary account not found")
		return
	}

	if err := transfer(&debtor, &beneficiary, amount); err != nil {
		sendError(ctx, http.StatusInternalServerError, "internal error, please try again later")
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account balance updated successfully",
	})
}

func transfer(debtor *Account, beneficiary *Account, amount decimal.Decimal) error {
	var err error

	debtorBalance := debtor.Balance.Sub(amount)
	beneficiaryBalance := beneficiary.Balance.Add(amount)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE account SET balance = $1 WHERE id = $2", debtorBalance, debtor.Id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE account SET balance = $1 WHERE id = $2", beneficiaryBalance, beneficiary.Id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
