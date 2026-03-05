/**
 * Practice 09: Bank Account Collection
 *
 * This is an extended practice of Practice 04. In this exercise, you will
 * create a list of Bank Accounts, perform various transactions (such as
 * withdrawals and deposits) across multiple accounts, and calculate the
 * total balance of all the accounts owned by the Bank.
 *
 * Key Concepts:
 *   - Reusing the BankAccount class from Practice 04
 *   - ArrayList of objects (object collection)
 *   - Iterating over a collection to aggregate data
 *   - Business operations across multiple objects
 *
 * Course: Professional OOP — by Zohirul Alam Tiemoon
 */

import java.util.ArrayList;
import java.util.List;

public class Practice09 {

    /** Represents a bank account with basic operations. */
    static class BankAccount {
        private String accountNumber;
        private String accountName;
        private double balance;

        public BankAccount(String accountNumber, String accountName, double balance) {
            this.accountNumber = accountNumber;
            this.accountName = accountName;
            this.balance = balance;
        }

        public String getAccountNumber() { return accountNumber; }
        public String getAccountName() { return accountName; }
        public double getBalance() { return balance; }

        public void deposit(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Deposit amount must be positive.");
                return;
            }
            balance += amount;
            System.out.printf("  [OK] Deposited %.2f to %s. New balance: %.2f%n",
                    amount, accountName, balance);
        }

        public void withdraw(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Withdrawal amount must be positive.");
                return;
            }
            if (amount > balance) {
                System.out.printf("  [Error] Insufficient balance in %s. Available: %.2f%n",
                        accountName, balance);
                return;
            }
            balance -= amount;
            System.out.printf("  [OK] Withdrew %.2f from %s. New balance: %.2f%n",
                    amount, accountName, balance);
        }

        public void transfer(BankAccount to, double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Transfer amount must be positive.");
                return;
            }
            if (amount > balance) {
                System.out.printf("  [Error] Insufficient balance in %s. Available: %.2f%n",
                        accountName, balance);
                return;
            }
            balance -= amount;
            to.balance += amount;
            System.out.printf("  [OK] Transferred %.2f from %s to %s%n",
                    amount, accountName, to.accountName);
        }

        public void printInfo() {
            System.out.printf("  %-10s | %-10s | Balance: %10.2f%n",
                    accountNumber, accountName, balance);
        }
    }

    /** Manages a collection of bank accounts. */
    static class Bank {
        private String bankName;
        private List<BankAccount> accounts;

        public Bank(String bankName) {
            this.bankName = bankName;
            this.accounts = new ArrayList<>();
        }

        public void addAccount(BankAccount account) {
            accounts.add(account);
            System.out.printf("  [OK] Added account %s (%s) to %s.%n",
                    account.getAccountNumber(), account.getAccountName(), bankName);
        }

        public void showAllAccounts() {
            System.out.printf("  %s — All Accounts (%d):%n", bankName, accounts.size());
            for (BankAccount acc : accounts) {
                acc.printInfo();
            }
        }

        public double getTotalBalance() {
            double total = 0;
            for (BankAccount acc : accounts) {
                total += acc.getBalance();
            }
            return total;
        }

        public void printTotalBalance() {
            System.out.printf("  Total balance of %s: %.2f%n", bankName, getTotalBalance());
        }
    }

    public static void main(String[] args) {
        System.out.println("=== Practice 09: Bank Account Collection ===");
        System.out.println();

        // Create bank
        Bank bank = new Bank("City Bank");

        // Create accounts
        BankAccount acc1 = new BankAccount("ACC-1001", "Imtiaz", 50000);
        BankAccount acc2 = new BankAccount("ACC-1002", "Faria", 30000);
        BankAccount acc3 = new BankAccount("ACC-1003", "Rafi", 45000);
        BankAccount acc4 = new BankAccount("ACC-1004", "Salma", 60000);

        // Add accounts to bank
        System.out.println("--- Add Accounts ---");
        bank.addAccount(acc1);
        bank.addAccount(acc2);
        bank.addAccount(acc3);
        bank.addAccount(acc4);
        System.out.println();

        // Show all accounts
        System.out.println("--- All Accounts ---");
        bank.showAllAccounts();
        System.out.println();

        // Total balance before transactions
        System.out.println("--- Total Balance (Before) ---");
        bank.printTotalBalance();
        System.out.println();

        // Perform transactions
        System.out.println("--- Transactions ---");
        acc1.deposit(10000);
        acc2.withdraw(5000);
        acc3.transfer(acc4, 15000);
        acc1.transfer(acc2, 20000);
        System.out.println();

        // Show all accounts after transactions
        System.out.println("--- All Accounts (After Transactions) ---");
        bank.showAllAccounts();
        System.out.println();

        // Total balance after transactions (should be same — money moves within bank)
        System.out.println("--- Total Balance (After) ---");
        bank.printTotalBalance();
    }
}
