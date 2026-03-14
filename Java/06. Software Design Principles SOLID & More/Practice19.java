/**
 * Practice 19: Applying DRY (Don't Repeat Yourself)
 * Task: Refactor duplicated discount and invoice logic into reusable components.
 *
 * How to compile and run:
 *   javac Practice19.java
 *   java Practice19
 */

import java.util.ArrayList;
import java.util.List;

public class Practice19 {

    static class OrderItem {
        String name;
        double unitPrice;
        int quantity;

        OrderItem(String name, double unitPrice, int quantity) {
            this.name = name;
            this.unitPrice = unitPrice;
            this.quantity = quantity;
        }

        double lineTotal() {
            return unitPrice * quantity;
        }
    }

    static class DiscountPolicy {
        static double calculate(double subtotal, String customerType) {
            String key = customerType.toLowerCase();
            double rate;
            if ("gold".equals(key)) {
                rate = 0.10;
            } else if ("silver".equals(key)) {
                rate = 0.05;
            } else {
                rate = 0.00;
            }
            return subtotal * rate;
        }
    }

    static class InvoiceSummary {
        double subtotal;
        double discount;
        double vat;
        double total;
    }

    static class InvoiceCalculator {
        private static final double VAT_RATE = 0.05;

        InvoiceSummary summarize(List<OrderItem> items, String customerType) {
            double subtotal = 0.0;
            for (OrderItem item : items) {
                subtotal += item.lineTotal();
            }

            double discount = DiscountPolicy.calculate(subtotal, customerType);
            double taxable = subtotal - discount;
            double vat = taxable * VAT_RATE;

            InvoiceSummary summary = new InvoiceSummary();
            summary.subtotal = subtotal;
            summary.discount = discount;
            summary.vat = vat;
            summary.total = taxable + vat;
            return summary;
        }
    }

    static void printInvoice(String customerName, String customerType, List<OrderItem> items) {
        InvoiceCalculator calculator = new InvoiceCalculator();
        InvoiceSummary summary = calculator.summarize(items, customerType);

        System.out.println("Customer: " + customerName + " (" + customerType + ")");
        for (OrderItem item : items) {
            System.out.printf("  - %-10s x%-2d @ $%6.2f = $%7.2f%n",
                item.name, item.quantity, item.unitPrice, item.lineTotal());
        }

        System.out.printf("Subtotal : $%.2f%n", summary.subtotal);
        System.out.printf("Discount : $%.2f%n", summary.discount);
        System.out.printf("VAT (5%%) : $%.2f%n", summary.vat);
        System.out.printf("Total    : $%.2f%n", summary.total);
        System.out.println();
    }

    public static void main(String[] args) {
        List<OrderItem> items = new ArrayList<>();
        items.add(new OrderItem("Keyboard", 45.0, 2));
        items.add(new OrderItem("Mouse", 20.0, 1));
        items.add(new OrderItem("USBC", 12.5, 3));

        printInvoice("Samia", "regular", items);
        printInvoice("Afsana", "silver", items);
        printInvoice("Hasan", "gold", items);
    }
}
