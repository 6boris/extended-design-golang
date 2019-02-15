package controllers

import (
	"extended-design-golang/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOrders(c *gin.Context) {
	orders := services.GetAllOrders()

	c.JSON(http.StatusOK, gin.H{
		"orders:": orders,
	})
}

func GetOneOrder(c *gin.Context) {
	order_num := c.Param("order_num")

	order := services.GetOneOrder(order_num)

	c.JSON(http.StatusOK, gin.H{
		"order:": order,
	})
}

func CreateOrder(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"order:": "CreateOrder",
	})
}
