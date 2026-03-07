/**
 * Practice 15: Restaurant Order Management System
 * Task: Customers place orders with multiple menu items.
 *       Calculate total bill. Kitchen updates order status.
 *       Payment completion after confirmation.
 *
 * How to compile and run:
 *   javac Practice15.java
 *   java Practice15
 *
 * Key Concepts:
 *   - Association between Customer, Order, MenuItem, and Kitchen
 *   - Order status management (Pending, Cooking, Served)
 *   - Bill calculation
 */

import java.util.ArrayList;
import java.util.List;

public class Practice15 {

    /** A single item on the restaurant menu. */
    static class MenuItem {
        private String itemId;
        private String name;
        private double price;

        public MenuItem(String itemId, String name, double price) {
            this.itemId = itemId;
            this.name = name;
            this.price = price;
        }

        public String getName() { return name; }
        public double getPrice() { return price; }

        public void showInfo() {
            System.out.printf("  %-8s | %-25s | Price: %.2f%n", itemId, name, price);
        }
    }

    /** An order placed by a customer, containing multiple MenuItems. */
    static class Order {
        private String orderId;
        private String customerName;
        private List<int[]> itemIndices = new ArrayList<>(); // index + quantity
        private List<MenuItem> items = new ArrayList<>();
        private List<Integer> quantities = new ArrayList<>();
        private String status = "Pending";
        private boolean isPaid = false;

        public Order(String orderId, String customerName) {
            this.orderId = orderId;
            this.customerName = customerName;
        }

        public String getOrderId() { return orderId; }
        public String getCustomerName() { return customerName; }
        public String getStatus() { return status; }
        public void setStatus(String status) { this.status = status; }
        public boolean getIsPaid() { return isPaid; }
        public void setIsPaid(boolean paid) { this.isPaid = paid; }

        public void addItem(MenuItem item, int quantity) {
            if (!status.equals("Pending")) {
                System.out.printf("  [Error] Cannot modify order %s — status is '%s'.%n", orderId, status);
                return;
            }
            items.add(item);
            quantities.add(quantity);
            System.out.printf("  [OK] Added %dx %s to order %s.%n", quantity, item.getName(), orderId);
        }

        public double calculateTotal() {
            double total = 0;
            for (int i = 0; i < items.size(); i++) {
                total += items.get(i).getPrice() * quantities.get(i);
            }
            return total;
        }

        public void showInfo() {
            System.out.println("  Order ID  : " + orderId);
            System.out.println("  Customer  : " + customerName);
            System.out.println("  Status    : " + status);
            System.out.println("  Paid      : " + (isPaid ? "Yes" : "No"));
            if (!items.isEmpty()) {
                System.out.printf("  %-25s | %3s | %10s | %10s%n", "Item", "Qty", "Unit Price", "Total");
                System.out.println("  " + "-".repeat(60));
                for (int i = 0; i < items.size(); i++) {
                    MenuItem item = items.get(i);
                    int qty = quantities.get(i);
                    System.out.printf("  %-25s | %3d | %10.2f | %10.2f%n",
                            item.getName(), qty, item.getPrice(), item.getPrice() * qty);
                }
                System.out.println("  " + "-".repeat(60));
                System.out.printf("  %-25s | %3s | %10s | %10.2f%n", "Total Bill", "", "", calculateTotal());
            }
            System.out.println();
        }
    }

    /** Kitchen that processes orders and updates their status. */
    static class Kitchen {
        private List<Order> orders = new ArrayList<>();

        public void receiveOrder(Order order) {
            orders.add(order);
            System.out.printf("  [OK] Kitchen received order %s from %s.%n",
                    order.getOrderId(), order.getCustomerName());
        }

        public void startCooking(String orderId) {
            Order order = findOrder(orderId);
            if (order == null) return;
            if (!order.getStatus().equals("Pending")) {
                System.out.printf("  [Error] Order %s is already '%s'.%n", orderId, order.getStatus());
                return;
            }
            order.setStatus("Cooking");
            System.out.printf("  [OK] Order %s is now being cooked.%n", orderId);
        }

        public void serveOrder(String orderId) {
            Order order = findOrder(orderId);
            if (order == null) return;
            if (!order.getStatus().equals("Cooking")) {
                System.out.printf("  [Error] Order %s must be 'Cooking' before serving (currently: '%s').%n",
                        orderId, order.getStatus());
                return;
            }
            order.setStatus("Served");
            System.out.printf("  [OK] Order %s has been served.%n", orderId);
        }

        private Order findOrder(String orderId) {
            for (Order order : orders) {
                if (order.getOrderId().equals(orderId)) return order;
            }
            System.out.printf("  [Error] Order %s not found in kitchen.%n", orderId);
            return null;
        }

        public void showQueue() {
            if (orders.isEmpty()) {
                System.out.println("  Kitchen queue is empty.");
                return;
            }
            System.out.printf("  %-10s | %-15s | %-10s | %10s%n", "Order ID", "Customer", "Status", "Total");
            System.out.println("  " + "-".repeat(55));
            for (Order order : orders) {
                System.out.printf("  %-10s | %-15s | %-10s | %10.2f%n",
                        order.getOrderId(), order.getCustomerName(),
                        order.getStatus(), order.calculateTotal());
            }
            System.out.println();
        }
    }

    /** Restaurant with menu, kitchen, and order management. */
    static class Restaurant {
        private String name;
        private List<MenuItem> menu = new ArrayList<>();
        private Kitchen kitchen = new Kitchen();
        private int orderCounter = 0;

        public Restaurant(String name) {
            this.name = name;
        }

        public void addMenuItem(MenuItem item) {
            menu.add(item);
        }

        public Order placeOrder(String customerName, MenuItem[] items, int[] quantities) {
            orderCounter++;
            Order order = new Order(String.format("ORD-%03d", orderCounter), customerName);
            for (int i = 0; i < items.length; i++) {
                order.addItem(items[i], quantities[i]);
            }
            kitchen.receiveOrder(order);
            return order;
        }

        public void completePayment(Order order) {
            if (!order.getStatus().equals("Served")) {
                System.out.printf("  [Error] Order %s is not served yet (status: '%s').%n",
                        order.getOrderId(), order.getStatus());
                return;
            }
            order.setIsPaid(true);
            System.out.printf("  [OK] Payment of %.2f completed for order %s by %s.%n",
                    order.calculateTotal(), order.getOrderId(), order.getCustomerName());
        }

        public void showMenu() {
            System.out.println("  === " + name + " Menu ===");
            for (MenuItem item : menu) {
                item.showInfo();
            }
            System.out.println();
        }

        public Kitchen getKitchen() { return kitchen; }
    }

    public static void main(String[] args) {
        Restaurant restaurant = new Restaurant("Spice Garden");
        MenuItem biryani = new MenuItem("M-01", "Chicken Biryani", 350);
        MenuItem kebab = new MenuItem("M-02", "Seekh Kebab", 180);
        MenuItem naan = new MenuItem("M-03", "Butter Naan", 60);
        MenuItem lassi = new MenuItem("M-04", "Mango Lassi", 120);
        MenuItem dessert = new MenuItem("M-05", "Gulab Jamun", 80);

        for (MenuItem item : new MenuItem[]{biryani, kebab, naan, lassi, dessert}) {
            restaurant.addMenuItem(item);
        }
        restaurant.showMenu();

        System.out.println("=== Placing Orders ===");
        Order order1 = restaurant.placeOrder("Tareq",
                new MenuItem[]{biryani, kebab, lassi}, new int[]{2, 3, 2});
        System.out.println();
        Order order2 = restaurant.placeOrder("Afsana",
                new MenuItem[]{naan, biryani, dessert}, new int[]{4, 1, 2});
        System.out.println();

        System.out.println("=== Kitchen Queue ===");
        restaurant.getKitchen().showQueue();

        System.out.println("=== Kitchen Processing ===");
        restaurant.getKitchen().startCooking("ORD-001");
        restaurant.getKitchen().startCooking("ORD-002");
        System.out.println();

        restaurant.getKitchen().serveOrder("ORD-001");
        System.out.println();

        System.out.println("=== Kitchen Queue After Processing ===");
        restaurant.getKitchen().showQueue();

        System.out.println("=== Payment ===");
        restaurant.completePayment(order1);
        restaurant.completePayment(order2); // Should fail
        System.out.println();

        restaurant.getKitchen().serveOrder("ORD-002");
        restaurant.completePayment(order2);
        System.out.println();

        System.out.println("=== Final Order Details ===");
        order1.showInfo();
        order2.showInfo();
    }
}
