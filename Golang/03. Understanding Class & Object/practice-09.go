/**
 * Practice 09: Bank Account Collection
 *
 * This is an extended practice of Practice 04. In this exercise, you will
 * create a list of Bank Accounts, perform various transactions (such as
 * withdrawals and deposits) across multiple accounts, and calculate the
 * total balance of all the accounts owned by the Bank.
 *
 * Key Concepts:
 *   - Reusing the BankAccount struct from Practice 04
 *   - Slice of structs (object collection)
 *   - Iterating over a collection to aggregate data
 *   - Business operations across multiple objects
 *
 * Course: Professional OOP — by Zohirul Alam Tiemoon
 */

package main

import "fmt"

// BankAccount represents a bank account with basic operations.
type BankAccount struct {
	accountNumber string
	accountName   string
	balance       float64
}

/** NewBankAccount creates and returns a new BankAccount. */
func NewBankAccount(accountNumber, accountName string, balance float64) BankAccount {
	return BankAccount{
		accountNumber: accountNumber,
		accountName:   accountName,
		balance:       balance,
	}
}

/** Deposit adds the given amount to the balance. */
func (a *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Deposit amount must be positive.")
		return
	}
	a.balance += amount
	fmt.Printf("  [OK] Deposited %.2f to %s. New balance: %.2f\n",
		amount, a.accountName, a.balance)
}

/** Withdraw subtracts the given amount from the balance. */
func (a *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Withdrawal amount must be positive.")
		return
	}
	if amount > a.balance {
		fmt.Printf("  [Error] Insufficient balance in %s. Available: %.2f\n",
			a.accountName, a.balance)
		return
	}
	a.balance -= amount
	fmt.Printf("  [OK] Withdrew %.2f from %s. New balance: %.2f\n",
		amount, a.accountName, a.balance)
}

/** Transfer moves the given amount from this account to another. */
func (a *BankAccount) Transfer(to *BankAccount, amount float64) {
	if amount <= 0 {
		fmt.Println("  [Error] Transfer amount must be positive.")
		return
	}
	if amount > a.balance {
		fmt.Printf("  [Error] Insufficient balance in %s. Available: %.2f\n",
			a.accountName, a.balance)
		return
	}
	a.balance -= amount
	to.balance += amount
	fmt.Printf("  [OK] Transferred %.2f from %s to %s\n",
		amount, a.accountName, to.accountName)
}

/** PrintInfo displays account details. */
func (a *BankAccount) PrintInfo() {
	fmt.Printf("  %-10s | %-10s | Balance: %10.2f\n",
		a.accountNumber, a.accountName, a.balance)
}

// Bank manages a collection of accounts.
type Bank struct {
	bankName string
	accounts []*BankAccount
}

/** NewBank creates a bank with an empty account list. */
func NewBank(bankName string) Bank {
	return Bank{bankName: bankName, accounts: []*BankAccount{}}
}

/** AddAccount adds an account to the bank. */
func (b *Bank) AddAccount(account *BankAccount) {
	b.accounts = append(b.accounts, account)
	fmt.Printf("  [OK] Added account %s (%s) to %s.\n",
		account.accountNumber, account.accountName, b.bankName)
}

/** ShowAllAccounts lists all accounts in the bank. */
func (b *Bank) ShowAllAccounts() {
	fmt.Printf("  %s — All Accounts (%d):\n", b.bankName, len(b.accounts))
	for _, acc := range b.accounts {
		acc.PrintInfo()
	}
}

/** GetTotalBalance calculates and returns the sum of all account balances. */
func (b *Bank) GetTotalBalance() float64 {
	total := 0.0
	for _, acc := range b.accounts {
		total += acc.balance
	}
	return total
}

/** PrintTotalBalance displays the total balance across all accounts. */
func (b *Bank) PrintTotalBalance() {
	fmt.Printf("  Total balance of %s: %.2f\n", b.bankName, b.GetTotalBalance())
}

func main() {
	fmt.Println("=== Practice 09: Bank Account Collection ===")
	fmt.Println()

	// Create bank
	bank := NewBank("City Bank")

	// Create accounts
	acc1 := NewBankAccount("ACC-1001", "Imtiaz", 50000)
	acc2 := NewBankAccount("ACC-1002", "Faria", 30000)
	acc3 := NewBankAccount("ACC-1003", "Rafi", 45000)
	acc4 := NewBankAccount("ACC-1004", "Salma", 60000)

	// Add accounts to bank
	fmt.Println("--- Add Accounts ---")
	bank.AddAccount(&acc1)
	bank.AddAccount(&acc2)
	bank.AddAccount(&acc3)
	bank.AddAccount(&acc4)
	fmt.Println()

	// Show all accounts
	fmt.Println("--- All Accounts ---")
	bank.ShowAllAccounts()
	fmt.Println()

	// Total balance before transactions
	fmt.Println("--- Total Balance (Before) ---")
	bank.PrintTotalBalance()
	fmt.Println()

	// Perform transactions
	fmt.Println("--- Transactions ---")
	acc1.Deposit(10000)
	acc2.Withdraw(5000)
	acc3.Transfer(&acc4, 15000)
	acc1.Transfer(&acc2, 20000)
	fmt.Println()

	// Show all accounts after transactions
	fmt.Println("--- All Accounts (After Transactions) ---")
	bank.ShowAllAccounts()
	fmt.Println()

	// Total balance after transactions (should be same — money moves within bank)
	fmt.Println("--- Total Balance (After) ---")
	bank.PrintTotalBalance()
}
