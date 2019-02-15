package models

type PayMethodMap struct {
	Id    int                `json:"id"`
	Name  string             `json:"name"`
	Code  string             `json:"code"`
	Types []PayMethodMapType `json:"types"`
}

type PayMethodMapType struct {
	Code        string `json:"code"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

type PayMethod struct {
	ChannelId       int    `json:"channel_id"`
	ChannelCode     string `json:"channel_code"`
	ChannelName     string `json:"channel_name"`
	TypeId          int    `json:"type_id"`
	TypeName        string `json:"type_name"`
	TypeCode        string `json:"type_code"`
	TypeDescription string `json:"type_description"`
}
