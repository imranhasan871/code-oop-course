/**
 * Practice 05: Credit Card
 * Task: Create a CreditCard class with spending limits.
 *       Max credit limit: 500,000. Cash withdrawal: daily limit 100,000,
 *       per-transaction limit 20,000. Bill payments: no per-txn limit.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice05.cs
 *
 * Or create a Console App and copy contents into Program.cs:
 *   dotnet new console -n Practice05
 *   dotnet run --project Practice05
 *
 * Key Concepts:
 *   - Class with constants and instance tracking
 *   - Multiple validation rules in methods
 *   - Business logic enforcement
 */

using System;

class Practice05
{
    /** CreditCard represents a credit card with spending limits. */
    class CreditCard
    {
        private const double MaxCreditLimit = 500000.00;
        private const double DailyCashLimit = 100000.00;
        private const double PerTxnCashLimit = 20000.00;

        private string cardNumber;
        private string cardHolder;
        private double totalSpent;
        private double dailyCashWithdrawn;

        /** Creates a new CreditCard with the given details. */
        public CreditCard(string cardNumber, string cardHolder)
        {
            this.cardNumber = cardNumber;
            this.cardHolder = cardHolder;
            this.totalSpent = 0;
            this.dailyCashWithdrawn = 0;
        }

        /** Withdraws cash with per-transaction and daily limits. */
        public void WithdrawCash(double amount)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Amount must be greater than 0.");
                return;
            }
            if (amount > PerTxnCashLimit)
            {
                Console.WriteLine($"  [Error] Cash withdrawal exceeds per-transaction limit of {PerTxnCashLimit:F2} (Requested: {amount:F2})");
                return;
            }
            if (dailyCashWithdrawn + amount > DailyCashLimit)
            {
                Console.WriteLine($"  [Error] Cash withdrawal exceeds daily limit of {DailyCashLimit:F2} (Already withdrawn today: {dailyCashWithdrawn:F2}, Requested: {amount:F2})");
                return;
            }
            if (totalSpent + amount > MaxCreditLimit)
            {
                Console.WriteLine($"  [Error] Total spending would exceed credit limit of {MaxCreditLimit:F2} (Already spent: {totalSpent:F2}, Requested: {amount:F2})");
                return;
            }
            totalSpent += amount;
            dailyCashWithdrawn += amount;
            Console.WriteLine($"  [OK] Cash withdrawn: {amount:F2}");
        }

        /** Pays a bill amount. No per-transaction limit, but total must stay within max limit. */
        public void PayBill(double amount)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Amount must be greater than 0.");
                return;
            }
            if (totalSpent + amount > MaxCreditLimit)
            {
                Console.WriteLine($"  [Error] Bill payment would exceed credit limit of {MaxCreditLimit:F2} (Already spent: {totalSpent:F2}, Requested: {amount:F2})");
                return;
            }
            totalSpent += amount;
            Console.WriteLine($"  [OK] Bill paid: {amount:F2}");
        }

        /** Prints the card details and current spending. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Card Number          : {cardNumber}");
            Console.WriteLine($"  Card Holder          : {cardHolder}");
            Console.WriteLine($"  Credit Limit         : {MaxCreditLimit:F2}");
            Console.WriteLine($"  Total Spent          : {totalSpent:F2}");
            Console.WriteLine($"  Available Limit      : {MaxCreditLimit - totalSpent:F2}");
            Console.WriteLine($"  Daily Cash Withdrawn : {dailyCashWithdrawn:F2} / {DailyCashLimit:F2}");
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        // --- Create a credit card ---
        CreditCard card = new CreditCard("4532-1234-5678-9012", "Tareq");

        Console.WriteLine("=== Initial Card Info ===");
        card.ShowInfo();

        // --- Valid cash withdrawal ---
        Console.WriteLine("=== Withdraw Cash 15,000 ===");
        card.WithdrawCash(15000);
        card.ShowInfo();

        // --- Exceed per-transaction limit ---
        Console.WriteLine("=== Withdraw Cash 25,000 (exceeds per-txn limit) ===");
        card.WithdrawCash(25000);
        card.ShowInfo();

        // --- Multiple valid withdrawals to approach daily limit ---
        Console.WriteLine("=== Withdraw Cash 20,000 x4 (should hit daily limit) ===");
        card.WithdrawCash(20000);
        card.WithdrawCash(20000);
        card.WithdrawCash(20000);
        card.WithdrawCash(20000);
        card.ShowInfo();

        // --- Bill payment ---
        Console.WriteLine("=== Pay Bill 200,000 ===");
        card.PayBill(200000);
        card.ShowInfo();

        // --- Try to exceed total credit limit ---
        Console.WriteLine("=== Pay Bill 300,000 (should exceed total limit) ===");
        card.PayBill(300000);
        card.ShowInfo();
    }
}
