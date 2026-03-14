/**
 * Practice 18: Applying DIP (Dependency Inversion Principle)
 * Task: Build a payment system where PaymentService depends on abstraction,
 *       not concrete payment providers.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice18.cs
 */

using System;
using System.Collections.Generic;

class Practice18
{
    interface IPaymentMethod
    {
        string Pay(decimal amount);
    }

    class CreditCardPayment : IPaymentMethod
    {
        public string Pay(decimal amount) => $"Paid ${amount:F2} using Credit Card";
    }

    class BkashPayment : IPaymentMethod
    {
        public string Pay(decimal amount) => $"Paid ${amount:F2} using bKash";
    }

    class PaypalPayment : IPaymentMethod
    {
        public string Pay(decimal amount) => $"Paid ${amount:F2} using PayPal";
    }

    class BankTransferPayment : IPaymentMethod
    {
        public string Pay(decimal amount) => $"Paid ${amount:F2} using Bank Transfer";
    }

    class PaymentService
    {
        public void ProcessPayment(IPaymentMethod paymentMethod, decimal amount)
        {
            if (amount <= 0)
                throw new ArgumentException("Amount must be greater than 0.");

            Console.WriteLine($"[PaymentService] {paymentMethod.Pay(amount)}");
        }
    }

    static void Main(string[] args)
    {
        var service = new PaymentService();
        var methods = new List<IPaymentMethod>
        {
            new CreditCardPayment(),
            new BkashPayment(),
            new PaypalPayment(),
            new BankTransferPayment()
        };

        foreach (var method in methods)
            service.ProcessPayment(method, 1500m);
    }
}
