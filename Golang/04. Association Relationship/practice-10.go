/**
 * Practice 10: 1-1 Association — Customer & Credit Card
 * Task: Manage customer profiles with associated credit cards.
 *       Card number must be digits only. Expiration date validation.
 *       Credit limit 500K. Track available credit. Age minimum 18.
 *
 * How to run:
 *   go run practice-10.go
 *
 * Key Concepts:
 *   - 1-1 Association (Customer has a CreditCard)
 *   - Input validation (digit-only card number, age check)
 *   - Date-based expiration validation
 */

package main

import (
	"fmt"
	"time"
	"unicode"
)

/** CreditCard with card number, expiration date, and credit limit. */
type CreditCard10 struct {
	CardNumber     string
	ExpirationDate time.Time
	CreditLimit    float64
	TotalSpent     float64
}

const maxCreditLimit10 = 500000.00

/** NewCreditCard creates a new CreditCard. Card number must contain only digits. */
func NewCreditCard10(cardNumber string, expirationDate time.Time, creditLimit float64) (*CreditCard10, error) {
	for _, c := range cardNumber {
		if !unicode.IsDigit(c) {
			return nil, fmt.Errorf("card number must contain only digits, got: '%s'", cardNumber)
		}
	}
	if creditLimit > maxCreditLimit10 {
		return nil, fmt.Errorf("credit limit cannot exceed %.2f", maxCreditLimit10)
	}
	return &CreditCard10{
		CardNumber:     cardNumber,
		ExpirationDate: expirationDate,
		CreditLimit:    creditLimit,
	}, nil
}

func (c *CreditCard10) IsValid() bool {
	return !time.Now().After(c.ExpirationDate)
}

func (c *CreditCard10) AvailableCredit() float64 {
	return c.CreditLimit - c.TotalSpent
}

func (c *CreditCard10) OutstandingBalance() float64 {
	return c.TotalSpent
}

func (c *CreditCard10) MakePurchase(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Purchase amount must be greater than 0.")
		return
	}
	if !c.IsValid() {
		fmt.Println("  [Error] Credit card has expired. Cannot make purchase.")
		return
	}
	if amount > c.AvailableCredit() {
		fmt.Printf("  [Error] Purchase of %.2f exceeds available credit (%.2f).\n",
			amount, c.AvailableCredit())
		return
	}
	c.TotalSpent += amount
	fmt.Printf("  [OK] Purchase of %.2f successful.\n", amount)
}

func (c *CreditCard10) ShowInfo() {
	status := "Valid"
	if !c.IsValid() {
		status = "Expired"
	}
	fmt.Println("  Card Number        :", c.CardNumber)
	fmt.Println("  Expiration Date    :", c.ExpirationDate.Format("2006-01-02"))
	fmt.Println("  Status             :", status)
	fmt.Printf("  Credit Limit       : %.2f\n", c.CreditLimit)
	fmt.Printf("  Total Spent        : %.2f\n", c.TotalSpent)
	fmt.Printf("  Available Credit   : %.2f\n", c.AvailableCredit())
	fmt.Printf("  Outstanding Balance: %.2f\n", c.OutstandingBalance())
}

/** Customer with personal info and an associated CreditCard (1-1). */
type Customer10 struct {
	Name        string
	DateOfBirth time.Time
	CreditCard  *CreditCard10
}

const minAge10 = 18

func calculateAge10(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

/** NewCustomer creates a Customer. Must be at least 18 years old. */
func NewCustomer10(name string, dob time.Time, card *CreditCard10) (*Customer10, error) {
	age := calculateAge10(dob)
	if age < minAge10 {
		return nil, fmt.Errorf("customer must be at least %d years old. %s is %d years old",
			minAge10, name, age)
	}
	return &Customer10{Name: name, DateOfBirth: dob, CreditCard: card}, nil
}

func (c *Customer10) GetAge() int {
	return calculateAge10(c.DateOfBirth)
}

func (c *Customer10) MakePurchase(amount float64) {
	fmt.Printf("  %s making purchase of %.2f...\n", c.Name, amount)
	c.CreditCard.MakePurchase(amount)
}

func (c *Customer10) ShowInfo() {
	fmt.Println("  Customer Name      :", c.Name)
	fmt.Println("  Date of Birth      :", c.DateOfBirth.Format("2006-01-02"))
	fmt.Println("  Age                :", c.GetAge())
	c.CreditCard.ShowInfo()
	fmt.Println()
}

func main() {
	// --- Create credit cards ---
	card1, _ := NewCreditCard10("4532123456789012", time.Date(2028, 12, 31, 0, 0, 0, 0, time.Local), maxCreditLimit10)
	card2, _ := NewCreditCard10("5678901234567890", time.Date(2025, 6, 15, 0, 0, 0, 0, time.Local), maxCreditLimit10)

	// --- Create customers ---
	customer1, _ := NewCustomer10("Tareq", time.Date(1990, 5, 15, 0, 0, 0, 0, time.Local), card1)
	customer2, _ := NewCustomer10("Afsana", time.Date(1995, 8, 20, 0, 0, 0, 0, time.Local), card2)

	fmt.Println("=== Customer Info ===")
	customer1.ShowInfo()
	customer2.ShowInfo()

	// --- Make purchases ---
	fmt.Println("=== Purchases ===")
	customer1.MakePurchase(150000)
	customer1.MakePurchase(200000)
	customer1.MakePurchase(200000) // Should fail
	fmt.Println()

	customer2.MakePurchase(50000) // Should fail — expired
	fmt.Println()

	fmt.Println("=== After Purchases ===")
	customer1.ShowInfo()

	// --- Invalid card number ---
	fmt.Println("=== Invalid Card Number ===")
	_, err := NewCreditCard10("ABCD-1234", time.Date(2028, 1, 1, 0, 0, 0, 0, time.Local), maxCreditLimit10)
	if err != nil {
		fmt.Println("  [Error]", err)
	}
	fmt.Println()

	// --- Underage customer ---
	fmt.Println("=== Underage Customer ===")
	youngCard, _ := NewCreditCard10("1111222233334444", time.Date(2030, 1, 1, 0, 0, 0, 0, time.Local), maxCreditLimit10)
	_, err = NewCustomer10("Junior", time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local), youngCard)
	if err != nil {
		fmt.Println("  [Error]", err)
	}
}
