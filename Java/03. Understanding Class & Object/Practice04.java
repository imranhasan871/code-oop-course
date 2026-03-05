/**
 * Practice 04: Bank Account
 * Task: Create a BankAccount class with deposit, withdraw, and transfer methods.
 *       No negative balance is allowed.
 *
 * How to compile and run:
 *   javac Practice04.java
 *   java Practice04
 *
 * Key Concepts:
 *   - Class with private fields and public methods (encapsulation)
 *   - Constructor for initializing object state
 *   - Input validation in methods
 */

public class Practice04 {

    /** BankAccount represents a bank account with basic operations. */
    static class BankAccount {
        private String accountNumber;
        private String accountName;
        private double balance;

        /** Creates a new BankAccount with the given details. */
        public BankAccount(String accountNumber, String accountName, double balance) {
            this.accountNumber = accountNumber;
            this.accountName = accountName;
            this.balance = balance;
        }

        /** Deposits the given amount into the account. Amount must be positive. */
        public void deposit(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Deposit amount must be greater than 0.");
                return;
            }
            balance += amount;
            System.out.printf("  [OK] Deposited %.2f to %s%n", amount, accountName);
        }

        /** Withdraws the given amount from the account. No negative balance allowed. */
        public void withdraw(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Withdrawal amount must be greater than 0.");
                return;
            }
            if (balance < amount) {
                System.out.printf("  [Error] Insufficient balance in %s (Balance: %.2f, Requested: %.2f)%n",
                        accountName, balance, amount);
                return;
            }
            balance -= amount;
            System.out.printf("  [OK] Withdrew %.2f from %s%n", amount, accountName);
        }

        /** Transfers the given amount from this account to the target account. */
        public void transfer(double amount, BankAccount target) {
            if (amount <= 0) {
                System.out.println("  [Error] Transfer amount must be greater than 0.");
                return;
            }
            if (balance < amount) {
                System.out.printf("  [Error] Insufficient balance in %s for transfer (Balance: %.2f, Requested: %.2f)%n",
                        accountName, balance, amount);
                return;
            }
            balance -= amount;
            target.balance += amount;
            System.out.printf("  [OK] Transferred %.2f from %s to %s%n", amount, accountName, target.accountName);
        }

        /** Prints the account number, name, and current balance. */
        public void showInfo() {
            System.out.printf("  Account Number : %s%n", accountNumber);
            System.out.printf("  Account Name   : %s%n", accountName);
            System.out.printf("  Balance        : %.2f%n", balance);
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // --- Create two bank accounts ---
        BankAccount account1 = new BankAccount("ACC-1001", "Tareq", 5000.00);
        BankAccount account2 = new BankAccount("ACC-1002", "Afsana", 3000.00);

        System.out.println("=== Initial Account Info ===");
        account1.showInfo();
        account2.showInfo();

        // --- Deposit ---
        System.out.println("=== Deposit 2000 to Tareq ===");
        account1.deposit(2000);
        account1.showInfo();

        // --- Withdraw ---
        System.out.println("=== Withdraw 1500 from Afsana ===");
        account2.withdraw(1500);
        account2.showInfo();

        // --- Withdraw more than balance (should fail) ---
        System.out.println("=== Withdraw 5000 from Afsana (should fail) ===");
        account2.withdraw(5000);
        account2.showInfo();

        // --- Transfer ---
        System.out.println("=== Transfer 3000 from Tareq to Afsana ===");
        account1.transfer(3000, account2);
        System.out.println("After transfer:");
        account1.showInfo();
        account2.showInfo();
    }
}
