package services

import (
	"extended-design-golang/app/util"
	"fmt"
	"testing"
)

func TestGetAllPayTypes(t *testing.T) {
	Setup()

	pay_types := GetAllPayTypes()
	type_json := util.JsonMarshal(pay_types)

	fmt.Println(pay_types)
	fmt.Println(string(type_json))

}
