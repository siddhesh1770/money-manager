package handlers

import (
	"money-manager/internal/db"
	"money-manager/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {

	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	db.DB.Create(&account)

	c.JSON(http.StatusOK, account)
}

func GetAccounts(c *gin.Context) {

	var accounts []models.Account

	db.DB.Find(&accounts)

	c.JSON(http.StatusOK, accounts)
}