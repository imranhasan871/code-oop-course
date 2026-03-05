/**
 * Practice 06: Car Rental System (OOAD)
 * Task: Model a car rental company with a fleet of vehicles.
 *       Customers can rent and return cars. Track availability and maintenance.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice06.cs
 *
 * Or create a Console App and copy contents into Program.cs:
 *   dotnet new console -n Practice06
 *   dotnet run --project Practice06
 *
 * Key Concepts:
 *   - Multiple classes working together (Car, Customer)
 *   - Object references between classes
 *   - State management (available, rented, maintenance)
 */

using System;

class Practice06
{
    /** Car represents a vehicle in the rental fleet. */
    class Car
    {
        private string licensePlate;
        string brand;
        string model;
        private string carType;
        private double dailyRate;
        private bool isAvailable;
        private bool needsMaintenance;

        /** Creates a new Car that is available for rent. */
        public Car(string licensePlate, string brand, string model, string carType, double dailyRate)
        {
            this.licensePlate = licensePlate;
            this.brand = brand;
            this.model = model;
            this.carType = carType;
            this.dailyRate = dailyRate;
            this.isAvailable = true;
            this.needsMaintenance = false;
        }

        /** Marks the car as rented (not available). Returns true if successful. */
        public bool Rent()
        {
            if (!isAvailable)
            {
                Console.WriteLine($"  [Error] {brand} {model} ({licensePlate}) is not available for rent.");
                return false;
            }
            if (needsMaintenance)
            {
                Console.WriteLine($"  [Error] {brand} {model} ({licensePlate}) needs maintenance and cannot be rented.");
                return false;
            }
            isAvailable = false;
            return true;
        }

        /** Marks the car as available again. */
        public void ReturnCar()
        {
            isAvailable = true;
            Console.WriteLine($"  [OK] {brand} {model} ({licensePlate}) has been returned.");
        }

        /** Marks the car as needing maintenance. */
        public void SendToMaintenance()
        {
            needsMaintenance = true;
            isAvailable = false;
            Console.WriteLine($"  [OK] {brand} {model} ({licensePlate}) sent to maintenance.");
        }

        /** Marks maintenance as done and makes car available. */
        public void CompleteMaintenance()
        {
            needsMaintenance = false;
            isAvailable = true;
            Console.WriteLine($"  [OK] {brand} {model} ({licensePlate}) maintenance completed. Now available.");
        }

        /** Prints car details. */
        public void ShowInfo()
        {
            string status = "Available";
            if (needsMaintenance)
                status = "In Maintenance";
            else if (!isAvailable)
                status = "Rented";
            Console.WriteLine($"  [{status}] {brand} {model} | Type: {carType} | Rate: {dailyRate:F2}/day | Plate: {licensePlate}");
        }
    }

    /** Customer represents a person who can rent a car. */
    class Customer
    {
        private string name;
        private string phone;
        private Car rentedCar;

        /** Creates a new Customer with no rented car. */
        public Customer(string name, string phone)
        {
            this.name = name;
            this.phone = phone;
            this.rentedCar = null;
        }

        /** Allows the customer to rent an available car. */
        public void RentCar(Car car)
        {
            if (rentedCar != null)
            {
                Console.WriteLine($"  [Error] {name} already has a rented car.");
                return;
            }
            if (car.Rent())
            {
                rentedCar = car;
                Console.WriteLine($"  [OK] {name} rented a car.");
            }
        }

        /** Allows the customer to return their rented car. */
        public void ReturnCar()
        {
            if (rentedCar == null)
            {
                Console.WriteLine($"  [Error] {name} has no car to return.");
                return;
            }
            rentedCar.ReturnCar();
            Console.WriteLine($"  [OK] {name} returned the car.");
            rentedCar = null;
        }

        /** Prints customer details. */
        public void ShowInfo()
        {
            Console.Write($"  Customer: {name} | Phone: {phone}");
            if (rentedCar != null)
                Console.Write(" | Rented: Yes");
            else
                Console.Write(" | Rented: None");
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        // --- Create fleet of cars ---
        Car car1 = new Car("DHK-1001", "Toyota", "Corolla", "Sedan", 3500);
        Car car2 = new Car("DHK-1002", "Honda", "CR-V", "SUV", 5000);
        Car car3 = new Car("DHK-1003", "Toyota", "HiAce", "Van", 7000);

        Console.WriteLine("=== Fleet Status ===");
        car1.ShowInfo();
        car2.ShowInfo();
        car3.ShowInfo();
        Console.WriteLine();

        // --- Create customers ---
        Customer cust1 = new Customer("Tareq", "01700-000001");
        Customer cust2 = new Customer("Afsana", "01700-000002");

        // --- Rent cars ---
        Console.WriteLine("=== Renting Cars ===");
        cust1.RentCar(car1);
        cust2.RentCar(car2);
        Console.WriteLine();

        Console.WriteLine("=== After Renting ===");
        car1.ShowInfo();
        car2.ShowInfo();
        car3.ShowInfo();
        cust1.ShowInfo();
        cust2.ShowInfo();
        Console.WriteLine();

        // --- Try to rent already rented car ---
        Console.WriteLine("=== Try to rent an already rented car ===");
        cust2.RentCar(car1);
        Console.WriteLine();

        // --- Return car ---
        Console.WriteLine("=== Tareq returns car ===");
        cust1.ReturnCar();
        cust1.ShowInfo();
        car1.ShowInfo();
        Console.WriteLine();

        // --- Send car to maintenance ---
        Console.WriteLine("=== Send car1 to maintenance ===");
        car1.SendToMaintenance();
        car1.ShowInfo();
        Console.WriteLine();

        // --- Try renting car in maintenance ---
        Console.WriteLine("=== Try renting car in maintenance ===");
        cust1.RentCar(car1);
        Console.WriteLine();

        // --- Complete maintenance ---
        Console.WriteLine("=== Complete maintenance ===");
        car1.CompleteMaintenance();
        car1.ShowInfo();
    }
}
