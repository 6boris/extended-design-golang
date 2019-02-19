package pay_channel

import (
	"extended-design-golang/app/models"
	"extended-design-golang/app/services/pay"
	"extended-design-golang/app/util"
	"fmt"
)

func GetPayChannelInstance(channel models.PayInfo) pay.PayChannelInterface {

	fmt.Println(string(util.JsonMarshal(channel)))

	switch channel.PayMethod.ChannelCode {
	case "kuaijiepay":
		return new(PayChannelKuaiJiePay)
	case "xinbeipay":
		return new(PayChannelXinBeiPay)
	default:
		return nil
	}
}
