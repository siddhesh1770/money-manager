package routes

import (
	"money-manager/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(rg *gin.RouterGroup) {

    accounts := rg.Group("/accounts")

    accounts.POST("/", handlers.CreateAccount)
    accounts.GET("/", handlers.GetAccounts)
}