/**
 * Practice 09: Bank Account Collection
 * Task: Manage a list of BankAccount objects. Perform transactions,
 *       calculate total balance, and find accounts with highest/lowest balance.
 *
 * How to compile and run:
 *   javac Practice09.java
 *   java Practice09
 *
 * Key Concepts:
 *   - ArrayList of objects (collection of objects)
 *   - Iterating over a collection to compute aggregates
 *   - Finding min/max in a collection of objects
 */

import java.util.ArrayList;
import java.util.List;

public class Practice09 {

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

        /** Deposits the given amount into the account. */
        public void deposit(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Deposit amount must be greater than 0.");
                return;
            }
            balance += amount;
            System.out.printf("  [OK] Deposited %.2f to %s%n", amount, accountName);
        }

        /** Withdraws the given amount from the account. */
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

        /** Prints the account details in a table row format. */
        public void showInfo() {
            System.out.printf("  %-10s | %-12s | Balance: %10.2f%n", accountNumber, accountName, balance);
        }
    }

    /** Calculates the total balance of all accounts. */
    public static double calculateTotalBalance(List<BankAccount> accounts) {
        double total = 0;
        for (BankAccount acc : accounts) {
            total += acc.balance;
        }
        return total;
    }

    /** Finds the account with the highest balance. */
    public static BankAccount findHighestBalance(List<BankAccount> accounts) {
        BankAccount highest = accounts.get(0);
        for (int i = 1; i < accounts.size(); i++) {
            if (accounts.get(i).balance > highest.balance) {
                highest = accounts.get(i);
            }
        }
        return highest;
    }

    /** Finds the account with the lowest balance. */
    public static BankAccount findLowestBalance(List<BankAccount> accounts) {
        BankAccount lowest = accounts.get(0);
        for (int i = 1; i < accounts.size(); i++) {
            if (accounts.get(i).balance < lowest.balance) {
                lowest = accounts.get(i);
            }
        }
        return lowest;
    }

    /** Prints all accounts in a formatted table. */
    public static void showAllAccounts(List<BankAccount> accounts) {
        System.out.printf("  %-10s | %-12s | %s%n", "Account No", "Name", "Balance");
        System.out.println("  -------------------------------------------");
        for (BankAccount acc : accounts) {
            acc.showInfo();
        }
        System.out.println();
    }

    public static void main(String[] args) {
        // --- Create a list of bank accounts ---
        List<BankAccount> accounts = new ArrayList<>();
        accounts.add(new BankAccount("ACC-1001", "Tareq", 15000));
        accounts.add(new BankAccount("ACC-1002", "Afsana", 22000));
        accounts.add(new BankAccount("ACC-1003", "Imtiaz", 8500));
        accounts.add(new BankAccount("ACC-1004", "Pulok", 31000));
        accounts.add(new BankAccount("ACC-1005", "Samia", 12000));

        System.out.println("=== All Accounts (Initial) ===");
        showAllAccounts(accounts);

        // --- Perform transactions ---
        System.out.println("=== Performing Transactions ===");
        accounts.get(0).deposit(5000);   // Tareq deposits 5000
        accounts.get(1).withdraw(3000);  // Afsana withdraws 3000
        accounts.get(2).deposit(1500);   // Imtiaz deposits 1500
        accounts.get(3).withdraw(10000); // Pulok withdraws 10000
        accounts.get(4).deposit(8000);   // Samia deposits 8000
        System.out.println();

        System.out.println("=== All Accounts (After Transactions) ===");
        showAllAccounts(accounts);

        // --- Calculate total balance ---
        double total = calculateTotalBalance(accounts);
        System.out.printf("  Total Balance of All Accounts: %.2f%n%n", total);

        // --- Find highest and lowest balance ---
        BankAccount highest = findHighestBalance(accounts);
        BankAccount lowest = findLowestBalance(accounts);

        System.out.println("=== Highest Balance ===");
        highest.showInfo();
        System.out.println();

        System.out.println("=== Lowest Balance ===");
        lowest.showInfo();
    }
}
