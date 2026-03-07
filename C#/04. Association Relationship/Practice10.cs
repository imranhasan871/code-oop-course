/**
 * Practice 10: 1-1 Association — Customer & Credit Card
 * Task: Manage customer profiles with associated credit cards.
 *       Card number must be digits only. Expiration date validation.
 *       Credit limit 500K. Track available credit. Age minimum 18.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice10.cs
 *
 * Key Concepts:
 *   - 1-1 Association (Customer has a CreditCard)
 *   - Input validation (digit-only card number, age check)
 *   - Date-based expiration validation
 */

using System;
using System.Linq;

class Practice10
{
    /** CreditCard with card number, expiration date, and credit limit. */
    class CreditCard
    {
        private const double MaxCreditLimit = 500000.00;

        private string cardNumber;
        private DateTime expirationDate;
        private double creditLimit;
        private double totalSpent;

        /** Creates a new CreditCard. Card number must contain only digits. */
        public CreditCard(string cardNumber, DateTime expirationDate, double creditLimit = 500000.00)
        {
            if (!cardNumber.All(char.IsDigit))
                throw new ArgumentException($"Card number must contain only digits, got: '{cardNumber}'");
            if (creditLimit > MaxCreditLimit)
                throw new ArgumentException($"Credit limit cannot exceed {MaxCreditLimit:F2}");
            this.cardNumber = cardNumber;
            this.expirationDate = expirationDate;
            this.creditLimit = creditLimit;
            this.totalSpent = 0;
        }

        public bool IsValid() => DateTime.Today <= expirationDate;
        public double AvailableCredit() => creditLimit - totalSpent;
        public double OutstandingBalance() => totalSpent;

        public void MakePurchase(double amount)
        {
            if (amount <= 0)
            {
                Console.WriteLine("  [Error] Purchase amount must be greater than 0.");
                return;
            }
            if (!IsValid())
            {
                Console.WriteLine("  [Error] Credit card has expired. Cannot make purchase.");
                return;
            }
            if (amount > AvailableCredit())
            {
                Console.WriteLine($"  [Error] Purchase of {amount:F2} exceeds available credit ({AvailableCredit():F2}).");
                return;
            }
            totalSpent += amount;
            Console.WriteLine($"  [OK] Purchase of {amount:F2} successful.");
        }

        public void ShowInfo()
        {
            string status = IsValid() ? "Valid" : "Expired";
            Console.WriteLine($"  Card Number        : {cardNumber}");
            Console.WriteLine($"  Expiration Date    : {expirationDate:yyyy-MM-dd}");
            Console.WriteLine($"  Status             : {status}");
            Console.WriteLine($"  Credit Limit       : {creditLimit:F2}");
            Console.WriteLine($"  Total Spent        : {totalSpent:F2}");
            Console.WriteLine($"  Available Credit   : {AvailableCredit():F2}");
            Console.WriteLine($"  Outstanding Balance: {OutstandingBalance():F2}");
        }
    }

    /** Customer with personal info and an associated CreditCard (1-1). */
    class Customer
    {
        private const int MinAge = 18;

        private string name;
        private DateTime dateOfBirth;
        private CreditCard creditCard;

        public Customer(string name, DateTime dateOfBirth, CreditCard creditCard)
        {
            int age = CalculateAge(dateOfBirth);
            if (age < MinAge)
                throw new ArgumentException($"Customer must be at least {MinAge} years old. {name} is {age} years old.");
            this.name = name;
            this.dateOfBirth = dateOfBirth;
            this.creditCard = creditCard;
        }

        private static int CalculateAge(DateTime dob)
        {
            var today = DateTime.Today;
            int age = today.Year - dob.Year;
            if (dob.Date > today.AddYears(-age)) age--;
            return age;
        }

        public int GetAge() => CalculateAge(dateOfBirth);

        public void MakePurchase(double amount)
        {
            Console.WriteLine($"  {name} making purchase of {amount:F2}...");
            creditCard.MakePurchase(amount);
        }

        public void ShowInfo()
        {
            Console.WriteLine($"  Customer Name      : {name}");
            Console.WriteLine($"  Date of Birth      : {dateOfBirth:yyyy-MM-dd}");
            Console.WriteLine($"  Age                : {GetAge()}");
            creditCard.ShowInfo();
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        // --- Create credit cards ---
        var card1 = new CreditCard("4532123456789012", new DateTime(2028, 12, 31));
        var card2 = new CreditCard("5678901234567890", new DateTime(2025, 6, 15)); // Expired

        // --- Create customers ---
        var customer1 = new Customer("Tareq", new DateTime(1990, 5, 15), card1);
        var customer2 = new Customer("Afsana", new DateTime(1995, 8, 20), card2);

        Console.WriteLine("=== Customer Info ===");
        customer1.ShowInfo();
        customer2.ShowInfo();

        // --- Make purchases ---
        Console.WriteLine("=== Purchases ===");
        customer1.MakePurchase(150000);
        customer1.MakePurchase(200000);
        customer1.MakePurchase(200000); // Should fail
        Console.WriteLine();

        customer2.MakePurchase(50000); // Should fail — expired
        Console.WriteLine();

        Console.WriteLine("=== After Purchases ===");
        customer1.ShowInfo();

        // --- Invalid card number ---
        Console.WriteLine("=== Invalid Card Number ===");
        try
        {
            new CreditCard("ABCD-1234", new DateTime(2028, 1, 1));
        }
        catch (ArgumentException e)
        {
            Console.WriteLine($"  [Error] {e.Message}");
        }
        Console.WriteLine();

        // --- Underage customer ---
        Console.WriteLine("=== Underage Customer ===");
        try
        {
            var youngCard = new CreditCard("1111222233334444", new DateTime(2030, 1, 1));
            new Customer("Junior", new DateTime(2015, 1, 1), youngCard);
        }
        catch (ArgumentException e)
        {
            Console.WriteLine($"  [Error] {e.Message}");
        }
    }
}
