"""
Practice 09: Bank Account Collection
Task: Manage a list of BankAccount objects. Perform transactions,
      calculate total balance, and find accounts with highest/lowest balance.

How to run:
  python practice_09.py

Key Concepts:
  - List of objects (collection of objects)
  - Iterating over a collection to compute aggregates
  - Finding min/max in a collection of objects
"""


class BankAccount:
    """BankAccount represents a bank account with basic operations."""

    def __init__(self, account_number: str, account_name: str, balance: float):
        """Creates a new BankAccount with the given details."""
        self.account_number = account_number
        self.account_name = account_name
        self.balance = balance

    def deposit(self, amount: float):
        """Deposits the given amount into the account."""
        if amount <= 0:
            print("  [Error] Deposit amount must be greater than 0.")
            return
        self.balance += amount
        print(f"  [OK] Deposited {amount:.2f} to {self.account_name}")

    def withdraw(self, amount: float):
        """Withdraws the given amount from the account."""
        if amount <= 0:
            print("  [Error] Withdrawal amount must be greater than 0.")
            return
        if self.balance < amount:
            print(f"  [Error] Insufficient balance in {self.account_name} "
                  f"(Balance: {self.balance:.2f}, Requested: {amount:.2f})")
            return
        self.balance -= amount
        print(f"  [OK] Withdrew {amount:.2f} from {self.account_name}")

    def show_info(self):
        """Prints the account details in a table row format."""
        print(f"  {self.account_number:<10} | {self.account_name:<12} | Balance: {self.balance:>10.2f}")


def calculate_total_balance(accounts: list) -> float:
    """Returns the sum of all account balances."""
    total = 0.0
    for acc in accounts:
        total += acc.balance
    return total


def find_highest_balance(accounts: list):
    """Returns the account with the highest balance."""
    highest = accounts[0]
    for acc in accounts[1:]:
        if acc.balance > highest.balance:
            highest = acc
    return highest


def find_lowest_balance(accounts: list):
    """Returns the account with the lowest balance."""
    lowest = accounts[0]
    for acc in accounts[1:]:
        if acc.balance < lowest.balance:
            lowest = acc
    return lowest


def show_all_accounts(accounts: list):
    """Prints all accounts in a formatted table."""
    print(f"  {'Account No':<10} | {'Name':<12} | {'Balance'}")
    print("  " + "-" * 43)
    for acc in accounts:
        acc.show_info()
    print()


def main():
    # --- Create a list of bank accounts ---
    accounts = [
        BankAccount("ACC-1001", "Tareq", 15000),
        BankAccount("ACC-1002", "Afsana", 22000),
        BankAccount("ACC-1003", "Imtiaz", 8500),
        BankAccount("ACC-1004", "Pulok", 31000),
        BankAccount("ACC-1005", "Samia", 12000),
    ]

    print("=== All Accounts (Initial) ===")
    show_all_accounts(accounts)

    # --- Perform transactions ---
    print("=== Performing Transactions ===")
    accounts[0].deposit(5000)   # Tareq deposits 5000
    accounts[1].withdraw(3000)  # Afsana withdraws 3000
    accounts[2].deposit(1500)   # Imtiaz deposits 1500
    accounts[3].withdraw(10000) # Pulok withdraws 10000
    accounts[4].deposit(8000)   # Samia deposits 8000
    print()

    print("=== All Accounts (After Transactions) ===")
    show_all_accounts(accounts)

    # --- Calculate total balance ---
    total = calculate_total_balance(accounts)
    print(f"  Total Balance of All Accounts: {total:.2f}")
    print()

    # --- Find highest and lowest balance ---
    highest = find_highest_balance(accounts)
    lowest = find_lowest_balance(accounts)

    print("=== Highest Balance ===")
    highest.show_info()
    print()

    print("=== Lowest Balance ===")
    lowest.show_info()


if __name__ == "__main__":
    main()
