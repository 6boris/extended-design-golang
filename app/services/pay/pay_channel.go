package pay

import (
	"extended-design-golang/app/models"
	"extended-design-golang/app/util"
	"fmt"
)

type PayInstacne struct {
	OrderId     string `json:"order_id"`
	ChannelCode string `json:"channel_code"`
	TypeCode    string `json:"type_code"`
}

//
type PayChannelInterface interface {
	GetPayInstance()
	GetPayInfo() string
	GetPayCallback() string
}

func GetPayChannel(order models.PayInfo) string {
	//channel := pay_channel.GetPayChannelInstance(order)
	channel := GetPayChannelInstance(order)

	return channel.GetPayInfo()
}

func GetPayChannelInstance(channel models.PayInfo) PayChannelInterface {

	fmt.Println(string(util.JsonMarshal(channel)))

	switch channel.PayMethod.ChannelCode {
	case "kuaijiepay":
		return new(PayChannelKuaiJiePay)
	case "xinbeipay":
		return new(PayChannelXinBeiPay)
	case "yipay":
		return new(PayChannelYiPay)
	default:
		return nil
	}
}
