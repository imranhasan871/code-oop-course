/**
 * Practice 19: Vehicle Rental — Runtime Polymorphism
 * Task: Display brand, model, year, and rental price for one car,
 *       two bikes, and one truck if each is rented for 10 days.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice19.cs
 *
 * Key Concepts:
 *   - Upcasting (subclass → base class reference)
 *   - Runtime polymorphism (correct method called at runtime)
 *   - Iterating an array of base class references
 */

using System;

class Practice19
{
    class Vehicle
    {
        public string Brand { get; }
        public string Model { get; }
        public int Year { get; }

        public Vehicle(string brand, string model, int year)
        {
            Brand = brand;
            Model = model;
            Year = year;
        }

        public virtual double CalculateRentalCost(int days) => 0.0;
        public virtual string GetVehicleType() => "Vehicle";

        public override string ToString() => $"{Brand} {Model} ({Year})";
    }

    class Car : Vehicle
    {
        private const double BaseRate = 50.0;

        public Car(string brand, string model, int year) : base(brand, model, year) { }

        public override double CalculateRentalCost(int days)
        {
            double cost = BaseRate * days;
            if (2026 - Year > 5)
                cost *= 0.9;
            return cost;
        }

        public override string GetVehicleType() => "Car";
        public override string ToString() => $"[Car] {base.ToString()}";
    }

    class Bike : Vehicle
    {
        private const double BaseRate = 15.0;

        public Bike(string brand, string model, int year) : base(brand, model, year) { }

        public override double CalculateRentalCost(int days)
        {
            double cost = BaseRate * days;
            if (days > 7)
                cost *= 0.85;
            return cost;
        }

        public override string GetVehicleType() => "Bike";
        public override string ToString() => $"[Bike] {base.ToString()}";
    }

    class Truck : Vehicle
    {
        private const double BaseRate = 100.0;
        private const double AdditionalFee = 100.0;

        public Truck(string brand, string model, int year) : base(brand, model, year) { }

        public override double CalculateRentalCost(int days)
        {
            return (BaseRate + AdditionalFee) * days;
        }

        public override string GetVehicleType() => "Truck";
        public override string ToString() => $"[Truck] {base.ToString()}";
    }

    static void Main(string[] args)
    {
        // Upcasting: all stored as Vehicle references
        Vehicle[] vehicles = {
            new Car("Toyota", "Corolla", 2019),       // 1 Car
            new Bike("Yamaha", "R15", 2023),           // 2 Bikes
            new Bike("Honda", "CBR", 2024),
            new Truck("Volvo", "FH16", 2022),          // 1 Truck
        };

        int rentalDays = 10;

        Console.WriteLine($"=== Vehicle Rental Report ({rentalDays} Days) ===");
        Console.WriteLine();
        Console.WriteLine($"  {"Type",-8} | {"Brand",-10} | {"Model",-10} | {"Year",-6} | {"Rental Cost",12}");
        Console.WriteLine("  " + new string('-', 60));

        double totalCost = 0;
        foreach (var vehicle in vehicles)
        {
            double cost = vehicle.CalculateRentalCost(rentalDays); // Runtime polymorphism
            totalCost += cost;
            Console.WriteLine($"  {vehicle.GetVehicleType(),-8} | {vehicle.Brand,-10} | {vehicle.Model,-10} | " +
                              $"{vehicle.Year,-6} | ${cost,10:F2}");
        }

        Console.WriteLine("  " + new string('-', 60));
        Console.WriteLine($"  {"Total",-8} | {"",10} | {"",10} | {"",6} | ${totalCost,10:F2}");
        Console.WriteLine();

        // Demonstrate runtime polymorphism
        Console.WriteLine("=== Runtime Polymorphism ===");
        Console.WriteLine();
        foreach (var vehicle in vehicles)
        {
            Console.WriteLine($"  {vehicle}");
            Console.WriteLine($"    Type at runtime : {vehicle.GetType().Name}");
            Console.WriteLine($"    Rental (10 days): ${vehicle.CalculateRentalCost(10):F2}");
            Console.WriteLine();
        }
    }
}
