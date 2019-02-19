package pay

type PayChannelXinBeiPay struct {
	PayChannelInterface
}

func (self PayChannelXinBeiPay) GetPayInstance() {
}

func (self PayChannelXinBeiPay) GetPayInfo() string {
	return "xinbei"
}
func (self PayChannelXinBeiPay) GetPayCallback() string {
	return "xinbei"
}
