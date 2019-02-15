package models

type PayChannel struct {
	Id     int       `gorm:"column:id" json:"id"`
	Name   string    `gorm:"column:name" json:"name"`
	Code   string    `gorm:"column:code" json:"code"`
	Enable int       `gorm:"column:enable" json:"enable"`
	Types  []PayType `json:"types"`
}

func (self PayChannel) TableName() string {
	return "pay_channels"
}
