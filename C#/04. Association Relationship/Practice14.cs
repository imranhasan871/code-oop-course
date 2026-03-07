/**
 * Practice 14: Smart Parking System
 * Task: Manage vehicle entry, slot assignment, payment, and exit.
 *       Calculate parking charge based on duration.
 *       Update slot allocation and availability in real time.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice14.cs
 *
 * Key Concepts:
 *   - Association between Vehicle, ParkingSlot, and ParkingLot
 *   - Duration-based charge calculation
 *   - Real-time availability tracking
 */

using System;
using System.Collections.Generic;
using System.Linq;

class Practice14
{
    class Vehicle
    {
        public string LicensePlate { get; }
        public string VehicleType { get; }

        public Vehicle(string licensePlate, string vehicleType)
        {
            LicensePlate = licensePlate;
            VehicleType = vehicleType;
        }
    }

    class ParkingSlot
    {
        public string SlotId { get; }
        public string SlotType { get; }
        public double RatePerHour { get; }
        public Vehicle Vehicle { get; set; }
        public DateTime? EntryTime { get; set; }

        public ParkingSlot(string slotId, string slotType, double ratePerHour)
        {
            SlotId = slotId;
            SlotType = slotType;
            RatePerHour = ratePerHour;
        }

        public bool IsAvailable() => Vehicle == null;

        public void AssignVehicle(Vehicle vehicle, DateTime entryTime)
        {
            Vehicle = vehicle;
            EntryTime = entryTime;
        }

        public void ReleaseVehicle()
        {
            Vehicle = null;
            EntryTime = null;
        }

        public void ShowInfo()
        {
            string status = IsAvailable()
                ? "Available"
                : $"Occupied by {Vehicle.LicensePlate}";
            Console.WriteLine($"  {SlotId,-8} | Type: {SlotType,-10} | Rate: {RatePerHour:F0}/hr | {status}");
        }
    }

    class ParkingLot
    {
        public string Name { get; }
        private List<ParkingSlot> slots = new List<ParkingSlot>();

        public ParkingLot(string name) { Name = name; }

        public void AddSlot(ParkingSlot slot) { slots.Add(slot); }

        public ParkingSlot FindAvailableSlot(string slotType)
        {
            return slots.FirstOrDefault(s => s.SlotType == slotType && s.IsAvailable());
        }

        public ParkingSlot ParkVehicle(Vehicle vehicle, DateTime entryTime)
        {
            var slot = FindAvailableSlot(vehicle.VehicleType);
            if (slot == null)
            {
                Console.WriteLine($"  [Error] No available {vehicle.VehicleType} slot for {vehicle.LicensePlate}.");
                return null;
            }
            slot.AssignVehicle(vehicle, entryTime);
            Console.WriteLine($"  [OK] {vehicle.LicensePlate} parked in slot {slot.SlotId} " +
                              $"at {entryTime:yyyy-MM-dd HH:mm}.");
            return slot;
        }

        public double ExitVehicle(string licensePlate, DateTime exitTime)
        {
            foreach (var slot in slots)
            {
                if (!slot.IsAvailable() && slot.Vehicle.LicensePlate == licensePlate)
                {
                    double hours = (exitTime - slot.EntryTime.Value).TotalHours;
                    if (hours < 1) hours = 1;
                    double charge = hours * slot.RatePerHour;
                    Console.WriteLine($"  [OK] {licensePlate} exiting slot {slot.SlotId}.");
                    Console.WriteLine($"       Duration : {hours:F1} hours");
                    Console.WriteLine($"       Rate     : {slot.RatePerHour:F0}/hr");
                    Console.WriteLine($"       Charge   : {charge:F2}");
                    slot.ReleaseVehicle();
                    return charge;
                }
            }
            Console.WriteLine($"  [Error] Vehicle {licensePlate} not found in parking lot.");
            return 0;
        }

        public void ShowStatus()
        {
            int available = slots.Count(s => s.IsAvailable());
            int occupied = slots.Count - available;
            Console.WriteLine($"  Parking Lot: {Name}");
            Console.WriteLine($"  Total Slots: {slots.Count} | Available: {available} | Occupied: {occupied}");
            Console.WriteLine();
            foreach (var slot in slots)
                slot.ShowInfo();
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        var lot = new ParkingLot("City Center Parking");
        lot.AddSlot(new ParkingSlot("A-01", "Car", 50));
        lot.AddSlot(new ParkingSlot("A-02", "Car", 50));
        lot.AddSlot(new ParkingSlot("A-03", "Car", 50));
        lot.AddSlot(new ParkingSlot("B-01", "Bike", 20));
        lot.AddSlot(new ParkingSlot("B-02", "Bike", 20));
        lot.AddSlot(new ParkingSlot("C-01", "Truck", 100));

        Console.WriteLine("=== Initial Parking Status ===");
        lot.ShowStatus();

        var now = new DateTime(2026, 3, 7, 9, 0, 0);
        Console.WriteLine("=== Parking Vehicles ===");
        lot.ParkVehicle(new Vehicle("DHK-1234", "Car"), now);
        lot.ParkVehicle(new Vehicle("DHK-5678", "Car"), now.AddMinutes(15));
        lot.ParkVehicle(new Vehicle("DHK-9012", "Bike"), now.AddMinutes(30));
        lot.ParkVehicle(new Vehicle("CTG-3456", "Truck"), now.AddHours(1));
        Console.WriteLine();

        Console.WriteLine("=== After Parking ===");
        lot.ShowStatus();

        Console.WriteLine("=== Vehicle Exits ===");
        lot.ExitVehicle("DHK-1234", now.AddHours(3));
        Console.WriteLine();
        lot.ExitVehicle("DHK-9012", now.AddHours(1).AddMinutes(30));
        Console.WriteLine();
        lot.ExitVehicle("CTG-3456", now.AddHours(5));
        Console.WriteLine();

        Console.WriteLine("=== Final Parking Status ===");
        lot.ShowStatus();

        Console.WriteLine("=== Parking When Slots Are Limited ===");
        lot.ParkVehicle(new Vehicle("SYL-7777", "Car"), now.AddHours(4));
        lot.ParkVehicle(new Vehicle("SYL-8888", "Car"), now.AddHours(4));
        lot.ParkVehicle(new Vehicle("SYL-9999", "Car"), now.AddHours(4));
    }
}
