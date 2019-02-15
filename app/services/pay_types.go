package services

import "extended-design-golang/app/models"

func GetAllPayTypes() []models.PayType {
	pay_types := []models.PayType{}

	db.Table("pay_types").Find(&pay_types)

	return pay_types
}
