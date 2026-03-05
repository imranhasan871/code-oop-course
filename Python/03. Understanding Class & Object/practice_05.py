"""
Practice 05: Credit Card
Task: Create a CreditCard class with spending limits.
      Max credit limit: 500,000. Cash withdrawal: daily limit 100,000,
      per-transaction limit 20,000. Bill payments: no per-txn limit.

How to run:
  python practice_05.py

Key Concepts:
  - Class with class-level constants
  - Multiple validation rules in methods
  - Business logic enforcement
"""


class CreditCard:
    """CreditCard represents a credit card with spending limits."""

    MAX_CREDIT_LIMIT = 500000.00
    DAILY_CASH_LIMIT = 100000.00
    PER_TXN_CASH_LIMIT = 20000.00

    def __init__(self, card_number: str, card_holder: str):
        """Creates a new CreditCard with the given details."""
        self.card_number = card_number
        self.card_holder = card_holder
        self.total_spent = 0.0
        self.daily_cash_withdrawn = 0.0

    def withdraw_cash(self, amount: float):
        """Withdraws cash with per-transaction and daily limits."""
        if amount <= 0:
            print("  [Error] Amount must be greater than 0.")
            return
        if amount > self.PER_TXN_CASH_LIMIT:
            print(f"  [Error] Cash withdrawal exceeds per-transaction limit of "
                  f"{self.PER_TXN_CASH_LIMIT:.2f} (Requested: {amount:.2f})")
            return
        if self.daily_cash_withdrawn + amount > self.DAILY_CASH_LIMIT:
            print(f"  [Error] Cash withdrawal exceeds daily limit of "
                  f"{self.DAILY_CASH_LIMIT:.2f} (Already withdrawn today: "
                  f"{self.daily_cash_withdrawn:.2f}, Requested: {amount:.2f})")
            return
        if self.total_spent + amount > self.MAX_CREDIT_LIMIT:
            print(f"  [Error] Total spending would exceed credit limit of "
                  f"{self.MAX_CREDIT_LIMIT:.2f} (Already spent: "
                  f"{self.total_spent:.2f}, Requested: {amount:.2f})")
            return
        self.total_spent += amount
        self.daily_cash_withdrawn += amount
        print(f"  [OK] Cash withdrawn: {amount:.2f}")

    def pay_bill(self, amount: float):
        """Pays a bill amount. No per-txn limit, but total must stay within max limit."""
        if amount <= 0:
            print("  [Error] Amount must be greater than 0.")
            return
        if self.total_spent + amount > self.MAX_CREDIT_LIMIT:
            print(f"  [Error] Bill payment would exceed credit limit of "
                  f"{self.MAX_CREDIT_LIMIT:.2f} (Already spent: "
                  f"{self.total_spent:.2f}, Requested: {amount:.2f})")
            return
        self.total_spent += amount
        print(f"  [OK] Bill paid: {amount:.2f}")

    def show_info(self):
        """Prints the card details and current spending."""
        print(f"  Card Number          : {self.card_number}")
        print(f"  Card Holder          : {self.card_holder}")
        print(f"  Credit Limit         : {self.MAX_CREDIT_LIMIT:.2f}")
        print(f"  Total Spent          : {self.total_spent:.2f}")
        print(f"  Available Limit      : {self.MAX_CREDIT_LIMIT - self.total_spent:.2f}")
        print(f"  Daily Cash Withdrawn : {self.daily_cash_withdrawn:.2f} / {self.DAILY_CASH_LIMIT:.2f}")
        print()


def main():
    # --- Create a credit card ---
    card = CreditCard("4532-1234-5678-9012", "Tareq")

    print("=== Initial Card Info ===")
    card.show_info()

    # --- Valid cash withdrawal ---
    print("=== Withdraw Cash 15,000 ===")
    card.withdraw_cash(15000)
    card.show_info()

    # --- Exceed per-transaction limit ---
    print("=== Withdraw Cash 25,000 (exceeds per-txn limit) ===")
    card.withdraw_cash(25000)
    card.show_info()

    # --- Multiple valid withdrawals to approach daily limit ---
    print("=== Withdraw Cash 20,000 x4 (should hit daily limit) ===")
    card.withdraw_cash(20000)
    card.withdraw_cash(20000)
    card.withdraw_cash(20000)
    card.withdraw_cash(20000)
    card.show_info()

    # --- Bill payment ---
    print("=== Pay Bill 200,000 ===")
    card.pay_bill(200000)
    card.show_info()

    # --- Try to exceed total credit limit ---
    print("=== Pay Bill 300,000 (should exceed total limit) ===")
    card.pay_bill(300000)
    card.show_info()


if __name__ == "__main__":
    main()
