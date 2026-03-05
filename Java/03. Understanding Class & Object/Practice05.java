/**
 * Practice 05: Credit Card
 * Task: Create a CreditCard class with spending limits.
 *       Max credit limit: 500,000. Cash withdrawal: daily limit 100,000,
 *       per-transaction limit 20,000. Bill payments: no per-txn limit.
 *
 * How to compile and run:
 *   javac Practice05.java
 *   java Practice05
 *
 * Key Concepts:
 *   - Class with constants and instance tracking
 *   - Multiple validation rules in methods
 *   - Business logic enforcement
 */

public class Practice05 {

    /** CreditCard represents a credit card with spending limits. */
    static class CreditCard {
        private static final double MAX_CREDIT_LIMIT = 500000.00;
        private static final double DAILY_CASH_LIMIT = 100000.00;
        private static final double PER_TXN_CASH_LIMIT = 20000.00;

        private String cardNumber;
        private String cardHolder;
        private double totalSpent;
        private double dailyCashWithdrawn;

        /** Creates a new CreditCard with the given details. */
        public CreditCard(String cardNumber, String cardHolder) {
            this.cardNumber = cardNumber;
            this.cardHolder = cardHolder;
            this.totalSpent = 0;
            this.dailyCashWithdrawn = 0;
        }

        /** Withdraws cash with per-transaction and daily limits. */
        public void withdrawCash(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Amount must be greater than 0.");
                return;
            }
            if (amount > PER_TXN_CASH_LIMIT) {
                System.out.printf("  [Error] Cash withdrawal exceeds per-transaction limit of %.2f (Requested: %.2f)%n",
                        PER_TXN_CASH_LIMIT, amount);
                return;
            }
            if (dailyCashWithdrawn + amount > DAILY_CASH_LIMIT) {
                System.out.printf("  [Error] Cash withdrawal exceeds daily limit of %.2f (Already withdrawn today: %.2f, Requested: %.2f)%n",
                        DAILY_CASH_LIMIT, dailyCashWithdrawn, amount);
                return;
            }
            if (totalSpent + amount > MAX_CREDIT_LIMIT) {
                System.out.printf("  [Error] Total spending would exceed credit limit of %.2f (Already spent: %.2f, Requested: %.2f)%n",
                        MAX_CREDIT_LIMIT, totalSpent, amount);
                return;
            }
            totalSpent += amount;
            dailyCashWithdrawn += amount;
            System.out.printf("  [OK] Cash withdrawn: %.2f%n", amount);
        }

        /** Pays a bill amount. No per-transaction limit, but total must stay within max limit. */
        public void payBill(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Amount must be greater than 0.");
                return;
            }
            if (totalSpent + amount > MAX_CREDIT_LIMIT) {
                System.out.printf("  [Error] Bill payment would exceed credit limit of %.2f (Already spent: %.2f, Requested: %.2f)%n",
                        MAX_CREDIT_LIMIT, totalSpent, amount);
                return;
            }
            totalSpent += amount;
            System.out.printf("  [OK] Bill paid: %.2f%n", amount);
        }

        /** Prints the card details and current spending. */
        public void showInfo() {
            System.out.printf("  Card Number          : %s%n", cardNumber);
            System.out.printf("  Card Holder          : %s%n", cardHolder);
            System.out.printf("  Credit Limit         : %.2f%n", MAX_CREDIT_LIMIT);
            System.out.printf("  Total Spent          : %.2f%n", totalSpent);
            System.out.printf("  Available Limit      : %.2f%n", MAX_CREDIT_LIMIT - totalSpent);
            System.out.printf("  Daily Cash Withdrawn : %.2f / %.2f%n", dailyCashWithdrawn, DAILY_CASH_LIMIT);
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // --- Create a credit card ---
        CreditCard card = new CreditCard("4532-1234-5678-9012", "Tareq");

        System.out.println("=== Initial Card Info ===");
        card.showInfo();

        // --- Valid cash withdrawal ---
        System.out.println("=== Withdraw Cash 15,000 ===");
        card.withdrawCash(15000);
        card.showInfo();

        // --- Exceed per-transaction limit ---
        System.out.println("=== Withdraw Cash 25,000 (exceeds per-txn limit) ===");
        card.withdrawCash(25000);
        card.showInfo();

        // --- Multiple valid withdrawals to approach daily limit ---
        System.out.println("=== Withdraw Cash 20,000 x4 (should hit daily limit) ===");
        card.withdrawCash(20000);
        card.withdrawCash(20000);
        card.withdrawCash(20000);
        card.withdrawCash(20000);
        card.showInfo();

        // --- Bill payment ---
        System.out.println("=== Pay Bill 200,000 ===");
        card.payBill(200000);
        card.showInfo();

        // --- Try to exceed total credit limit ---
        System.out.println("=== Pay Bill 300,000 (should exceed total limit) ===");
        card.payBill(300000);
        card.showInfo();
    }
}
