package models

import "time"

type PayMapChannelType struct {
	Id          int        `gorm:"column:id" json:"id"`
	ChannelCode string     `gorm:"column:channel_code" json:"channel_code"`
	TypeCode    string     `gorm:"column:type_code" json:"type_code"`
	Created_at  *time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (self PayMapChannelType) TableName() string {
	return "pay_map_channel_type"
}
