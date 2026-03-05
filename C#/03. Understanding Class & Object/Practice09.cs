/**
 * Practice 09: Bank Account Collection
 *
 * This is an extended practice of Practice 04. In this exercise, you will
 * create a list of Bank Accounts, perform various transactions (such as
 * withdrawals and deposits) across multiple accounts, and calculate the
 * total balance of all the accounts owned by the Bank.
 *
 * Key Concepts:
 *   - Reusing the BankAccount class from Practice 04
 *   - List<T> of objects (object collection)
 *   - Iterating over a collection to aggregate data
 *   - Business operations across multiple objects
 *
 * Course: Professional OOP — by Zohirul Alam Tiemoon
 */

using System;
using System.Collections.Generic;

/** Represents a bank account with basic operations. */
class BankAccount
{
    private string accountNumber;
    private string accountName;
    private double balance;

    public BankAccount(string accountNumber, string accountName, double balance)
    {
        this.accountNumber = accountNumber;
        this.accountName = accountName;
        this.balance = balance;
    }

    public string AccountNumber { get { return accountNumber; } }
    public string AccountName { get { return accountName; } }
    public double Balance { get { return balance; } }

    public void Deposit(double amount)
    {
        if (amount <= 0)
        {
            Console.WriteLine("  [Error] Deposit amount must be positive.");
            return;
        }
        balance += amount;
        Console.WriteLine($"  [OK] Deposited {amount:F2} to {accountName}. New balance: {balance:F2}");
    }

    public void Withdraw(double amount)
    {
        if (amount <= 0)
        {
            Console.WriteLine("  [Error] Withdrawal amount must be positive.");
            return;
        }
        if (amount > balance)
        {
            Console.WriteLine($"  [Error] Insufficient balance in {accountName}. Available: {balance:F2}");
            return;
        }
        balance -= amount;
        Console.WriteLine($"  [OK] Withdrew {amount:F2} from {accountName}. New balance: {balance:F2}");
    }

    public void Transfer(BankAccount to, double amount)
    {
        if (amount <= 0)
        {
            Console.WriteLine("  [Error] Transfer amount must be positive.");
            return;
        }
        if (amount > balance)
        {
            Console.WriteLine($"  [Error] Insufficient balance in {accountName}. Available: {balance:F2}");
            return;
        }
        balance -= amount;
        to.balance += amount;
        Console.WriteLine($"  [OK] Transferred {amount:F2} from {accountName} to {to.accountName}");
    }

    public void PrintInfo()
    {
        Console.WriteLine($"  {accountNumber,-10} | {accountName,-10} | Balance: {balance,10:F2}");
    }
}

/** Manages a collection of bank accounts. */
class Bank
{
    private string bankName;
    private List<BankAccount> accounts;

    public Bank(string bankName)
    {
        this.bankName = bankName;
        this.accounts = new List<BankAccount>();
    }

    public void AddAccount(BankAccount account)
    {
        accounts.Add(account);
        Console.WriteLine($"  [OK] Added account {account.AccountNumber} ({account.AccountName}) to {bankName}.");
    }

    public void ShowAllAccounts()
    {
        Console.WriteLine($"  {bankName} — All Accounts ({accounts.Count}):");
        foreach (BankAccount acc in accounts)
        {
            acc.PrintInfo();
        }
    }

    public double GetTotalBalance()
    {
        double total = 0;
        foreach (BankAccount acc in accounts)
        {
            total += acc.Balance;
        }
        return total;
    }

    public void PrintTotalBalance()
    {
        Console.WriteLine($"  Total balance of {bankName}: {GetTotalBalance():F2}");
    }
}

class Practice09
{
    static void Main(string[] args)
    {
        Console.WriteLine("=== Practice 09: Bank Account Collection ===");
        Console.WriteLine();

        // Create bank
        Bank bank = new Bank("City Bank");

        // Create accounts
        BankAccount acc1 = new BankAccount("ACC-1001", "Imtiaz", 50000);
        BankAccount acc2 = new BankAccount("ACC-1002", "Faria", 30000);
        BankAccount acc3 = new BankAccount("ACC-1003", "Rafi", 45000);
        BankAccount acc4 = new BankAccount("ACC-1004", "Salma", 60000);

        // Add accounts to bank
        Console.WriteLine("--- Add Accounts ---");
        bank.AddAccount(acc1);
        bank.AddAccount(acc2);
        bank.AddAccount(acc3);
        bank.AddAccount(acc4);
        Console.WriteLine();

        // Show all accounts
        Console.WriteLine("--- All Accounts ---");
        bank.ShowAllAccounts();
        Console.WriteLine();

        // Total balance before transactions
        Console.WriteLine("--- Total Balance (Before) ---");
        bank.PrintTotalBalance();
        Console.WriteLine();

        // Perform transactions
        Console.WriteLine("--- Transactions ---");
        acc1.Deposit(10000);
        acc2.Withdraw(5000);
        acc3.Transfer(acc4, 15000);
        acc1.Transfer(acc2, 20000);
        Console.WriteLine();

        // Show all accounts after transactions
        Console.WriteLine("--- All Accounts (After Transactions) ---");
        bank.ShowAllAccounts();
        Console.WriteLine();

        // Total balance after transactions (should be same — money moves within bank)
        Console.WriteLine("--- Total Balance (After) ---");
        bank.PrintTotalBalance();
    }
}
