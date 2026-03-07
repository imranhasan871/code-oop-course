"""
Practice 14: Smart Parking System
Task: Manage vehicle entry, slot assignment, payment, and exit.
      Calculate parking charge based on duration.
      Update slot allocation and availability in real time.

How to run:
  python practice_14.py

Key Concepts:
  - Association between Vehicle, ParkingSlot, and ParkingLot
  - Duration-based charge calculation
  - Real-time availability tracking
"""

from datetime import datetime, timedelta


class Vehicle:
    """Vehicle with license plate and type."""

    def __init__(self, license_plate: str, vehicle_type: str):
        """Creates a new Vehicle."""
        self.license_plate = license_plate
        self.vehicle_type = vehicle_type

    def show_info(self):
        """Prints vehicle details."""
        print(f"  License Plate : {self.license_plate}")
        print(f"  Vehicle Type  : {self.vehicle_type}")


class ParkingSlot:
    """ParkingSlot that can hold one vehicle at a time."""

    def __init__(self, slot_id: str, slot_type: str, rate_per_hour: float):
        """Creates a new ParkingSlot."""
        self.slot_id = slot_id
        self.slot_type = slot_type
        self.rate_per_hour = rate_per_hour
        self.vehicle = None
        self.entry_time = None

    def is_available(self) -> bool:
        """Returns True if the slot is empty."""
        return self.vehicle is None

    def assign_vehicle(self, vehicle: Vehicle, entry_time: datetime):
        """Assigns a vehicle to this slot."""
        self.vehicle = vehicle
        self.entry_time = entry_time

    def release_vehicle(self):
        """Releases the vehicle from this slot."""
        self.vehicle = None
        self.entry_time = None

    def show_info(self):
        """Prints slot details."""
        status = "Available" if self.is_available() else f"Occupied by {self.vehicle.license_plate}"
        print(f"  {self.slot_id:<8} | Type: {self.slot_type:<10} | "
              f"Rate: {self.rate_per_hour:.0f}/hr | {status}")


class ParkingLot:
    """ParkingLot manages a collection of ParkingSlots."""

    def __init__(self, name: str):
        """Creates a new ParkingLot."""
        self.name = name
        self.slots = []

    def add_slot(self, slot: ParkingSlot):
        """Adds a parking slot to the lot."""
        self.slots.append(slot)

    def find_available_slot(self, slot_type: str):
        """Finds the first available slot of the given type."""
        for slot in self.slots:
            if slot.slot_type == slot_type and slot.is_available():
                return slot
        return None

    def park_vehicle(self, vehicle: Vehicle, entry_time: datetime):
        """Parks a vehicle in the first available matching slot."""
        slot = self.find_available_slot(vehicle.vehicle_type)
        if slot is None:
            print(f"  [Error] No available {vehicle.vehicle_type} slot for "
                  f"{vehicle.license_plate}.")
            return None
        slot.assign_vehicle(vehicle, entry_time)
        print(f"  [OK] {vehicle.license_plate} parked in slot {slot.slot_id} "
              f"at {entry_time.strftime('%Y-%m-%d %H:%M')}.")
        return slot

    def exit_vehicle(self, license_plate: str, exit_time: datetime) -> float:
        """Removes a vehicle and calculates the parking charge."""
        for slot in self.slots:
            if not slot.is_available() and slot.vehicle.license_plate == license_plate:
                duration = exit_time - slot.entry_time
                hours = duration.total_seconds() / 3600
                if hours < 1:
                    hours = 1  # Minimum 1 hour charge
                charge = hours * slot.rate_per_hour
                print(f"  [OK] {license_plate} exiting slot {slot.slot_id}.")
                print(f"       Duration : {hours:.1f} hours")
                print(f"       Rate     : {slot.rate_per_hour:.0f}/hr")
                print(f"       Charge   : {charge:.2f}")
                slot.release_vehicle()
                return charge
        print(f"  [Error] Vehicle {license_plate} not found in parking lot.")
        return 0.0

    def show_status(self):
        """Prints the status of all slots."""
        available = sum(1 for s in self.slots if s.is_available())
        occupied = len(self.slots) - available
        print(f"  Parking Lot: {self.name}")
        print(f"  Total Slots: {len(self.slots)} | Available: {available} | Occupied: {occupied}")
        print()
        for slot in self.slots:
            slot.show_info()
        print()


def main():
    # --- Create parking lot with slots ---
    lot = ParkingLot("City Center Parking")
    lot.add_slot(ParkingSlot("A-01", "Car", 50.0))
    lot.add_slot(ParkingSlot("A-02", "Car", 50.0))
    lot.add_slot(ParkingSlot("A-03", "Car", 50.0))
    lot.add_slot(ParkingSlot("B-01", "Bike", 20.0))
    lot.add_slot(ParkingSlot("B-02", "Bike", 20.0))
    lot.add_slot(ParkingSlot("C-01", "Truck", 100.0))

    print("=== Initial Parking Status ===")
    lot.show_status()

    # --- Park vehicles ---
    now = datetime(2026, 3, 7, 9, 0)
    print("=== Parking Vehicles ===")
    v1 = Vehicle("DHK-1234", "Car")
    v2 = Vehicle("DHK-5678", "Car")
    v3 = Vehicle("DHK-9012", "Bike")
    v4 = Vehicle("CTG-3456", "Truck")

    lot.park_vehicle(v1, now)
    lot.park_vehicle(v2, now + timedelta(minutes=15))
    lot.park_vehicle(v3, now + timedelta(minutes=30))
    lot.park_vehicle(v4, now + timedelta(hours=1))
    print()

    print("=== After Parking ===")
    lot.show_status()

    # --- Exit vehicles ---
    print("=== Vehicle Exits ===")
    lot.exit_vehicle("DHK-1234", now + timedelta(hours=3))          # 3 hours
    print()
    lot.exit_vehicle("DHK-9012", now + timedelta(hours=1, minutes=30))  # 1.5 hours
    print()
    lot.exit_vehicle("CTG-3456", now + timedelta(hours=5))          # 4 hours (entered 1hr late)
    print()

    print("=== Final Parking Status ===")
    lot.show_status()

    # --- Try parking when full ---
    print("=== Parking When Slots Are Limited ===")
    v5 = Vehicle("SYL-7777", "Car")
    v6 = Vehicle("SYL-8888", "Car")
    v7 = Vehicle("SYL-9999", "Car")  # Should fail — only 2 Car slots left
    lot.park_vehicle(v5, now + timedelta(hours=4))
    lot.park_vehicle(v6, now + timedelta(hours=4))
    lot.park_vehicle(v7, now + timedelta(hours=4))  # No available slot


if __name__ == "__main__":
    main()
