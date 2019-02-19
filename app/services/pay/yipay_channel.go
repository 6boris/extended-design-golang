package pay

type PayChannelYiPay struct {
	PayChannelInterface
}

func (self PayChannelYiPay) GetPayInstance() {
}

func (self PayChannelYiPay) GetPayInfo() string {
	return "yipay"
}
func (self PayChannelYiPay) GetPayCallback() string {
	return "yipay"
}
