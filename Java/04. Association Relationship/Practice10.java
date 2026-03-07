/**
 * Practice 10: 1-1 Association — Customer & Credit Card
 * Task: Manage customer profiles with associated credit cards.
 *       Card number must be digits only. Expiration date validation.
 *       Credit limit 500K. Track available credit. Age minimum 18.
 *
 * How to compile and run:
 *   javac Practice10.java
 *   java Practice10
 *
 * Key Concepts:
 *   - 1-1 Association (Customer has a CreditCard)
 *   - Input validation (digit-only card number, age check)
 *   - Date-based expiration validation
 */

import java.time.LocalDate;
import java.time.Period;

public class Practice10 {

    /** CreditCard with card number, expiration date, and credit limit. */
    static class CreditCard {
        private static final double MAX_CREDIT_LIMIT = 500000.00;

        private String cardNumber;
        private LocalDate expirationDate;
        private double creditLimit;
        private double totalSpent;

        /** Creates a new CreditCard. Card number must contain only digits. */
        public CreditCard(String cardNumber, LocalDate expirationDate, double creditLimit) {
            if (!cardNumber.matches("\\d+")) {
                throw new IllegalArgumentException("Card number must contain only digits, got: '" + cardNumber + "'");
            }
            if (creditLimit > MAX_CREDIT_LIMIT) {
                throw new IllegalArgumentException("Credit limit cannot exceed " + MAX_CREDIT_LIMIT);
            }
            this.cardNumber = cardNumber;
            this.expirationDate = expirationDate;
            this.creditLimit = creditLimit;
            this.totalSpent = 0;
        }

        public CreditCard(String cardNumber, LocalDate expirationDate) {
            this(cardNumber, expirationDate, MAX_CREDIT_LIMIT);
        }

        public boolean isValid() {
            return !LocalDate.now().isAfter(expirationDate);
        }

        public double availableCredit() {
            return creditLimit - totalSpent;
        }

        public double outstandingBalance() {
            return totalSpent;
        }

        public void makePurchase(double amount) {
            if (amount <= 0) {
                System.out.println("  [Error] Purchase amount must be greater than 0.");
                return;
            }
            if (!isValid()) {
                System.out.println("  [Error] Credit card has expired. Cannot make purchase.");
                return;
            }
            if (amount > availableCredit()) {
                System.out.printf("  [Error] Purchase of %.2f exceeds available credit (%.2f).%n",
                        amount, availableCredit());
                return;
            }
            totalSpent += amount;
            System.out.printf("  [OK] Purchase of %.2f successful.%n", amount);
        }

        public void showInfo() {
            String status = isValid() ? "Valid" : "Expired";
            System.out.println("  Card Number        : " + cardNumber);
            System.out.println("  Expiration Date    : " + expirationDate);
            System.out.println("  Status             : " + status);
            System.out.printf("  Credit Limit       : %.2f%n", creditLimit);
            System.out.printf("  Total Spent        : %.2f%n", totalSpent);
            System.out.printf("  Available Credit   : %.2f%n", availableCredit());
            System.out.printf("  Outstanding Balance: %.2f%n", outstandingBalance());
        }
    }

    /** Customer with personal info and an associated CreditCard (1-1). */
    static class Customer {
        private static final int MIN_AGE = 18;

        private String name;
        private LocalDate dateOfBirth;
        private CreditCard creditCard;

        /** Creates a Customer. Must be at least 18 years old. */
        public Customer(String name, LocalDate dateOfBirth, CreditCard creditCard) {
            int age = calculateAge(dateOfBirth);
            if (age < MIN_AGE) {
                throw new IllegalArgumentException(
                        "Customer must be at least " + MIN_AGE + " years old. " + name + " is " + age + " years old.");
            }
            this.name = name;
            this.dateOfBirth = dateOfBirth;
            this.creditCard = creditCard;
        }

        private static int calculateAge(LocalDate dob) {
            return Period.between(dob, LocalDate.now()).getYears();
        }

        public int getAge() {
            return calculateAge(dateOfBirth);
        }

        public void makePurchase(double amount) {
            System.out.printf("  %s making purchase of %.2f...%n", name, amount);
            creditCard.makePurchase(amount);
        }

        public void showInfo() {
            System.out.println("  Customer Name      : " + name);
            System.out.println("  Date of Birth      : " + dateOfBirth);
            System.out.println("  Age                : " + getAge());
            creditCard.showInfo();
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // --- Create credit cards ---
        CreditCard card1 = new CreditCard("4532123456789012", LocalDate.of(2028, 12, 31));
        CreditCard card2 = new CreditCard("5678901234567890", LocalDate.of(2025, 6, 15)); // Expired

        // --- Create customers ---
        Customer customer1 = new Customer("Tareq", LocalDate.of(1990, 5, 15), card1);
        Customer customer2 = new Customer("Afsana", LocalDate.of(1995, 8, 20), card2);

        System.out.println("=== Customer Info ===");
        customer1.showInfo();
        customer2.showInfo();

        // --- Make purchases ---
        System.out.println("=== Purchases ===");
        customer1.makePurchase(150000);
        customer1.makePurchase(200000);
        customer1.makePurchase(200000); // Should fail
        System.out.println();

        customer2.makePurchase(50000); // Should fail — expired
        System.out.println();

        System.out.println("=== After Purchases ===");
        customer1.showInfo();

        // --- Invalid card number ---
        System.out.println("=== Invalid Card Number ===");
        try {
            new CreditCard("ABCD-1234", LocalDate.of(2028, 1, 1));
        } catch (IllegalArgumentException e) {
            System.out.println("  [Error] " + e.getMessage());
        }
        System.out.println();

        // --- Underage customer ---
        System.out.println("=== Underage Customer ===");
        try {
            CreditCard youngCard = new CreditCard("1111222233334444", LocalDate.of(2030, 1, 1));
            new Customer("Junior", LocalDate.of(2015, 1, 1), youngCard);
        } catch (IllegalArgumentException e) {
            System.out.println("  [Error] " + e.getMessage());
        }
    }
}
