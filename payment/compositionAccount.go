package payment

import "fmt"

type CompositionAccount struct {

}

func (a CompositionAccount) ProcessPayment(amount float32) bool {
	fmt.Println("processing payment", amount)
	return true
}

type CreditAmount struct {
	CompositionAccount
}