"""
Practice 06: Car Rental System (OOAD)
Task: Model a car rental company with a fleet of vehicles.
      Customers can rent and return cars. Track availability and maintenance.

How to run:
  python practice_06.py

Key Concepts:
  - Multiple classes working together (Car, Customer)
  - Object references between classes
  - State management (available, rented, maintenance)
"""


class Car:
    """Car represents a vehicle in the rental fleet."""

    def __init__(self, license_plate: str, brand: str, model: str, car_type: str, daily_rate: float):
        """Creates a new Car that is available for rent."""
        self.license_plate = license_plate
        self.brand = brand
        self.model = model
        self.car_type = car_type
        self.daily_rate = daily_rate
        self.is_available = True
        self.needs_maintenance = False

    def rent(self) -> bool:
        """Marks the car as rented. Returns True if successful."""
        if not self.is_available:
            print(f"  [Error] {self.brand} {self.model} ({self.license_plate}) is not available for rent.")
            return False
        if self.needs_maintenance:
            print(f"  [Error] {self.brand} {self.model} ({self.license_plate}) needs maintenance and cannot be rented.")
            return False
        self.is_available = False
        return True

    def return_car(self):
        """Marks the car as available again."""
        self.is_available = True
        print(f"  [OK] {self.brand} {self.model} ({self.license_plate}) has been returned.")

    def send_to_maintenance(self):
        """Marks the car as needing maintenance."""
        self.needs_maintenance = True
        self.is_available = False
        print(f"  [OK] {self.brand} {self.model} ({self.license_plate}) sent to maintenance.")

    def complete_maintenance(self):
        """Marks maintenance as done and makes car available."""
        self.needs_maintenance = False
        self.is_available = True
        print(f"  [OK] {self.brand} {self.model} ({self.license_plate}) maintenance completed. Now available.")

    def show_info(self):
        """Prints car details."""
        if self.needs_maintenance:
            status = "In Maintenance"
        elif not self.is_available:
            status = "Rented"
        else:
            status = "Available"
        print(f"  [{status}] {self.brand} {self.model} | Type: {self.car_type} "
              f"| Rate: {self.daily_rate:.2f}/day | Plate: {self.license_plate}")


class Customer:
    """Customer represents a person who can rent a car."""

    def __init__(self, name: str, phone: str):
        """Creates a new Customer with no rented car."""
        self.name = name
        self.phone = phone
        self.rented_car = None

    def rent_car(self, car: Car):
        """Allows the customer to rent an available car."""
        if self.rented_car is not None:
            print(f"  [Error] {self.name} already has a rented car "
                  f"({self.rented_car.brand} {self.rented_car.model}).")
            return
        if car.rent():
            self.rented_car = car
            print(f"  [OK] {self.name} rented {car.brand} {car.model} "
                  f"({car.license_plate}) at {car.daily_rate:.2f}/day.")

    def return_car(self):
        """Allows the customer to return their rented car."""
        if self.rented_car is None:
            print(f"  [Error] {self.name} has no car to return.")
            return
        self.rented_car.return_car()
        print(f"  [OK] {self.name} returned the car.")
        self.rented_car = None

    def show_info(self):
        """Prints customer details."""
        rented = "None"
        if self.rented_car is not None:
            rented = f"{self.rented_car.brand} {self.rented_car.model} ({self.rented_car.license_plate})"
        print(f"  Customer: {self.name} | Phone: {self.phone} | Rented: {rented}")


def main():
    # --- Create fleet of cars ---
    car1 = Car("DHK-1001", "Toyota", "Corolla", "Sedan", 3500)
    car2 = Car("DHK-1002", "Honda", "CR-V", "SUV", 5000)
    car3 = Car("DHK-1003", "Toyota", "HiAce", "Van", 7000)

    print("=== Fleet Status ===")
    car1.show_info()
    car2.show_info()
    car3.show_info()
    print()

    # --- Create customers ---
    cust1 = Customer("Tareq", "01700-000001")
    cust2 = Customer("Afsana", "01700-000002")

    # --- Rent cars ---
    print("=== Renting Cars ===")
    cust1.rent_car(car1)
    cust2.rent_car(car2)
    print()

    print("=== After Renting ===")
    car1.show_info()
    car2.show_info()
    car3.show_info()
    cust1.show_info()
    cust2.show_info()
    print()

    # --- Try to rent already rented car ---
    print("=== Try to rent an already rented car ===")
    cust2.rent_car(car1)
    print()

    # --- Return car ---
    print("=== Tareq returns car ===")
    cust1.return_car()
    cust1.show_info()
    car1.show_info()
    print()

    # --- Send car to maintenance ---
    print("=== Send car1 to maintenance ===")
    car1.send_to_maintenance()
    car1.show_info()
    print()

    # --- Try renting car in maintenance ---
    print("=== Try renting car in maintenance ===")
    cust1.rent_car(car1)
    print()

    # --- Complete maintenance ---
    print("=== Complete maintenance ===")
    car1.complete_maintenance()
    car1.show_info()


if __name__ == "__main__":
    main()
