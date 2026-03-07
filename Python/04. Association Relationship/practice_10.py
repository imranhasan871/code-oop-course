"""
Practice 10: 1-1 Association — Customer & Credit Card
Task: Manage customer profiles with associated credit cards.
      Card number must be digits only. Expiration date validation.
      Credit limit 500K. Track available credit. Age minimum 18.

How to run:
  python practice_10.py

Key Concepts:
  - 1-1 Association (Customer has a CreditCard)
  - Input validation (digit-only card number, age check)
  - Date-based expiration validation
"""

from datetime import date


class CreditCard:
    """CreditCard with card number, expiration date, and credit limit."""

    MAX_CREDIT_LIMIT = 500000.00

    def __init__(self, card_number: str, expiration_date: date, credit_limit: float = 500000.00):
        """Creates a new CreditCard. Card number must contain only digits."""
        if not card_number.isdigit():
            raise ValueError(f"Card number must contain only digits, got: '{card_number}'")
        if credit_limit > self.MAX_CREDIT_LIMIT:
            raise ValueError(f"Credit limit cannot exceed {self.MAX_CREDIT_LIMIT:.2f}")
        self.card_number = card_number
        self.expiration_date = expiration_date
        self.credit_limit = credit_limit
        self.total_spent = 0.0

    def is_valid(self) -> bool:
        """Returns True if the card has not expired."""
        return date.today() <= self.expiration_date

    def available_credit(self) -> float:
        """Returns the remaining credit available."""
        return self.credit_limit - self.total_spent

    def outstanding_balance(self) -> float:
        """Returns the outstanding balance (amount spent)."""
        return self.total_spent

    def make_purchase(self, amount: float):
        """Makes a purchase if within available credit."""
        if amount <= 0:
            print("  [Error] Purchase amount must be greater than 0.")
            return
        if not self.is_valid():
            print("  [Error] Credit card has expired. Cannot make purchase.")
            return
        if amount > self.available_credit():
            print(f"  [Error] Purchase of {amount:.2f} exceeds available credit "
                  f"({self.available_credit():.2f}).")
            return
        self.total_spent += amount
        print(f"  [OK] Purchase of {amount:.2f} successful.")

    def show_info(self):
        """Prints credit card details."""
        status = "Valid" if self.is_valid() else "Expired"
        print(f"  Card Number        : {self.card_number}")
        print(f"  Expiration Date    : {self.expiration_date}")
        print(f"  Status             : {status}")
        print(f"  Credit Limit       : {self.credit_limit:.2f}")
        print(f"  Total Spent        : {self.total_spent:.2f}")
        print(f"  Available Credit   : {self.available_credit():.2f}")
        print(f"  Outstanding Balance: {self.outstanding_balance():.2f}")


class Customer:
    """Customer with personal info and an associated CreditCard (1-1)."""

    MIN_AGE = 18

    def __init__(self, name: str, date_of_birth: date, credit_card: CreditCard):
        """Creates a Customer. Must be at least 18 years old."""
        age = self._calculate_age(date_of_birth)
        if age < self.MIN_AGE:
            raise ValueError(f"Customer must be at least {self.MIN_AGE} years old. "
                             f"{name} is {age} years old.")
        self.name = name
        self.date_of_birth = date_of_birth
        self.credit_card = credit_card

    @staticmethod
    def _calculate_age(dob: date) -> int:
        """Calculates age from date of birth."""
        today = date.today()
        age = today.year - dob.year
        if (today.month, today.day) < (dob.month, dob.day):
            age -= 1
        return age

    def get_age(self) -> int:
        """Returns the customer's current age."""
        return self._calculate_age(self.date_of_birth)

    def make_purchase(self, amount: float):
        """Makes a purchase using the customer's credit card."""
        print(f"  {self.name} making purchase of {amount:.2f}...")
        self.credit_card.make_purchase(amount)

    def show_info(self):
        """Prints customer and credit card details."""
        print(f"  Customer Name      : {self.name}")
        print(f"  Date of Birth      : {self.date_of_birth}")
        print(f"  Age                : {self.get_age()}")
        self.credit_card.show_info()
        print()


def main():
    # --- Create credit cards ---
    card1 = CreditCard("4532123456789012", date(2028, 12, 31))
    card2 = CreditCard("5678901234567890", date(2025, 6, 15))  # Already expired

    # --- Create customers ---
    customer1 = Customer("Tareq", date(1990, 5, 15), card1)
    customer2 = Customer("Afsana", date(1995, 8, 20), card2)

    print("=== Customer Info ===")
    customer1.show_info()
    customer2.show_info()

    # --- Make purchases ---
    print("=== Purchases ===")
    customer1.make_purchase(150000)
    customer1.make_purchase(200000)
    customer1.make_purchase(200000)  # Should fail — exceeds available credit
    print()

    customer2.make_purchase(50000)   # Should fail — expired card
    print()

    print("=== After Purchases ===")
    customer1.show_info()

    # --- Invalid card number ---
    print("=== Invalid Card Number ===")
    try:
        bad_card = CreditCard("ABCD-1234", date(2028, 1, 1))
    except ValueError as e:
        print(f"  [Error] {e}")
    print()

    # --- Underage customer ---
    print("=== Underage Customer ===")
    try:
        young_card = CreditCard("1111222233334444", date(2030, 1, 1))
        young_customer = Customer("Junior", date(2015, 1, 1), young_card)
    except ValueError as e:
        print(f"  [Error] {e}")


if __name__ == "__main__":
    main()
