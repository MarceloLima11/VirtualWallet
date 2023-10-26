package balance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccountBalance(ctx *gin.Context) {
	id := ctx.Param("id")

	var account Account

	row := db.QueryRow("SELECT balance FROM account WHERE id = $1", id)

	if err := row.Scan(&account.Balance); err != nil {
		ctx.JSON(http.StatusNotFound, "account not found")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"balance": account.Balance,
	})
}
