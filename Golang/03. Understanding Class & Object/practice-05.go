/**
 * Practice 05: Credit Card
 * Task: Create a CreditCard struct with spending limits.
 *       Max credit limit: 500,000. Cash withdrawal: daily limit 100,000,
 *       per-transaction limit 20,000. Bill payments: no per-txn limit.
 *
 * How to run:
 *   go run practice-05.go
 *
 * Key Concepts:
 *   - Struct with multiple validation rules
 *   - Methods with pointer receivers
 *   - Business logic enforcement
 */

package main

import "fmt"

/** CreditCard represents a credit card with spending limits. */
type CreditCard struct {
	CardNumber         string
	CardHolder         string
	MaxLimit           float64
	TotalSpent         float64
	DailyCashWithdrawn float64
}

const (
	maxCreditLimit    = 500000.00
	dailyCashLimit    = 100000.00
	perTxnCashLimit   = 20000.00
)

/** NewCreditCard creates a new CreditCard with the given details. */
func NewCreditCard(cardNumber string, cardHolder string) CreditCard {
	return CreditCard{
		CardNumber: cardNumber,
		CardHolder: cardHolder,
		MaxLimit:   maxCreditLimit,
	}
}

/** WithdrawCash withdraws cash with per-transaction and daily limits. */
func (c *CreditCard) WithdrawCash(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Amount must be greater than 0.")
		return
	}
	if amount > perTxnCashLimit {
		fmt.Printf("  [Error] Cash withdrawal exceeds per-transaction limit of %.2f (Requested: %.2f)\n",
			perTxnCashLimit, amount)
		return
	}
	if c.DailyCashWithdrawn+amount > dailyCashLimit {
		fmt.Printf("  [Error] Cash withdrawal exceeds daily limit of %.2f (Already withdrawn today: %.2f, Requested: %.2f)\n",
			dailyCashLimit, c.DailyCashWithdrawn, amount)
		return
	}
	if c.TotalSpent+amount > c.MaxLimit {
		fmt.Printf("  [Error] Total spending would exceed credit limit of %.2f (Already spent: %.2f, Requested: %.2f)\n",
			c.MaxLimit, c.TotalSpent, amount)
		return
	}
	c.TotalSpent += amount
	c.DailyCashWithdrawn += amount
	fmt.Printf("  [OK] Cash withdrawn: %.2f\n", amount)
}

/** PayBill pays a bill amount. No per-transaction limit, but total spending must stay within max limit. */
func (c *CreditCard) PayBill(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Amount must be greater than 0.")
		return
	}
	if c.TotalSpent+amount > c.MaxLimit {
		fmt.Printf("  [Error] Bill payment would exceed credit limit of %.2f (Already spent: %.2f, Requested: %.2f)\n",
			c.MaxLimit, c.TotalSpent, amount)
		return
	}
	c.TotalSpent += amount
	fmt.Printf("  [OK] Bill paid: %.2f\n", amount)
}

/** ShowInfo prints the card details and current spending. */
func (c *CreditCard) ShowInfo() {
	fmt.Printf("  Card Number          : %s\n", c.CardNumber)
	fmt.Printf("  Card Holder          : %s\n", c.CardHolder)
	fmt.Printf("  Credit Limit         : %.2f\n", c.MaxLimit)
	fmt.Printf("  Total Spent          : %.2f\n", c.TotalSpent)
	fmt.Printf("  Available Limit      : %.2f\n", c.MaxLimit-c.TotalSpent)
	fmt.Printf("  Daily Cash Withdrawn : %.2f / %.2f\n", c.DailyCashWithdrawn, dailyCashLimit)
	fmt.Println()
}

func main() {
	// --- Create a credit card ---
	card := NewCreditCard("4532-1234-5678-9012", "Tareq")

	fmt.Println("=== Initial Card Info ===")
	card.ShowInfo()

	// --- Valid cash withdrawal ---
	fmt.Println("=== Withdraw Cash 15,000 ===")
	card.WithdrawCash(15000)
	card.ShowInfo()

	// --- Exceed per-transaction limit ---
	fmt.Println("=== Withdraw Cash 25,000 (exceeds per-txn limit) ===")
	card.WithdrawCash(25000)
	card.ShowInfo()

	// --- Multiple valid withdrawals to approach daily limit ---
	fmt.Println("=== Withdraw Cash 20,000 x4 (should hit daily limit) ===")
	card.WithdrawCash(20000)
	card.WithdrawCash(20000)
	card.WithdrawCash(20000)
	card.WithdrawCash(20000)
	card.ShowInfo()

	// --- Bill payment ---
	fmt.Println("=== Pay Bill 200,000 ===")
	card.PayBill(200000)
	card.ShowInfo()

	// --- Try to exceed total credit limit ---
	fmt.Println("=== Pay Bill 300,000 (should exceed total limit) ===")
	card.PayBill(300000)
	card.ShowInfo()
}
