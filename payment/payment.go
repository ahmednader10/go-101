package payment

import (
	"errors"
	"fmt"
)

type PaymentOption interface {
	ExposedProcessPayment(float32) bool
}

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

func CreateCreditAccount(ownerName string, cardNumber string, expirationMonth int, expirationYear int, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

func CreateCreditAccountWithChannel(chargeChan chan float32, ownerName string, cardNumber string, expirationMonth int, expirationYear int, securityCode int) *CreditCard {
	creditCard := &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
	go func(chargeChan chan float32) {
		for amount := range chargeChan {
			creditCard.internalProcessPayment(amount)
		}
	}(chargeChan)

	return creditCard
}

func (c CreditCard) internalProcessPayment(value float32) bool {
	fmt.Println("processing received credit card payment", value)
	return true
}

func (c *CreditCard) ExposedProcessPayment(value float32) bool {
	fmt.Println("processing received credit card payment")
	return true
}

func (c *CreditCard) OwnerName() string {
	return c.ownerName
}

func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid name provided")
	}
	c.ownerName = value
	return nil
}

func (c *CreditCard) CardNumber() string {
	return c.cardNumber
}

func (c *CreditCard) ExpirationData() (int, int) {
	return c.expirationMonth, c.expirationYear
}

func (c *CreditCard) SecurityCode() int {
	return c.securityCode
}

func (c *CreditCard) SetSecurityCode(value int) error {
	if value < 100 || value > 999 {
		return errors.New("invalid value")
	}
	c.securityCode = value
	return nil
}

func (c *CreditCard) AvailableCredit() float32 {
	return 5000
}
