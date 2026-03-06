package handlers

import (
	"money-manager/internal/db"
	"money-manager/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {

	var tx models.Transaction

	if err := c.BindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	var account models.Account

	if err := db.DB.First(&account, tx.AccountID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "account not found",
		})
		return
	}

	if tx.Type == "debit" {
		account.Balance -= tx.Amount
	} else if tx.Type == "credit" {
		account.Balance += tx.Amount
	}

	db.DB.Save(&account)
	db.DB.Create(&tx)

	c.JSON(http.StatusOK, tx)
}

func GetTransactions(c *gin.Context) {

	var txs []models.Transaction

	db.DB.Find(&txs)

	c.JSON(http.StatusOK, txs)
}
