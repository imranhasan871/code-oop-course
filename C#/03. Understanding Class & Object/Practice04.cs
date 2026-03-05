/**
 * Practice 04: Bank Account
 * Task: Create a BankAccount class with deposit, withdraw, and transfer methods.
 *       No negative balance is allowed.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice04.cs
 *
 * Or create a Console App and copy contents into Program.cs:
 *   dotnet new console -n Practice04
 *   dotnet run --project Practice04
 *
 * Key Concepts:
 *   - Class with private fields and public methods (encapsulation)
 *   - Constructor for initializing object state
 *   - Input validation in methods
 */

using System;

class Practice04
{
    /** BankAccount represents a bank account with basic operations. */
    class BankAccount
    {
        private string accountNumber;
        private string accountName;
        private double balance;

        /** Creates a new BankAccount with the given details. */
        public BankAccount(string accountNumber, string accountName, double balance)
        {
            this.accountNumber = accountNumber;
            this.accountName = accountName;
            this.balance = balance;
        }

        /** Deposits the given amount into the account. Amount must be positive. */
        public void Deposit(double amount)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Deposit amount must be greater than 0.");
                return;
            }
            balance += amount;
            Console.WriteLine($"  [OK] Deposited {amount:F2} to {accountName}");
        }

        /** Withdraws the given amount from the account. No negative balance allowed. */
        public void Withdraw(double amount)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Withdrawal amount must be greater than 0.");
                return;
            }
            if (balance < amount)
            {
                Console.WriteLine($"  [Error] Insufficient balance in {accountName} (Balance: {balance:F2}, Requested: {amount:F2})");
                return;
            }
            balance -= amount;
            Console.WriteLine($"  [OK] Withdrew {amount:F2} from {accountName}");
        }

        /** Transfers the given amount from this account to the target account. */
        public void Transfer(double amount, BankAccount target)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Transfer amount must be greater than 0.");
                return;
            }
            if (balance < amount)
            {
                Console.WriteLine($"  [Error] Insufficient balance in {accountName} for transfer (Balance: {balance:F2}, Requested: {amount:F2})");
                return;
            }
            balance -= amount;
            target.balance += amount;
            Console.WriteLine($"  [OK] Transferred {amount:F2} from {accountName} to {target.accountName}");
        }

        /** Prints the account number, name, and current balance. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Account Number : {accountNumber}");
            Console.WriteLine($"  Account Name   : {accountName}");
            Console.WriteLine($"  Balance        : {balance:F2}");
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        // --- Create two bank accounts ---
        BankAccount account1 = new BankAccount("ACC-1001", "Tareq", 5000.00);
        BankAccount account2 = new BankAccount("ACC-1002", "Afsana", 3000.00);

        Console.WriteLine("=== Initial Account Info ===");
        account1.ShowInfo();
        account2.ShowInfo();

        // --- Deposit ---
        Console.WriteLine("=== Deposit 2000 to Tareq ===");
        account1.Deposit(2000);
        account1.ShowInfo();

        // --- Withdraw ---
        Console.WriteLine("=== Withdraw 1500 from Afsana ===");
        account2.Withdraw(1500);
        account2.ShowInfo();

        // --- Withdraw more than balance (should fail) ---
        Console.WriteLine("=== Withdraw 5000 from Afsana (should fail) ===");
        account2.Withdraw(5000);
        account2.ShowInfo();

        // --- Transfer ---
        Console.WriteLine("=== Transfer 3000 from Tareq to Afsana ===");
        account1.Transfer(3000, account2);
        Console.WriteLine("After transfer:");
        account1.ShowInfo();
        account2.ShowInfo();
    }
}
