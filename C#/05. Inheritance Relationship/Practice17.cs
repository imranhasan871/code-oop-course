/**
 * Practice 17: Vehicle Rental System — Inheritance
 * Task: A vehicle rental company rents out Cars, Bikes, and Trucks.
 *       All vehicles share Brand, Model, and Year of Manufacture.
 *       Each type has different rental pricing and discount rules.
 *
 * Rules:
 *   - Car  : $50/day. 10% discount if the car is older than 5 years.
 *   - Bike : $15/day. 15% discount if rental period is more than 7 days.
 *   - Truck: $100/day + $100/day additional maintenance fee.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice17.cs
 *
 * Key Concepts:
 *   - Inheritance (IS-A relationship)
 *   - Base class and subclasses
 *   - Shared properties via parent class
 */

using System;

class Practice17
{
    class Vehicle
    {
        public string Brand;
        public string Model;
        public int Year;

        public void ShowInfo()
        {
            Console.WriteLine($"  Brand : {Brand}");
            Console.WriteLine($"  Model : {Model}");
            Console.WriteLine($"  Year  : {Year}");
        }
    }

    class Car : Vehicle
    {
        private const double BaseRate = 50.0;

        public Car(string brand, string model, int year)
        {
            Brand = brand;
            Model = model;
            Year = year;
        }

        public double CalculateRentalCost(int days)
        {
            double cost = BaseRate * days;
            if (2026 - Year > 5)
                cost *= 0.9;
            return cost;
        }
    }

    class Bike : Vehicle
    {
        private const double BaseRate = 15.0;

        public Bike(string brand, string model, int year)
        {
            Brand = brand;
            Model = model;
            Year = year;
        }

        public double CalculateRentalCost(int days)
        {
            double cost = BaseRate * days;
            if (days > 7)
                cost *= 0.85;
            return cost;
        }
    }

    class Truck : Vehicle
    {
        private const double BaseRate = 100.0;
        private const double AdditionalFee = 100.0;

        public Truck(string brand, string model, int year)
        {
            Brand = brand;
            Model = model;
            Year = year;
        }

        public double CalculateRentalCost(int days)
        {
            return (BaseRate + AdditionalFee) * days;
        }
    }

    static void Main(string[] args)
    {
        var car = new Car("Toyota", "Corolla", 2019);
        var bike = new Bike("Yamaha", "R15", 2023);
        var truck = new Truck("Volvo", "FH16", 2022);

        int days = 10;

        Console.WriteLine("=== Car Rental ===");
        car.ShowInfo();
        Console.WriteLine($"  Rental ({days} days): ${car.CalculateRentalCost(days):F2}");
        Console.WriteLine();

        Console.WriteLine("=== Bike Rental ===");
        bike.ShowInfo();
        Console.WriteLine($"  Rental ({days} days): ${bike.CalculateRentalCost(days):F2}");
        Console.WriteLine();

        Console.WriteLine("=== Truck Rental ===");
        truck.ShowInfo();
        Console.WriteLine($"  Rental ({days} days): ${truck.CalculateRentalCost(days):F2}");
        Console.WriteLine();

        Console.WriteLine("=== Discount Scenarios ===");

        var oldCar = new Car("Honda", "Civic", 2018);
        Console.WriteLine("Old Car (2018):");
        oldCar.ShowInfo();
        Console.WriteLine($"  Rental ({days} days): ${oldCar.CalculateRentalCost(days):F2}");
        Console.WriteLine($"  (10% discount applied — car is {2026 - oldCar.Year} years old)");
        Console.WriteLine();

        var shortBike = new Bike("Honda", "CBR", 2024);
        int shortDays = 5;
        Console.WriteLine($"Bike for {shortDays} days (no discount):");
        shortBike.ShowInfo();
        Console.WriteLine($"  Rental ({shortDays} days): ${shortBike.CalculateRentalCost(shortDays):F2}");
        Console.WriteLine();

        var longBike = new Bike("Suzuki", "Gixxer", 2023);
        int longDays = 10;
        Console.WriteLine($"Bike for {longDays} days (15% discount):");
        longBike.ShowInfo();
        Console.WriteLine($"  Rental ({longDays} days): ${longBike.CalculateRentalCost(longDays):F2}");
    }
}
