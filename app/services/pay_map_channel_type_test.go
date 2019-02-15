package services

import (
	"extended-design-golang/app/util"
	"fmt"
	"testing"
)

func TestGetAllPayMapChannelType(t *testing.T) {
	Setup()

	pay_map_channel_type := GetAllPayMapChannelType()
	json_data := util.JsonMarshal(pay_map_channel_type)

	fmt.Println(string(json_data))
}
