/**
 * Practice 18: Applying DIP (Dependency Inversion Principle)
 * Task: Build a payment system where PaymentService depends on abstraction,
 *       not concrete payment providers.
 *
 * How to compile and run:
 *   javac Practice18.java
 *   java Practice18
 */

import java.util.ArrayList;
import java.util.List;

public class Practice18 {

    interface PaymentMethod {
        String pay(double amount);
    }

    static class CreditCardPayment implements PaymentMethod {
        public String pay(double amount) {
            return String.format("Paid $%.2f using Credit Card", amount);
        }
    }

    static class BkashPayment implements PaymentMethod {
        public String pay(double amount) {
            return String.format("Paid $%.2f using bKash", amount);
        }
    }

    static class PaypalPayment implements PaymentMethod {
        public String pay(double amount) {
            return String.format("Paid $%.2f using PayPal", amount);
        }
    }

    static class BankTransferPayment implements PaymentMethod {
        public String pay(double amount) {
            return String.format("Paid $%.2f using Bank Transfer", amount);
        }
    }

    static class PaymentService {
        void processPayment(PaymentMethod paymentMethod, double amount) {
            if (amount <= 0) {
                throw new IllegalArgumentException("Amount must be greater than 0.");
            }
            System.out.println("[PaymentService] " + paymentMethod.pay(amount));
        }
    }

    public static void main(String[] args) {
        PaymentService service = new PaymentService();

        List<PaymentMethod> methods = new ArrayList<>();
        methods.add(new CreditCardPayment());
        methods.add(new BkashPayment());
        methods.add(new PaypalPayment());
        methods.add(new BankTransferPayment());

        for (PaymentMethod method : methods) {
            service.processPayment(method, 1500.0);
        }
    }
}
