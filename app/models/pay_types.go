package models

type PayType struct {
	Id          int    `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Code        string `gorm:"column:code" json:"code"`
	Description string `gorm:"column:description" json:"description"`
}

func (self PayType) TableName() string {
	return "pay_types"
}
