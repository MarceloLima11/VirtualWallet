package balance

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	User struct {
		Id      int64
		Balance float64
	}
)

func UserBalance(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Print(id)
}
