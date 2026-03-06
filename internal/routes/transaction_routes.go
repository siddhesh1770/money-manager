package routes

import (
	"money-manager/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(rg *gin.RouterGroup) {

    transactions := rg.Group("/transactions")

    transactions.POST("/", handlers.CreateTransaction)
    transactions.GET("/", handlers.GetTransactions)
}