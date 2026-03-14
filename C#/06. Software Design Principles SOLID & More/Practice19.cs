/**
 * Practice 19: Applying DRY (Don't Repeat Yourself)
 * Task: Refactor duplicated discount and invoice logic into reusable components.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice19.cs
 */

using System;
using System.Collections.Generic;

class Practice19
{
    class OrderItem
    {
        public string Name;
        public decimal UnitPrice;
        public int Quantity;

        public OrderItem(string name, decimal unitPrice, int quantity)
        {
            Name = name;
            UnitPrice = unitPrice;
            Quantity = quantity;
        }

        public decimal LineTotal => UnitPrice * Quantity;
    }

    static class DiscountPolicy
    {
        public static decimal Calculate(decimal subtotal, string customerType)
        {
            var key = customerType.ToLowerInvariant();
            decimal rate = key == "gold" ? 0.10m : key == "silver" ? 0.05m : 0.00m;
            return subtotal * rate;
        }
    }

    class InvoiceSummary
    {
        public decimal Subtotal;
        public decimal Discount;
        public decimal Vat;
        public decimal Total;
    }

    class InvoiceCalculator
    {
        private const decimal VatRate = 0.05m;

        public InvoiceSummary Summarize(List<OrderItem> items, string customerType)
        {
            decimal subtotal = 0m;
            foreach (var item in items)
                subtotal += item.LineTotal;

            decimal discount = DiscountPolicy.Calculate(subtotal, customerType);
            decimal taxable = subtotal - discount;
            decimal vat = taxable * VatRate;

            return new InvoiceSummary
            {
                Subtotal = subtotal,
                Discount = discount,
                Vat = vat,
                Total = taxable + vat
            };
        }
    }

    static void PrintInvoice(string customerName, string customerType, List<OrderItem> items)
    {
        var calculator = new InvoiceCalculator();
        var summary = calculator.Summarize(items, customerType);

        Console.WriteLine($"Customer: {customerName} ({customerType})");
        foreach (var item in items)
            Console.WriteLine($"  - {item.Name,-10} x{item.Quantity,-2} @ ${item.UnitPrice,6:F2} = ${item.LineTotal,7:F2}");

        Console.WriteLine($"Subtotal : ${summary.Subtotal:F2}");
        Console.WriteLine($"Discount : ${summary.Discount:F2}");
        Console.WriteLine($"VAT (5%) : ${summary.Vat:F2}");
        Console.WriteLine($"Total    : ${summary.Total:F2}");
        Console.WriteLine();
    }

    static void Main(string[] args)
    {
        var items = new List<OrderItem>
        {
            new OrderItem("Keyboard", 45.0m, 2),
            new OrderItem("Mouse", 20.0m, 1),
            new OrderItem("USBC", 12.5m, 3)
        };

        PrintInvoice("Samia", "regular", items);
        PrintInvoice("Afsana", "silver", items);
        PrintInvoice("Hasan", "gold", items);
    }
}
