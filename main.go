package main

import (
	"net/http"

	"github.com/bookkyjung1221/stockradar_challenge/controllers"
	"github.com/bookkyjung1221/stockradar_challenge/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {

	utils.AddTransactionToDB()

	r := gin.Default()

	r.GET("/user/sale_amount/:user_id", controllers.GetUserSaleAmount)
	r.GET("/user/point/:user_id", controllers.GetPoints)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 Page Not Found"})
	})

	r.Run()
}
