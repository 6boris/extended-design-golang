package services

import (
	"extended-design-golang/app/util"
	"fmt"
	"testing"
)

func TestGetAllPayChannels(t *testing.T) {
	Setup()

	pay_channels := GetAllPayChannels()
	json_channels := util.JsonMarshal(pay_channels)
	fmt.Println(string(json_channels))

}
