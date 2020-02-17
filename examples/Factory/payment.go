package creational

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash = iota
	DebitCard
	Other
)

type Payment struct {
	paymentType int
}

func NewPayment(paymentType int) Payment {
	return Payment{paymentType: paymentType}
}

func (p Payment) GetPaymentMethod() (PaymentMethod, error) {

	switch p.paymentType {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method: %d not recognized\n", p.paymentType))
	}
}
