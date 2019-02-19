package pay_channel

import (
	"extended-design-golang/app/services/pay"
)

type PayChannelKuaiJiePay struct {
	pay.PayChannelInterface
}

func (self PayChannelKuaiJiePay) GetPayInstance() {
}

func (self PayChannelKuaiJiePay) GetPayInfo() string {
	return "kuaijiepay"

}
func (self PayChannelKuaiJiePay) GetPayCallback() string {
	return "kuaijiepay"
}
