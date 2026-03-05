/**
 * Practice 09: Bank Account Collection
 * Task: Manage a list of BankAccount objects. Perform transactions,
 *       calculate total balance, and find accounts with highest/lowest balance.
 *
 * How to run:
 *   go run practice-09.go
 *
 * Key Concepts:
 *   - Slice of structs (collection of objects)
 *   - Iterating over a collection to compute aggregates
 *   - Finding min/max in a collection of objects
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
func NewBankAccount(accountNumber, accountName string, balance float64) BankAccount {
	return BankAccount{
		AccountNumber: accountNumber,
		AccountName:   accountName,
		Balance:       balance,
	}
}

/** Deposit adds the given amount to the account balance. */
func (a *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Deposit amount must be greater than 0.")
		return
	}
	a.Balance += amount
	fmt.Printf("  [OK] Deposited %.2f to %s\n", amount, a.AccountName)
}

/** Withdraw subtracts the given amount from the account balance. */
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

/** ShowInfo prints the account details. */
func (a *BankAccount) ShowInfo() {
	fmt.Printf("  %-10s | %-12s | Balance: %10.2f\n", a.AccountNumber, a.AccountName, a.Balance)
}

/** calculateTotalBalance returns the sum of all account balances. */
func calculateTotalBalance(accounts []BankAccount) float64 {
	total := 0.0
	for _, acc := range accounts {
		total += acc.Balance
	}
	return total
}

/** findHighestBalance returns the account with the highest balance. */
func findHighestBalance(accounts []BankAccount) BankAccount {
	highest := accounts[0]
	for _, acc := range accounts[1:] {
		if acc.Balance > highest.Balance {
			highest = acc
		}
	}
	return highest
}

/** findLowestBalance returns the account with the lowest balance. */
func findLowestBalance(accounts []BankAccount) BankAccount {
	lowest := accounts[0]
	for _, acc := range accounts[1:] {
		if acc.Balance < lowest.Balance {
			lowest = acc
		}
	}
	return lowest
}

/** showAllAccounts prints all accounts in the list. */
func showAllAccounts(accounts []BankAccount) {
	fmt.Printf("  %-10s | %-12s | %s\n", "Account No", "Name", "Balance")
	fmt.Println("  " + "-------------------------------------------")
	for i := range accounts {
		accounts[i].ShowInfo()
	}
	fmt.Println()
}

func main() {
	// --- Create a list of bank accounts ---
	accounts := []BankAccount{
		NewBankAccount("ACC-1001", "Tareq", 15000),
		NewBankAccount("ACC-1002", "Afsana", 22000),
		NewBankAccount("ACC-1003", "Imtiaz", 8500),
		NewBankAccount("ACC-1004", "Pulok", 31000),
		NewBankAccount("ACC-1005", "Samia", 12000),
	}

	fmt.Println("=== All Accounts (Initial) ===")
	showAllAccounts(accounts)

	// --- Perform transactions ---
	fmt.Println("=== Performing Transactions ===")
	accounts[0].Deposit(5000)   // Tareq deposits 5000
	accounts[1].Withdraw(3000)  // Afsana withdraws 3000
	accounts[2].Deposit(1500)   // Imtiaz deposits 1500
	accounts[3].Withdraw(10000) // Pulok withdraws 10000
	accounts[4].Deposit(8000)   // Samia deposits 8000
	fmt.Println()

	fmt.Println("=== All Accounts (After Transactions) ===")
	showAllAccounts(accounts)

	// --- Calculate total balance ---
	total := calculateTotalBalance(accounts)
	fmt.Printf("  Total Balance of All Accounts: %.2f\n\n", total)

	// --- Find highest and lowest balance ---
	highest := findHighestBalance(accounts)
	lowest := findLowestBalance(accounts)

	fmt.Println("=== Highest Balance ===")
	highest.ShowInfo()
	fmt.Println()

	fmt.Println("=== Lowest Balance ===")
	lowest.ShowInfo()
}
