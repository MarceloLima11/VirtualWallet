package balance

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func AccountBalance(ctx *gin.Context) {
	id := ctx.Param("id")

	var balance decimal.Decimal

	row := db.QueryRow("SELECT balance FROM account WHERE id = $1", id)

	if err := row.Scan(&balance); err != nil {
		ctx.JSON(http.StatusNotFound, "account not found")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}
