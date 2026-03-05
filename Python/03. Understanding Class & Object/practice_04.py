"""
Practice 04: Bank Account
Task: Create a BankAccount class with deposit, withdraw, and transfer methods.
      No negative balance is allowed.

How to run:
  python practice_04.py

Key Concepts:
  - Class with __init__ constructor
  - Instance methods for deposit, withdraw, transfer
  - Input validation (no negative balance)
"""


class BankAccount:
    """BankAccount represents a bank account with basic operations."""

    def __init__(self, account_number: str, account_name: str, balance: float):
        """Creates a new BankAccount with the given details."""
        self.account_number = account_number
        self.account_name = account_name
        self.balance = balance

    def deposit(self, amount: float):
        """Deposits the given amount into the account. Amount must be positive."""
        if amount <= 0:
            print("  [Error] Deposit amount must be greater than 0.")
            return
        self.balance += amount
        print(f"  [OK] Deposited {amount:.2f} to {self.account_name}")

    def withdraw(self, amount: float):
        """Withdraws the given amount from the account. No negative balance allowed."""
        if amount <= 0:
            print("  [Error] Withdrawal amount must be greater than 0.")
            return
        if self.balance < amount:
            print(f"  [Error] Insufficient balance in {self.account_name} "
                  f"(Balance: {self.balance:.2f}, Requested: {amount:.2f})")
            return
        self.balance -= amount
        print(f"  [OK] Withdrew {amount:.2f} from {self.account_name}")

    def transfer(self, amount: float, target: "BankAccount"):
        """Transfers the given amount from this account to the target account."""
        if amount <= 0:
            print("  [Error] Transfer amount must be greater than 0.")
            return
        if self.balance < amount:
            print(f"  [Error] Insufficient balance in {self.account_name} "
                  f"for transfer (Balance: {self.balance:.2f}, Requested: {amount:.2f})")
            return
        self.balance -= amount
        target.balance += amount
        print(f"  [OK] Transferred {amount:.2f} from {self.account_name} to {target.account_name}")

    def show_info(self):
        """Prints the account number, name, and current balance."""
        print(f"  Account Number : {self.account_number}")
        print(f"  Account Name   : {self.account_name}")
        print(f"  Balance        : {self.balance:.2f}")
        print()


def main():
    # --- Create two bank accounts ---
    account1 = BankAccount("ACC-1001", "Tareq", 5000.00)
    account2 = BankAccount("ACC-1002", "Afsana", 3000.00)

    print("=== Initial Account Info ===")
    account1.show_info()
    account2.show_info()

    # --- Deposit ---
    print("=== Deposit 2000 to Tareq ===")
    account1.deposit(2000)
    account1.show_info()

    # --- Withdraw ---
    print("=== Withdraw 1500 from Afsana ===")
    account2.withdraw(1500)
    account2.show_info()

    # --- Withdraw more than balance (should fail) ---
    print("=== Withdraw 5000 from Afsana (should fail) ===")
    account2.withdraw(5000)
    account2.show_info()

    # --- Transfer ---
    print("=== Transfer 3000 from Tareq to Afsana ===")
    account1.transfer(3000, account2)
    print("After transfer:")
    account1.show_info()
    account2.show_info()


if __name__ == "__main__":
    main()
