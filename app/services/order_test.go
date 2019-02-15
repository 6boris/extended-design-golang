package services

import (
	"fmt"
	"testing"
)

func TestGetAllOrders(t *testing.T) {
	Setup()
	orders := GetAllOrders()
	fmt.Println(orders)
}
