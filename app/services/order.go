package services

import (
	"extended-design-golang/app/models"
)

//	查询所有订单
func GetAllOrders() []models.Order {
	orders := []models.Order{}
	db.Table("orders").Find(&orders)
	return orders
}

//	查询单个订单
func GetOneOrder(order_num string) models.Order {
	order := models.Order{}
	order.OrderNnm = order_num

	db.Table("orders").Where("order_num = ?", order_num).Find(&order)
	return order
}
