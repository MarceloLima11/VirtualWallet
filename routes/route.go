package routes

import (
	"github.com/MarceloLima11/VirtualWallet/balance"
	"github.com/gin-gonic/gin"
)

const (
	path = "api/user"
)

func Init() {
	r := gin.Default()

	v1 := r.Group(path)
	{
		v1.GET("/:id", balance.UserBalance)
		v1.GET("/:sender/:receiver")
	}

	r.Run("8080")
}
