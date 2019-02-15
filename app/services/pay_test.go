package services

import (
	"extended-design-golang/app/util"
	"fmt"
	"testing"
)

func TestGetPayMap(t *testing.T) {
	Setup()

	maps1 := GetPayMethodMap()

	fmt.Println(string(util.JsonMarshal(maps1)))

}
func TestGetAllPay(t *testing.T) {
	Setup()

	data := GetAllPay()

	fmt.Println(string(util.JsonMarshal(data)))
}
