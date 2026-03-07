"""
Practice 11: 1-1 Association — Car & License Plate
Task: Manage car details and their associated license plates.
      Validate license plate expiration. Track owner, manufacturer,
      model, year. Car age must be 20 years or less for renewal.

How to run:
  python practice_11.py

Key Concepts:
  - 1-1 Association (Car has a LicensePlate)
  - Date-based expiration validation
  - Age calculation for renewal eligibility
"""

from datetime import date


class LicensePlate:
    """LicensePlate with plate number, registration and expiration dates."""

    def __init__(self, plate_number: str, registration_date: date, expiration_date: date):
        """Creates a new LicensePlate."""
        self.plate_number = plate_number
        self.registration_date = registration_date
        self.expiration_date = expiration_date

    def is_valid(self) -> bool:
        """Returns True if the license plate has not expired."""
        return date.today() <= self.expiration_date

    def show_info(self):
        """Prints license plate details."""
        status = "Valid" if self.is_valid() else "Expired"
        print(f"  Plate Number       : {self.plate_number}")
        print(f"  Registration Date  : {self.registration_date}")
        print(f"  Expiration Date    : {self.expiration_date}")
        print(f"  Status             : {status}")


class Car:
    """Car with owner, manufacturer, model, year, and a LicensePlate (1-1)."""

    MAX_AGE_FOR_RENEWAL = 20

    def __init__(self, owner: str, manufacturer: str, model: str,
                 year: int, license_plate: LicensePlate):
        """Creates a new Car with the given details."""
        self.owner = owner
        self.manufacturer = manufacturer
        self.model = model
        self.year = year
        self.license_plate = license_plate

    def car_age(self) -> int:
        """Returns the age of the car in years."""
        return date.today().year - self.year

    def qualifies_for_renewal(self) -> bool:
        """Returns True if the car is 20 years old or less."""
        return self.car_age() <= self.MAX_AGE_FOR_RENEWAL

    def renew_registration(self, new_expiration: date):
        """Renews the license plate if the car qualifies."""
        if not self.qualifies_for_renewal():
            print(f"  [Error] {self.manufacturer} {self.model} ({self.year}) is "
                  f"{self.car_age()} years old. Maximum age for renewal is "
                  f"{self.MAX_AGE_FOR_RENEWAL} years.")
            return
        self.license_plate.expiration_date = new_expiration
        print(f"  [OK] Registration renewed for {self.manufacturer} {self.model} "
              f"({self.license_plate.plate_number}). New expiration: {new_expiration}")

    def show_info(self):
        """Prints car and license plate details."""
        print(f"  Owner              : {self.owner}")
        print(f"  Manufacturer       : {self.manufacturer}")
        print(f"  Model              : {self.model}")
        print(f"  Year               : {self.year}")
        print(f"  Car Age            : {self.car_age()} years")
        print(f"  Qualifies Renewal  : {'Yes' if self.qualifies_for_renewal() else 'No'}")
        self.license_plate.show_info()
        print()


def main():
    # --- Create license plates ---
    plate1 = LicensePlate("DHK-GA-1234", date(2023, 1, 1), date(2028, 1, 1))
    plate2 = LicensePlate("DHK-KA-5678", date(2020, 6, 15), date(2025, 6, 15))  # Expired
    plate3 = LicensePlate("CTG-GA-9012", date(2015, 3, 10), date(2025, 3, 10))

    # --- Create cars ---
    car1 = Car("Tareq", "Toyota", "Corolla", 2020, plate1)
    car2 = Car("Afsana", "Honda", "Civic", 2018, plate2)
    car3 = Car("Imtiaz", "Nissan", "Sunny", 2000, plate3)  # Old car

    print("=== Car Information ===")
    car1.show_info()
    car2.show_info()
    car3.show_info()

    # --- Validate license plates ---
    print("=== License Plate Validation ===")
    for car in [car1, car2, car3]:
        valid = "Valid" if car.license_plate.is_valid() else "EXPIRED"
        print(f"  {car.manufacturer} {car.model} ({car.license_plate.plate_number}): {valid}")
    print()

    # --- Renew registration ---
    print("=== Registration Renewal ===")
    car1.renew_registration(date(2033, 1, 1))
    car2.renew_registration(date(2030, 6, 15))
    print()

    # --- Try to renew an old car ---
    print("=== Renew Old Car (should fail if too old) ===")
    old_car = Car("Robin", "Toyota", "Corona", 2003, LicensePlate("DHK-GA-0001", date(2005, 1, 1), date(2025, 1, 1)))
    old_car.show_info()
    old_car.renew_registration(date(2030, 1, 1))
    print()

    # --- Very old car that doesn't qualify ---
    print("=== Very Old Car ===")
    ancient_car = Car("Karim", "Datsun", "120Y", 1980, LicensePlate("DHK-TA-0002", date(1985, 1, 1), date(2000, 1, 1)))
    ancient_car.show_info()
    ancient_car.renew_registration(date(2030, 1, 1))


if __name__ == "__main__":
    main()
