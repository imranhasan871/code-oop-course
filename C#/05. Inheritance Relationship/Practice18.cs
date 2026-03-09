/**
 * Practice 18: Vehicle Rental — Method Overriding & Constructor Chaining
 * Task: Enhance the vehicle rental system with method overriding and
 *       constructor chaining between base class and subclasses.
 *
 * Changes from Practice 17:
 *   - Constructor chaining: subclass constructors call base()
 *   - Method overriding: CalculateRentalCost() is virtual in Vehicle, overridden in subclasses
 *   - ToString() overriding in each class
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice18.cs
 *
 * Key Concepts:
 *   - Method overriding (virtual / override)
 *   - Constructor chaining (: base())
 *   - ToString() overriding
 */

using System;

class Practice18
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

        public virtual double CalculateRentalCost(int days)
        {
            return 0.0;
        }

        public override string ToString()
        {
            return $"{Brand} {Model} ({Year})";
        }
    }

    class Car : Vehicle
    {
        private const double BaseRate = 50.0;

        public Car(string brand, string model, int year) : base(brand, model, year) { }  // Constructor chaining

        public override double CalculateRentalCost(int days)
        {
            double cost = BaseRate * days;
            if (2026 - Year > 5)
                cost *= 0.9;
            return cost;
        }

        public override string ToString() => $"[Car] {base.ToString()}";
    }

    class Bike : Vehicle
    {
        private const double BaseRate = 15.0;

        public Bike(string brand, string model, int year) : base(brand, model, year) { }  // Constructor chaining

        public override double CalculateRentalCost(int days)
        {
            double cost = BaseRate * days;
            if (days > 7)
                cost *= 0.85;
            return cost;
        }

        public override string ToString() => $"[Bike] {base.ToString()}";
    }

    class Truck : Vehicle
    {
        private const double BaseRate = 100.0;
        private const double AdditionalFee = 100.0;

        public Truck(string brand, string model, int year) : base(brand, model, year) { }  // Constructor chaining

        public override double CalculateRentalCost(int days)
        {
            return (BaseRate + AdditionalFee) * days;
        }

        public override string ToString() => $"[Truck] {base.ToString()}";
    }

    static void Main(string[] args)
    {
        Vehicle car = new Car("Toyota", "Corolla", 2019);
        Vehicle bike = new Bike("Yamaha", "R15", 2023);
        Vehicle truck = new Truck("Volvo", "FH16", 2022);

        int days = 10;

        Console.WriteLine("=== Vehicle Rental — Method Overriding ===");
        Console.WriteLine();

        foreach (var vehicle in new Vehicle[] { car, bike, truck })
        {
            Console.WriteLine($"  Vehicle : {vehicle}");                                             // ToString override
            Console.WriteLine($"  Rental ({days} days): ${vehicle.CalculateRentalCost(days):F2}");   // Method override
            Console.WriteLine();
        }

        Console.WriteLine("=== String Representation (ToString) ===");
        Console.WriteLine($"  Car   : {car}");
        Console.WriteLine($"  Bike  : {bike}");
        Console.WriteLine($"  Truck : {truck}");
        Console.WriteLine();

        Console.WriteLine("=== Discount Scenarios ===");

        Vehicle oldCar = new Car("Honda", "Civic", 2018);
        Console.WriteLine($"  {oldCar}");
        Console.WriteLine($"  Age: {2026 - oldCar.Year} years → 10% discount applied");
        Console.WriteLine($"  Rental ({days} days): ${oldCar.CalculateRentalCost(days):F2}");
        Console.WriteLine();

        Vehicle shortBike = new Bike("Honda", "CBR", 2024);
        int shortDays = 5;
        Console.WriteLine($"  {shortBike}");
        Console.WriteLine($"  Rental ({shortDays} days): ${shortBike.CalculateRentalCost(shortDays):F2} (no discount)");
        Console.WriteLine();

        Vehicle longBike = new Bike("Suzuki", "Gixxer", 2023);
        int longDays = 10;
        Console.WriteLine($"  {longBike}");
        Console.WriteLine($"  Rental ({longDays} days): ${longBike.CalculateRentalCost(longDays):F2} (15% discount)");
    }
}
