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
