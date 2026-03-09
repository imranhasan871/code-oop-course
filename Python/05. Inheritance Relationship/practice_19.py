"""
Practice 19: Vehicle Rental — Runtime Polymorphism
Task: Display the brand, model, year, and rental price for one car,
      two bikes, and one truck if each is rented separately for 10 days.

How to run:
  python practice_19.py

Key Concepts:
  - Upcasting (subclass instance treated as base class)
  - Runtime polymorphism (correct method called based on actual type)
  - Iterating a collection of base class references
"""


class Vehicle:
    """Base class for all rental vehicles."""

    def __init__(self, brand: str, model: str, year: int):
        self.brand = brand
        self.model = model
        self.year = year

    def calculate_rental_cost(self, days: int) -> float:
        """Base rental cost — overridden by subclasses."""
        return 0.0

    def get_vehicle_type(self) -> str:
        """Returns the type name — overridden by subclasses."""
        return "Vehicle"

    def __str__(self) -> str:
        return f"{self.brand} {self.model} ({self.year})"


class Car(Vehicle):
    """Car: $50/day, 10% discount if older than 5 years."""

    BASE_RATE = 50.0

    def __init__(self, brand: str, model: str, year: int):
        super().__init__(brand, model, year)

    def calculate_rental_cost(self, days: int) -> float:
        cost = self.BASE_RATE * days
        if 2026 - self.year > 5:
            cost *= 0.9
        return cost

    def get_vehicle_type(self) -> str:
        return "Car"

    def __str__(self) -> str:
        return f"[Car] {super().__str__()}"


class Bike(Vehicle):
    """Bike: $15/day, 15% discount if rental > 7 days."""

    BASE_RATE = 15.0

    def __init__(self, brand: str, model: str, year: int):
        super().__init__(brand, model, year)

    def calculate_rental_cost(self, days: int) -> float:
        cost = self.BASE_RATE * days
        if days > 7:
            cost *= 0.85
        return cost

    def get_vehicle_type(self) -> str:
        return "Bike"

    def __str__(self) -> str:
        return f"[Bike] {super().__str__()}"


class Truck(Vehicle):
    """Truck: $100/day + $100/day additional fee."""

    BASE_RATE = 100.0
    ADDITIONAL_FEE = 100.0

    def __init__(self, brand: str, model: str, year: int):
        super().__init__(brand, model, year)

    def calculate_rental_cost(self, days: int) -> float:
        return (self.BASE_RATE + self.ADDITIONAL_FEE) * days

    def get_vehicle_type(self) -> str:
        return "Truck"

    def __str__(self) -> str:
        return f"[Truck] {super().__str__()}"


def main():
    # --- Upcasting: all stored as Vehicle references ---
    vehicles: list[Vehicle] = [
        Car("Toyota", "Corolla", 2019),      # 1 Car
        Bike("Yamaha", "R15", 2023),          # 2 Bikes
        Bike("Honda", "CBR", 2024),
        Truck("Volvo", "FH16", 2022),         # 1 Truck
    ]

    rental_days = 10

    # --- Runtime polymorphism: correct method called for each type ---
    print(f"=== Vehicle Rental Report ({rental_days} Days) ===")
    print()
    print(f"  {'Type':<8} | {'Brand':<10} | {'Model':<10} | {'Year':<6} | {'Rental Cost':>12}")
    print("  " + "-" * 60)

    total_cost = 0.0
    for vehicle in vehicles:
        cost = vehicle.calculate_rental_cost(rental_days)  # Runtime polymorphism
        total_cost += cost
        print(f"  {vehicle.get_vehicle_type():<8} | {vehicle.brand:<10} | {vehicle.model:<10} | "
              f"{vehicle.year:<6} | ${cost:>10.2f}")

    print("  " + "-" * 60)
    print(f"  {'Total':<8} | {'':<10} | {'':<10} | {'':<6} | ${total_cost:>10.2f}")
    print()

    # --- Demonstrate runtime polymorphism ---
    print("=== Runtime Polymorphism ===")
    print()
    for vehicle in vehicles:
        print(f"  {vehicle}")
        print(f"    Type at runtime : {type(vehicle).__name__}")
        print(f"    Rental (10 days): ${vehicle.calculate_rental_cost(10):.2f}")
        print()


if __name__ == "__main__":
    main()
