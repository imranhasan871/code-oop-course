/**
 * Practice 04: Bank Account
 * Task: Create a BankAccount struct with deposit, withdraw, and transfer methods.
 *       No negative balance is allowed.
 *
 * How to run:
 *   go run practice-04.go
 *
 * Key Concepts:
 *   - Struct as a class equivalent in Go
 *   - Methods with pointer receivers to modify struct fields
 *   - Input validation (no negative balance)
 */

package main

import "fmt"

/** BankAccount represents a bank account with basic operations. */
type BankAccount struct {
	AccountNumber string
	AccountName   string
	Balance       float64
}

/** NewBankAccount creates a new BankAccount with the given details. */
func NewBankAccount(accountNumber string, accountName string, balance float64) BankAccount {
	return BankAccount{
		AccountNumber: accountNumber,
		AccountName:   accountName,
		Balance:       balance,
	}
}

/** Deposit adds the given amount to the account balance. Amount must be positive. */
func (a *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Deposit amount must be greater than 0.")
		return
	}
	a.Balance += amount
	fmt.Printf("  [OK] Deposited %.2f to %s\n", amount, a.AccountName)
}

/** Withdraw subtracts the given amount from the account balance. No negative balance allowed. */
func (a *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Withdrawal amount must be greater than 0.")
		return
	}
	if a.Balance < amount {
		fmt.Printf("  [Error] Insufficient balance in %s (Balance: %.2f, Requested: %.2f)\n",
			a.AccountName, a.Balance, amount)
		return
	}
	a.Balance -= amount
	fmt.Printf("  [OK] Withdrew %.2f from %s\n", amount, a.AccountName)
}

/** Transfer moves the given amount from this account to the target account. */
func (a *BankAccount) Transfer(amount float64, target *BankAccount) {
	if amount <= 0 {
		fmt.Println("  [Error] Transfer amount must be greater than 0.")
		return
	}
	if a.Balance < amount {
		fmt.Printf("  [Error] Insufficient balance in %s for transfer (Balance: %.2f, Requested: %.2f)\n",
			a.AccountName, a.Balance, amount)
		return
	}
	a.Balance -= amount
	target.Balance += amount
	fmt.Printf("  [OK] Transferred %.2f from %s to %s\n", amount, a.AccountName, target.AccountName)
}

/** ShowInfo prints the account number, name, and current balance. */
func (a *BankAccount) ShowInfo() {
	fmt.Printf("  Account Number : %s\n", a.AccountNumber)
	fmt.Printf("  Account Name   : %s\n", a.AccountName)
	fmt.Printf("  Balance        : %.2f\n", a.Balance)
	fmt.Println()
}

func main() {
	// --- Create two bank accounts ---
	account1 := NewBankAccount("ACC-1001", "Tareq", 5000.00)
	account2 := NewBankAccount("ACC-1002", "Afsana", 3000.00)

	fmt.Println("=== Initial Account Info ===")
	account1.ShowInfo()
	account2.ShowInfo()

	// --- Deposit ---
	fmt.Println("=== Deposit 2000 to Tareq ===")
	account1.Deposit(2000)
	account1.ShowInfo()

	// --- Withdraw ---
	fmt.Println("=== Withdraw 1500 from Afsana ===")
	account2.Withdraw(1500)
	account2.ShowInfo()

	// --- Withdraw more than balance (should fail) ---
	fmt.Println("=== Withdraw 5000 from Afsana (should fail) ===")
	account2.Withdraw(5000)
	account2.ShowInfo()

	// --- Transfer ---
	fmt.Println("=== Transfer 3000 from Tareq to Afsana ===")
	account1.Transfer(3000, &account2)
	fmt.Println("After transfer:")
	account1.ShowInfo()
	account2.ShowInfo()
}
