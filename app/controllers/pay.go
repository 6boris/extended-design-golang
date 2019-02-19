package controllers

import (
	"extended-design-golang/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllPayMaps(c *gin.Context) {
	pay_maps1 := services.GetPayMethodMap()
	pay_maps2 := services.GetAllPay()

	c.JSON(http.StatusOK, gin.H{
		"pay_maps1": pay_maps1,
		"pay_maps2": pay_maps2,
	})
}

func GetPayInfo(c *gin.Context) {
	order_num := c.Param("order_num")

	pay_info := services.GetPayInfo(order_num)

	c.JSON(200, gin.H{
		"pay_info": pay_info,
	})

}
