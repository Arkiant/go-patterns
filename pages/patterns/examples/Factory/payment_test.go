package creational

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment := NewPayment(CASH)
	paymentMethod, err := payment.GetPaymentMethod()
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exists")
	}

	msg := paymentMethod.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment := NewPayment(DEBITCARD)
	paymentMethod, err := payment.GetPaymentMethod()
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exists")
	}

	msg := paymentMethod.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	payment := NewPayment(OTHER)
	_, err := payment.GetPaymentMethod()
	if err == nil {
		t.Errorf("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
