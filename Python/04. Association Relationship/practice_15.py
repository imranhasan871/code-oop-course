"""
Practice 15: Restaurant Order Management System
Task: Customers place orders with multiple menu items.
      Calculate total bill. Kitchen updates order status.
      Payment completion after confirmation.

How to run:
  python practice_15.py

Key Concepts:
  - Association between Customer, Order, MenuItem, and Kitchen
  - Order status management (Pending, Cooking, Served)
  - Bill calculation
"""


class MenuItem:
    """A single item on the restaurant menu."""

    def __init__(self, item_id: str, name: str, price: float):
        """Creates a new MenuItem."""
        self.item_id = item_id
        self.name = name
        self.price = price

    def show_info(self):
        """Prints menu item details."""
        print(f"  {self.item_id:<8} | {self.name:<25} | Price: {self.price:.2f}")


class Order:
    """An order placed by a customer, containing multiple MenuItems."""

    def __init__(self, order_id: str, customer_name: str):
        """Creates a new Order."""
        self.order_id = order_id
        self.customer_name = customer_name
        self.items = []       # list of (MenuItem, quantity) tuples
        self.status = "Pending"
        self.is_paid = False

    def add_item(self, item: MenuItem, quantity: int = 1):
        """Adds a menu item to the order."""
        if self.status != "Pending":
            print(f"  [Error] Cannot modify order {self.order_id} — status is '{self.status}'.")
            return
        self.items.append((item, quantity))
        print(f"  [OK] Added {quantity}x {item.name} to order {self.order_id}.")

    def calculate_total(self) -> float:
        """Calculates the total bill for this order."""
        total = 0.0
        for item, qty in self.items:
            total += item.price * qty
        return total

    def show_info(self):
        """Prints order details."""
        print(f"  Order ID  : {self.order_id}")
        print(f"  Customer  : {self.customer_name}")
        print(f"  Status    : {self.status}")
        print(f"  Paid      : {'Yes' if self.is_paid else 'No'}")
        if self.items:
            print(f"  {'Item':<25} | {'Qty':>3} | {'Unit Price':>10} | {'Total':>10}")
            print("  " + "-" * 60)
            for item, qty in self.items:
                print(f"  {item.name:<25} | {qty:>3} | {item.price:>10.2f} | {item.price * qty:>10.2f}")
            print("  " + "-" * 60)
            print(f"  {'Total Bill':<25} | {'':>3} | {'':>10} | {self.calculate_total():>10.2f}")
        print()


class Kitchen:
    """Kitchen that processes orders and updates their status."""

    def __init__(self):
        """Creates a new Kitchen."""
        self.orders = []

    def receive_order(self, order: Order):
        """Receives an order into the kitchen queue."""
        self.orders.append(order)
        print(f"  [OK] Kitchen received order {order.order_id} from {order.customer_name}.")

    def start_cooking(self, order_id: str):
        """Updates order status to 'Cooking'."""
        order = self._find_order(order_id)
        if order is None:
            return
        if order.status != "Pending":
            print(f"  [Error] Order {order_id} is already '{order.status}'.")
            return
        order.status = "Cooking"
        print(f"  [OK] Order {order_id} is now being cooked.")

    def serve_order(self, order_id: str):
        """Updates order status to 'Served'."""
        order = self._find_order(order_id)
        if order is None:
            return
        if order.status != "Cooking":
            print(f"  [Error] Order {order_id} must be 'Cooking' before serving (currently: '{order.status}').")
            return
        order.status = "Served"
        print(f"  [OK] Order {order_id} has been served.")

    def _find_order(self, order_id: str):
        """Finds an order by ID."""
        for order in self.orders:
            if order.order_id == order_id:
                return order
        print(f"  [Error] Order {order_id} not found in kitchen.")
        return None

    def show_queue(self):
        """Prints all orders in the kitchen queue."""
        if not self.orders:
            print("  Kitchen queue is empty.")
            return
        print(f"  {'Order ID':<10} | {'Customer':<15} | {'Status':<10} | {'Total':>10}")
        print("  " + "-" * 55)
        for order in self.orders:
            print(f"  {order.order_id:<10} | {order.customer_name:<15} | "
                  f"{order.status:<10} | {order.calculate_total():>10.2f}")
        print()


class Restaurant:
    """Restaurant with menu, kitchen, and order management."""

    def __init__(self, name: str):
        """Creates a new Restaurant."""
        self.name = name
        self.menu = []
        self.kitchen = Kitchen()
        self.order_counter = 0

    def add_menu_item(self, item: MenuItem):
        """Adds an item to the restaurant menu."""
        self.menu.append(item)

    def place_order(self, customer_name: str, items: list) -> Order:
        """Creates a new order for the customer."""
        self.order_counter += 1
        order = Order(f"ORD-{self.order_counter:03d}", customer_name)
        for item, qty in items:
            order.add_item(item, qty)
        self.kitchen.receive_order(order)
        return order

    def complete_payment(self, order: Order):
        """Completes payment for a served order."""
        if order.status != "Served":
            print(f"  [Error] Order {order.order_id} is not served yet (status: '{order.status}').")
            return
        order.is_paid = True
        total = order.calculate_total()
        print(f"  [OK] Payment of {total:.2f} completed for order {order.order_id} "
              f"by {order.customer_name}.")

    def show_menu(self):
        """Prints the restaurant menu."""
        print(f"  === {self.name} Menu ===")
        for item in self.menu:
            item.show_info()
        print()


def main():
    # --- Setup restaurant and menu ---
    restaurant = Restaurant("Spice Garden")
    biryani = MenuItem("M-01", "Chicken Biryani", 350.00)
    kebab = MenuItem("M-02", "Seekh Kebab", 180.00)
    naan = MenuItem("M-03", "Butter Naan", 60.00)
    lassi = MenuItem("M-04", "Mango Lassi", 120.00)
    dessert = MenuItem("M-05", "Gulab Jamun", 80.00)

    for item in [biryani, kebab, naan, lassi, dessert]:
        restaurant.add_menu_item(item)

    restaurant.show_menu()

    # --- Place orders ---
    print("=== Placing Orders ===")
    order1 = restaurant.place_order("Tareq", [(biryani, 2), (kebab, 3), (lassi, 2)])
    print()
    order2 = restaurant.place_order("Afsana", [(naan, 4), (biryani, 1), (dessert, 2)])
    print()

    # --- Show kitchen queue ---
    print("=== Kitchen Queue ===")
    restaurant.kitchen.show_queue()

    # --- Kitchen processes orders ---
    print("=== Kitchen Processing ===")
    restaurant.kitchen.start_cooking("ORD-001")
    restaurant.kitchen.start_cooking("ORD-002")
    print()

    restaurant.kitchen.serve_order("ORD-001")
    print()

    print("=== Kitchen Queue After Processing ===")
    restaurant.kitchen.show_queue()

    # --- Payment ---
    print("=== Payment ===")
    restaurant.complete_payment(order1)
    restaurant.complete_payment(order2)  # Should fail — not served yet
    print()

    # --- Serve and pay order 2 ---
    restaurant.kitchen.serve_order("ORD-002")
    restaurant.complete_payment(order2)
    print()

    # --- Show final order details ---
    print("=== Final Order Details ===")
    order1.show_info()
    order2.show_info()


if __name__ == "__main__":
    main()
