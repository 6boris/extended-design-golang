package pay

type PayChannelKuaiJiePay struct {
	PayChannelInterface
}

func (self PayChannelKuaiJiePay) GetPayInstance() {
}

func (self PayChannelKuaiJiePay) GetPayInfo() string {
	return "kuaijiepay"

}
func (self PayChannelKuaiJiePay) GetPayCallback() string {
	return "kuaijiepay"
}
