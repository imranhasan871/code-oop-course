"""
Practice 09: Bank Account Collection

This is an extended practice of Practice 04. In this exercise, you will
create a list of Bank Accounts, perform various transactions (such as
withdrawals and deposits) across multiple accounts, and calculate the
total balance of all the accounts owned by the Bank.

Key Concepts:
  - Reusing the BankAccount class from Practice 04
  - List of objects (object collection)
  - Iterating over a collection to aggregate data
  - Business operations across multiple objects

Course: Professional OOP — by Zohirul Alam Tiemoon
"""


class BankAccount:
    """Represents a bank account with basic operations."""

    def __init__(self, account_number: str, account_name: str, balance: float):
        self._account_number = account_number
        self._account_name = account_name
        self._balance = balance

    @property
    def account_number(self) -> str:
        return self._account_number

    @property
    def account_name(self) -> str:
        return self._account_name

    @property
    def balance(self) -> float:
        return self._balance

    def deposit(self, amount: float):
        if amount <= 0:
            print("  [Error] Deposit amount must be positive.")
            return
        self._balance += amount
        print(f"  [OK] Deposited {amount:.2f} to {self._account_name}. "
              f"New balance: {self._balance:.2f}")

    def withdraw(self, amount: float):
        if amount <= 0:
            print("  [Error] Withdrawal amount must be positive.")
            return
        if amount > self._balance:
            print(f"  [Error] Insufficient balance in {self._account_name}. "
                  f"Available: {self._balance:.2f}")
            return
        self._balance -= amount
        print(f"  [OK] Withdrew {amount:.2f} from {self._account_name}. "
              f"New balance: {self._balance:.2f}")

    def transfer(self, to: "BankAccount", amount: float):
        if amount <= 0:
            print("  [Error] Transfer amount must be positive.")
            return
        if amount > self._balance:
            print(f"  [Error] Insufficient balance in {self._account_name}. "
                  f"Available: {self._balance:.2f}")
            return
        self._balance -= amount
        to._balance += amount
        print(f"  [OK] Transferred {amount:.2f} from {self._account_name} "
              f"to {to._account_name}")

    def print_info(self):
        print(f"  {self._account_number:<10} | {self._account_name:<10} | "
              f"Balance: {self._balance:>10.2f}")


class Bank:
    """Manages a collection of bank accounts."""

    def __init__(self, bank_name: str):
        self._bank_name = bank_name
        self._accounts = []

    def add_account(self, account: BankAccount):
        self._accounts.append(account)
        print(f"  [OK] Added account {account.account_number} "
              f"({account.account_name}) to {self._bank_name}.")

    def show_all_accounts(self):
        print(f"  {self._bank_name} — All Accounts ({len(self._accounts)}):")
        for acc in self._accounts:
            acc.print_info()

    def get_total_balance(self) -> float:
        total = 0.0
        for acc in self._accounts:
            total += acc.balance
        return total

    def print_total_balance(self):
        print(f"  Total balance of {self._bank_name}: {self.get_total_balance():.2f}")


if __name__ == "__main__":
    print("=== Practice 09: Bank Account Collection ===")
    print()

    # Create bank
    bank = Bank("City Bank")

    # Create accounts
    acc1 = BankAccount("ACC-1001", "Imtiaz", 50000)
    acc2 = BankAccount("ACC-1002", "Faria", 30000)
    acc3 = BankAccount("ACC-1003", "Rafi", 45000)
    acc4 = BankAccount("ACC-1004", "Salma", 60000)

    # Add accounts to bank
    print("--- Add Accounts ---")
    bank.add_account(acc1)
    bank.add_account(acc2)
    bank.add_account(acc3)
    bank.add_account(acc4)
    print()

    # Show all accounts
    print("--- All Accounts ---")
    bank.show_all_accounts()
    print()

    # Total balance before transactions
    print("--- Total Balance (Before) ---")
    bank.print_total_balance()
    print()

    # Perform transactions
    print("--- Transactions ---")
    acc1.deposit(10000)
    acc2.withdraw(5000)
    acc3.transfer(acc4, 15000)
    acc1.transfer(acc2, 20000)
    print()

    # Show all accounts after transactions
    print("--- All Accounts (After Transactions) ---")
    bank.show_all_accounts()
    print()

    # Total balance after transactions (should be same — money moves within bank)
    print("--- Total Balance (After) ---")
    bank.print_total_balance()
