package pay

type PayInstacne struct {
	OrderId     string `json:"order_id"`
	ChannelCode string `json:"channel_code"`
	TypeCode    string `json:"type_code"`
}

type PayAbstract interface {
	GetPayInfo()
	GetPayCallback()
}
