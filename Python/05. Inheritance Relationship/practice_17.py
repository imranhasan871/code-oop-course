"""
Practice 17: Vehicle Rental System — Inheritance
Task: A vehicle rental company rents out Cars, Bikes, and Trucks.
      All vehicles share Brand, Model, and Year of Manufacture.
      Each type has different rental pricing and discount rules.

Rules:
  - Car  : $50/day. 10% discount if the car is older than 5 years.
  - Bike : $15/day. 15% discount if rental period is more than 7 days.
  - Truck: $100/day + $100/day additional maintenance fee.

How to run:
  python practice_17.py

Key Concepts:
  - Inheritance (IS-A relationship)
  - Base class and subclasses
  - Shared properties via parent class
"""


class Vehicle:
    """Base class for all rental vehicles."""

    def __init__(self, brand: str, model: str, year: int):
        self.brand = brand
        self.model = model
        self.year = year

    def show_info(self):
        """Prints vehicle details."""
        print(f"  Brand : {self.brand}")
        print(f"  Model : {self.model}")
        print(f"  Year  : {self.year}")


class Car(Vehicle):
    """Car with $50/day base rate. 10% discount if older than 5 years."""

    BASE_RATE = 50.0

    def __init__(self, brand: str, model: str, year: int):
        self.brand = brand
        self.model = model
        self.year = year

    def calculate_rental_cost(self, days: int) -> float:
        """Calculates rental cost for the given number of days."""
        cost = self.BASE_RATE * days
        if 2026 - self.year > 5:
            cost *= 0.9  # 10% discount
        return cost


class Bike(Vehicle):
    """Bike with $15/day base rate. 15% discount if rental > 7 days."""

    BASE_RATE = 15.0

    def __init__(self, brand: str, model: str, year: int):
        self.brand = brand
        self.model = model
        self.year = year

    def calculate_rental_cost(self, days: int) -> float:
        """Calculates rental cost for the given number of days."""
        cost = self.BASE_RATE * days
        if days > 7:
            cost *= 0.85  # 15% discount
        return cost


class Truck(Vehicle):
    """Truck with $100/day base rate + $100/day additional maintenance fee."""

    BASE_RATE = 100.0
    ADDITIONAL_FEE = 100.0

    def __init__(self, brand: str, model: str, year: int):
        self.brand = brand
        self.model = model
        self.year = year

    def calculate_rental_cost(self, days: int) -> float:
        """Calculates rental cost for the given number of days."""
        return (self.BASE_RATE + self.ADDITIONAL_FEE) * days


def main():
    # --- Create vehicles ---
    car = Car("Toyota", "Corolla", 2019)
    bike = Bike("Yamaha", "R15", 2023)
    truck = Truck("Volvo", "FH16", 2022)

    days = 10

    # --- Show rental costs ---
    print("=== Car Rental ===")
    car.show_info()
    print(f"  Rental ({days} days): ${car.calculate_rental_cost(days):.2f}")
    print()

    print("=== Bike Rental ===")
    bike.show_info()
    print(f"  Rental ({days} days): ${bike.calculate_rental_cost(days):.2f}")
    print()

    print("=== Truck Rental ===")
    truck.show_info()
    print(f"  Rental ({days} days): ${truck.calculate_rental_cost(days):.2f}")
    print()

    # --- Discount scenarios ---
    print("=== Discount Scenarios ===")

    old_car = Car("Honda", "Civic", 2018)
    print("Old Car (2018):")
    old_car.show_info()
    print(f"  Rental ({days} days): ${old_car.calculate_rental_cost(days):.2f}")
    print(f"  (10% discount applied — car is {2026 - old_car.year} years old)")
    print()

    short_bike = Bike("Honda", "CBR", 2024)
    short_days = 5
    print(f"Bike for {short_days} days (no discount):")
    short_bike.show_info()
    print(f"  Rental ({short_days} days): ${short_bike.calculate_rental_cost(short_days):.2f}")
    print()

    long_bike = Bike("Suzuki", "Gixxer", 2023)
    long_days = 10
    print(f"Bike for {long_days} days (15% discount):")
    long_bike.show_info()
    print(f"  Rental ({long_days} days): ${long_bike.calculate_rental_cost(long_days):.2f}")


if __name__ == "__main__":
    main()
