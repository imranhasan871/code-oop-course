/**
 * Practice 11: 1-1 Association — Car & License Plate
 * Task: Manage car details and their associated license plates.
 *       Validate license plate expiration. Track owner, manufacturer,
 *       model, year. Car age must be 20 years or less for renewal.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice11.cs
 *
 * Key Concepts:
 *   - 1-1 Association (Car has a LicensePlate)
 *   - Date-based expiration validation
 *   - Age calculation for renewal eligibility
 */

using System;

class Practice11
{
    /** LicensePlate with plate number, registration and expiration dates. */
    class LicensePlate
    {
        public string PlateNumber { get; set; }
        public DateTime RegistrationDate { get; set; }
        public DateTime ExpirationDate { get; set; }

        public LicensePlate(string plateNumber, DateTime registrationDate, DateTime expirationDate)
        {
            PlateNumber = plateNumber;
            RegistrationDate = registrationDate;
            ExpirationDate = expirationDate;
        }

        public bool IsValid() => DateTime.Today <= ExpirationDate;

        public void ShowInfo()
        {
            string status = IsValid() ? "Valid" : "Expired";
            Console.WriteLine($"  Plate Number       : {PlateNumber}");
            Console.WriteLine($"  Registration Date  : {RegistrationDate:yyyy-MM-dd}");
            Console.WriteLine($"  Expiration Date    : {ExpirationDate:yyyy-MM-dd}");
            Console.WriteLine($"  Status             : {status}");
        }
    }

    /** Car with owner, manufacturer, model, year, and a LicensePlate (1-1). */
    class Car
    {
        private const int MaxAgeForRenewal = 20;

        public string Owner { get; set; }
        public string Manufacturer { get; set; }
        public string Model { get; set; }
        public int Year { get; set; }
        public LicensePlate LicensePlate { get; set; }

        public Car(string owner, string manufacturer, string model, int year, LicensePlate licensePlate)
        {
            Owner = owner;
            Manufacturer = manufacturer;
            Model = model;
            Year = year;
            LicensePlate = licensePlate;
        }

        public int CarAge() => DateTime.Today.Year - Year;
        public bool QualifiesForRenewal() => CarAge() <= MaxAgeForRenewal;

        public void RenewRegistration(DateTime newExpiration)
        {
            if (!QualifiesForRenewal())
            {
                Console.WriteLine($"  [Error] {Manufacturer} {Model} ({Year}) is {CarAge()} years old. " +
                                  $"Maximum age for renewal is {MaxAgeForRenewal} years.");
                return;
            }
            LicensePlate.ExpirationDate = newExpiration;
            Console.WriteLine($"  [OK] Registration renewed for {Manufacturer} {Model} " +
                              $"({LicensePlate.PlateNumber}). New expiration: {newExpiration:yyyy-MM-dd}");
        }

        public void ShowInfo()
        {
            Console.WriteLine($"  Owner              : {Owner}");
            Console.WriteLine($"  Manufacturer       : {Manufacturer}");
            Console.WriteLine($"  Model              : {Model}");
            Console.WriteLine($"  Year               : {Year}");
            Console.WriteLine($"  Car Age            : {CarAge()} years");
            Console.WriteLine($"  Qualifies Renewal  : {(QualifiesForRenewal() ? "Yes" : "No")}");
            LicensePlate.ShowInfo();
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        var plate1 = new LicensePlate("DHK-GA-1234", new DateTime(2023, 1, 1), new DateTime(2028, 1, 1));
        var plate2 = new LicensePlate("DHK-KA-5678", new DateTime(2020, 6, 15), new DateTime(2025, 6, 15));
        var plate3 = new LicensePlate("CTG-GA-9012", new DateTime(2015, 3, 10), new DateTime(2025, 3, 10));

        var car1 = new Car("Tareq", "Toyota", "Corolla", 2020, plate1);
        var car2 = new Car("Afsana", "Honda", "Civic", 2018, plate2);
        var car3 = new Car("Imtiaz", "Nissan", "Sunny", 2000, plate3);

        Console.WriteLine("=== Car Information ===");
        car1.ShowInfo();
        car2.ShowInfo();
        car3.ShowInfo();

        Console.WriteLine("=== License Plate Validation ===");
        Car[] cars = { car1, car2, car3 };
        foreach (var car in cars)
        {
            string valid = car.LicensePlate.IsValid() ? "Valid" : "EXPIRED";
            Console.WriteLine($"  {car.Manufacturer} {car.Model} ({car.LicensePlate.PlateNumber}): {valid}");
        }
        Console.WriteLine();

        Console.WriteLine("=== Registration Renewal ===");
        car1.RenewRegistration(new DateTime(2033, 1, 1));
        car2.RenewRegistration(new DateTime(2030, 6, 15));
        Console.WriteLine();

        Console.WriteLine("=== Renew Old Car (should fail if too old) ===");
        var oldCar = new Car("Robin", "Toyota", "Corona", 2003,
            new LicensePlate("DHK-GA-0001", new DateTime(2005, 1, 1), new DateTime(2025, 1, 1)));
        oldCar.ShowInfo();
        oldCar.RenewRegistration(new DateTime(2030, 1, 1));
        Console.WriteLine();

        Console.WriteLine("=== Very Old Car ===");
        var ancientCar = new Car("Karim", "Datsun", "120Y", 1980,
            new LicensePlate("DHK-TA-0002", new DateTime(1985, 1, 1), new DateTime(2000, 1, 1)));
        ancientCar.ShowInfo();
        ancientCar.RenewRegistration(new DateTime(2030, 1, 1));
    }
}
