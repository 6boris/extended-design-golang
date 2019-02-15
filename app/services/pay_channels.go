package services

import "extended-design-golang/app/models"

func GetAllPayChannels() []models.PayChannel {
	pay_channels := []models.PayChannel{}
	db.Table("pay_channels").Find(&pay_channels)
	return pay_channels
}

func GetPayChannelMap() {

}
