/**
 * Practice 09: Bank Account Collection
 * Task: Manage a list of BankAccount objects. Perform transactions,
 *       calculate total balance, and find accounts with highest/lowest balance.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice09.cs
 *
 * Or create a Console App and copy contents into Program.cs:
 *   dotnet new console -n Practice09
 *   dotnet run --project Practice09
 *
 * Key Concepts:
 *   - List of objects (collection of objects)
 *   - Iterating over a collection to compute aggregates
 *   - Finding min/max in a collection of objects
 */

using System;
using System.Collections.Generic;

class Practice09
{
    /** BankAccount represents a bank account with basic operations. */
    class BankAccount
    {
        public string AccountNumber;
        public string AccountName;
        public double Balance;

        /** Creates a new BankAccount with the given details. */
        public BankAccount(string accountNumber, string accountName, double balance)
        {
            AccountNumber = accountNumber;
            AccountName = accountName;
            Balance = balance;
        }

        /** Deposits the given amount into the account. */
        public void Deposit(double amount)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Deposit amount must be greater than 0.");
                return;
            }
            Balance += amount;
            Console.WriteLine($"  [OK] Deposited {amount:F2} to {AccountName}");
        }

        /** Withdraws the given amount from the account. */
        public void Withdraw(double amount)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Withdrawal amount must be greater than 0.");
                return;
            }
            if (Balance < amount)
            {
                Console.WriteLine($"  [Error] Insufficient balance in {AccountName} (Balance: {Balance:F2}, Requested: {amount:F2})");
                return;
            }
            Balance -= amount;
            Console.WriteLine($"  [OK] Withdrew {amount:F2} from {AccountName}");
        }

        /** Prints the account details in a table row format. */
        public void ShowInfo()
        {
            Console.WriteLine($"  {AccountNumber,-10} | {AccountName,-12} | Balance: {Balance,10:F2}");
        }
    }

    /** Calculates the total balance of all accounts. */
    static double CalculateTotalBalance(List<BankAccount> accounts)
    {
        double total = 0;
        foreach (BankAccount acc in accounts)
        {
            total += acc.Balance;
        }
        return total;
    }

    /** Finds the account with the highest balance. */
    static BankAccount FindHighestBalance(List<BankAccount> accounts)
    {
        BankAccount highest = accounts[0];
        for (int i = 1; i < accounts.Count; i++)
        {
            if (accounts[i].Balance > highest.Balance)
                highest = accounts[i];
        }
        return highest;
    }

    /** Finds the account with the lowest balance. */
    static BankAccount FindLowestBalance(List<BankAccount> accounts)
    {
        BankAccount lowest = accounts[0];
        for (int i = 1; i < accounts.Count; i++)
        {
            if (accounts[i].Balance < lowest.Balance)
                lowest = accounts[i];
        }
        return lowest;
    }

    /** Prints all accounts in a formatted table. */
    static void ShowAllAccounts(List<BankAccount> accounts)
    {
        Console.WriteLine($"  {"Account No",-10} | {"Name",-12} | {"Balance"}");
        Console.WriteLine("  -------------------------------------------");
        foreach (BankAccount acc in accounts)
        {
            acc.ShowInfo();
        }
        Console.WriteLine();
    }

    static void Main(string[] args)
    {
        // --- Create a list of bank accounts ---
        List<BankAccount> accounts = new List<BankAccount>
        {
            new BankAccount("ACC-1001", "Tareq", 15000),
            new BankAccount("ACC-1002", "Afsana", 22000),
            new BankAccount("ACC-1003", "Imtiaz", 8500),
            new BankAccount("ACC-1004", "Pulok", 31000),
            new BankAccount("ACC-1005", "Samia", 12000)
        };

        Console.WriteLine("=== All Accounts (Initial) ===");
        ShowAllAccounts(accounts);

        // --- Perform transactions ---
        Console.WriteLine("=== Performing Transactions ===");
        accounts[0].Deposit(5000);   // Tareq deposits 5000
        accounts[1].Withdraw(3000);  // Afsana withdraws 3000
        accounts[2].Deposit(1500);   // Imtiaz deposits 1500
        accounts[3].Withdraw(10000); // Pulok withdraws 10000
        accounts[4].Deposit(8000);   // Samia deposits 8000
        Console.WriteLine();

        Console.WriteLine("=== All Accounts (After Transactions) ===");
        ShowAllAccounts(accounts);

        // --- Calculate total balance ---
        double total = CalculateTotalBalance(accounts);
        Console.WriteLine($"  Total Balance of All Accounts: {total:F2}");
        Console.WriteLine();

        // --- Find highest and lowest balance ---
        BankAccount highest = FindHighestBalance(accounts);
        BankAccount lowest = FindLowestBalance(accounts);

        Console.WriteLine("=== Highest Balance ===");
        highest.ShowInfo();
        Console.WriteLine();

        Console.WriteLine("=== Lowest Balance ===");
        lowest.ShowInfo();
    }
}
