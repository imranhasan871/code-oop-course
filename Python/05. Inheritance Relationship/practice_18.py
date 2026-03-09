"""
Practice 18: Vehicle Rental — Method Overriding & Constructor Chaining
Task: Enhance the vehicle rental system with method overriding and
      constructor chaining between base class and subclasses.

Changes from Practice 17:
  - Constructor chaining: subclass constructors call super().__init__()
  - Method overriding: calculate_rental_cost() defined in Vehicle base class,
    overridden in each subclass
  - __str__ overriding: each class provides its own string representation

How to run:
  python practice_18.py

Key Concepts:
  - Method overriding
  - Constructor chaining (super().__init__())
  - __str__ / toString overriding
"""


class Vehicle:
    """Base class for all rental vehicles."""

    def __init__(self, brand: str, model: str, year: int):
        self.brand = brand
        self.model = model
        self.year = year

    def calculate_rental_cost(self, days: int) -> float:
        """Base rental cost calculation — overridden by subclasses."""
        return 0.0

    def __str__(self) -> str:
        return f"{self.brand} {self.model} ({self.year})"


class Car(Vehicle):
    """Car: $50/day, 10% discount if older than 5 years."""

    BASE_RATE = 50.0

    def __init__(self, brand: str, model: str, year: int):
        super().__init__(brand, model, year)  # Constructor chaining

    def calculate_rental_cost(self, days: int) -> float:
        """Overrides Vehicle.calculate_rental_cost()."""
        cost = self.BASE_RATE * days
        if 2026 - self.year > 5:
            cost *= 0.9
        return cost

    def __str__(self) -> str:
        return f"[Car] {super().__str__()}"


class Bike(Vehicle):
    """Bike: $15/day, 15% discount if rental > 7 days."""

    BASE_RATE = 15.0

    def __init__(self, brand: str, model: str, year: int):
        super().__init__(brand, model, year)  # Constructor chaining

    def calculate_rental_cost(self, days: int) -> float:
        """Overrides Vehicle.calculate_rental_cost()."""
        cost = self.BASE_RATE * days
        if days > 7:
            cost *= 0.85
        return cost

    def __str__(self) -> str:
        return f"[Bike] {super().__str__()}"


class Truck(Vehicle):
    """Truck: $100/day + $100/day additional fee."""

    BASE_RATE = 100.0
    ADDITIONAL_FEE = 100.0

    def __init__(self, brand: str, model: str, year: int):
        super().__init__(brand, model, year)  # Constructor chaining

    def calculate_rental_cost(self, days: int) -> float:
        """Overrides Vehicle.calculate_rental_cost()."""
        return (self.BASE_RATE + self.ADDITIONAL_FEE) * days

    def __str__(self) -> str:
        return f"[Truck] {super().__str__()}"


def main():
    car = Car("Toyota", "Corolla", 2019)
    bike = Bike("Yamaha", "R15", 2023)
    truck = Truck("Volvo", "FH16", 2022)

    days = 10

    # --- Method overriding: calculate_rental_cost() ---
    print("=== Vehicle Rental — Method Overriding ===")
    print()
    for vehicle in [car, bike, truck]:
        print(f"  Vehicle : {vehicle}")                                         # __str__ override
        print(f"  Rental ({days} days): ${vehicle.calculate_rental_cost(days):.2f}")  # Method override
        print()

    # --- __str__ overriding ---
    print("=== String Representation (__str__) ===")
    print(f"  Car   : {car}")
    print(f"  Bike  : {bike}")
    print(f"  Truck : {truck}")
    print()

    # --- Discount scenarios ---
    print("=== Discount Scenarios ===")

    old_car = Car("Honda", "Civic", 2018)
    print(f"  {old_car}")
    print(f"  Age: {2026 - old_car.year} years → 10% discount applied")
    print(f"  Rental ({days} days): ${old_car.calculate_rental_cost(days):.2f}")
    print()

    short_bike = Bike("Honda", "CBR", 2024)
    short_days = 5
    print(f"  {short_bike}")
    print(f"  Rental ({short_days} days): ${short_bike.calculate_rental_cost(short_days):.2f} (no discount)")
    print()

    long_bike = Bike("Suzuki", "Gixxer", 2023)
    long_days = 10
    print(f"  {long_bike}")
    print(f"  Rental ({long_days} days): ${long_bike.calculate_rental_cost(long_days):.2f} (15% discount)")


if __name__ == "__main__":
    main()
