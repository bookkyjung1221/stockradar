package controllers

import (
	"log"
	"strconv"

	"github.com/bookkyjung1221/stockradar_challenge/utils"
	"github.com/gin-gonic/gin"
)

func GetPoints(c *gin.Context) {

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

	var totalPoints float64

	errQuery := db.QueryRow(`
            SELECT SUM(payout * 0.2) AS total_points
            FROM transactions
            WHERE user_id = $1
        `, userId).Scan(&totalPoints)

	if errQuery != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	strNum := strconv.FormatFloat(totalPoints, 'f', 2, 64)

	totalpoints := strNum + " points"

	c.JSON(200, gin.H{"total_points": totalpoints})
}
