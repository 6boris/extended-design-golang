package services

import "extended-design-golang/app/models"

func GetAllPayMapChannelType() []models.PayMapChannelType {
	pay_map_channel_type := []models.PayMapChannelType{}
	db.Table("pay_map_channel_type").Find(&pay_map_channel_type)
	return pay_map_channel_type
}
