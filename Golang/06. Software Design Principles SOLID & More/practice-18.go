/**
 * Practice 18: Applying DIP (Dependency Inversion Principle)
 * Task: Build a payment system where PaymentService depends on abstraction,
 *       not concrete payment providers.
 *
 * How to run:
 *   go run practice-18.go
 */

package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCardPayment struct{}

func (p *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card", amount)
}

type BkashPayment struct{}

func (p *BkashPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using bKash", amount)
}

type PaypalPayment struct{}

func (p *PaypalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
}

type BankTransferPayment struct{}

func (p *BankTransferPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Bank Transfer", amount)
}

type PaymentService struct{}

func (s *PaymentService) ProcessPayment(paymentMethod PaymentMethod, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	fmt.Println("[PaymentService] " + paymentMethod.Pay(amount))
	return nil
}

func main() {
	service := &PaymentService{}
	methods := []PaymentMethod{
		&CreditCardPayment{},
		&BkashPayment{},
		&PaypalPayment{},
		&BankTransferPayment{},
	}

	for _, method := range methods {
		err := service.ProcessPayment(method, 1500.0)
		if err != nil {
			panic(err)
		}
	}
}
