package controllers

import (
	"log"
	"strconv"

	"github.com/bookkyjung1221/stockradar_challenge/utils"
	"github.com/gin-gonic/gin"
)

func GetUserSaleAmount(c *gin.Context) {

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	userId := c.Param("user_id")

	if userId == "" {
		c.JSON(400, gin.H{"error": "user_id not found please enter user_id"})
		return
	}

	var totalSale float64

	errQuery := db.QueryRow(`
            SELECT SUM(sale_amount * er.rate) AS total_sale
            FROM transactions s
            JOIN exchange_rates er ON s.currency = er.currency
            WHERE s.user_id = $1
        `, userId).Scan(&totalSale)

	if errQuery != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	strNum := strconv.FormatFloat(totalSale, 'f', 2, 64)

	totalamount := strNum + " THB"

	c.JSON(200, gin.H{"total_sale_amouunt": totalamount})
}
