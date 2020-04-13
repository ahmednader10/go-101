package payment

import "fmt"

type Cash struct {
}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (c Cash) ExposedProcessPayment(value float32) bool {
	fmt.Println("processing received cash payment")
	return true
}
