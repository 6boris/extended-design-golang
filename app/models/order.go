package models

import "time"

type Order struct {
	Id           int        `gorm:"column:id" json:"id"`
	UserId       int        `gorm:"column:user_id" json:"user_id"`
	PayTypeId    int        `gorm:"column:pay_type_id" json:"pay_type_id"`
	PayChannelId int        `gorm:"column:pay_channel_id" json:"pay_channel_id"`
	OrderNnm     string     `gorm:"column:order_num" json:"order_num"`
	Account      float64    `gorm:"column:account" json:"account"`
	Status       string     `gorm:"column:status" json:"status"`
	Created_at   *time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Finished_at  *time.Time `gorm:"column:finished_at" json:"finished_at"`
}

func (self Order) TableName() string {
	return "orders"
}

type OrderInfo struct {
	OrderNum      string  `json:"order_num"`
	Account       float64 `json:"account"`
	PayTypeCode   string  `json:"pay_type_code"`
	PayTypeName   string  `json:"pay_type_name"`
	PayMethodCode string  `json:"pay_method_code"`
	PayMethodName string  `json:"pay_method_name"`

	Created_at  *time.Time `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at"`
	Finished_at *time.Time `json:"finished_at"`
}
