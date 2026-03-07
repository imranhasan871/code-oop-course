/**
 * Practice 15: Restaurant Order Management System
 * Task: Customers place orders with multiple menu items.
 *       Calculate total bill. Kitchen updates order status.
 *       Payment completion after confirmation.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice15.cs
 *
 * Key Concepts:
 *   - Association between Customer, Order, MenuItem, and Kitchen
 *   - Order status management (Pending, Cooking, Served)
 *   - Bill calculation
 */

using System;
using System.Collections.Generic;

class Practice15
{
    class MenuItem
    {
        public string ItemId { get; }
        public string Name { get; }
        public double Price { get; }

        public MenuItem(string itemId, string name, double price)
        {
            ItemId = itemId;
            Name = name;
            Price = price;
        }

        public void ShowInfo()
        {
            Console.WriteLine($"  {ItemId,-8} | {Name,-25} | Price: {Price:F2}");
        }
    }

    class Order
    {
        public string OrderId { get; }
        public string CustomerName { get; }
        private List<MenuItem> items = new List<MenuItem>();
        private List<int> quantities = new List<int>();
        public string Status { get; set; } = "Pending";
        public bool IsPaid { get; set; } = false;

        public Order(string orderId, string customerName)
        {
            OrderId = orderId;
            CustomerName = customerName;
        }

        public void AddItem(MenuItem item, int quantity)
        {
            if (Status != "Pending")
            {
                Console.WriteLine($"  [Error] Cannot modify order {OrderId} — status is '{Status}'.");
                return;
            }
            items.Add(item);
            quantities.Add(quantity);
            Console.WriteLine($"  [OK] Added {quantity}x {item.Name} to order {OrderId}.");
        }

        public double CalculateTotal()
        {
            double total = 0;
            for (int i = 0; i < items.Count; i++)
                total += items[i].Price * quantities[i];
            return total;
        }

        public void ShowInfo()
        {
            Console.WriteLine($"  Order ID  : {OrderId}");
            Console.WriteLine($"  Customer  : {CustomerName}");
            Console.WriteLine($"  Status    : {Status}");
            Console.WriteLine($"  Paid      : {(IsPaid ? "Yes" : "No")}");
            if (items.Count > 0)
            {
                Console.WriteLine($"  {"Item",-25} | {"Qty",3} | {"Unit Price",10} | {"Total",10}");
                Console.WriteLine("  " + new string('-', 60));
                for (int i = 0; i < items.Count; i++)
                {
                    var item = items[i];
                    int qty = quantities[i];
                    Console.WriteLine($"  {item.Name,-25} | {qty,3} | {item.Price,10:F2} | {item.Price * qty,10:F2}");
                }
                Console.WriteLine("  " + new string('-', 60));
                Console.WriteLine($"  {"Total Bill",-25} | {"",3} | {"",10} | {CalculateTotal(),10:F2}");
            }
            Console.WriteLine();
        }
    }

    class Kitchen
    {
        private List<Order> orders = new List<Order>();

        public void ReceiveOrder(Order order)
        {
            orders.Add(order);
            Console.WriteLine($"  [OK] Kitchen received order {order.OrderId} from {order.CustomerName}.");
        }

        public void StartCooking(string orderId)
        {
            var order = FindOrder(orderId);
            if (order == null) return;
            if (order.Status != "Pending")
            {
                Console.WriteLine($"  [Error] Order {orderId} is already '{order.Status}'.");
                return;
            }
            order.Status = "Cooking";
            Console.WriteLine($"  [OK] Order {orderId} is now being cooked.");
        }

        public void ServeOrder(string orderId)
        {
            var order = FindOrder(orderId);
            if (order == null) return;
            if (order.Status != "Cooking")
            {
                Console.WriteLine($"  [Error] Order {orderId} must be 'Cooking' before serving (currently: '{order.Status}').");
                return;
            }
            order.Status = "Served";
            Console.WriteLine($"  [OK] Order {orderId} has been served.");
        }

        private Order FindOrder(string orderId)
        {
            foreach (var order in orders)
                if (order.OrderId == orderId) return order;
            Console.WriteLine($"  [Error] Order {orderId} not found in kitchen.");
            return null;
        }

        public void ShowQueue()
        {
            if (orders.Count == 0)
            {
                Console.WriteLine("  Kitchen queue is empty.");
                return;
            }
            Console.WriteLine($"  {"Order ID",-10} | {"Customer",-15} | {"Status",-10} | {"Total",10}");
            Console.WriteLine("  " + new string('-', 55));
            foreach (var order in orders)
                Console.WriteLine($"  {order.OrderId,-10} | {order.CustomerName,-15} | {order.Status,-10} | {order.CalculateTotal(),10:F2}");
            Console.WriteLine();
        }
    }

    class Restaurant
    {
        public string Name { get; }
        private List<MenuItem> menu = new List<MenuItem>();
        public Kitchen Kitchen { get; } = new Kitchen();
        private int orderCounter = 0;

        public Restaurant(string name) { Name = name; }

        public void AddMenuItem(MenuItem item) { menu.Add(item); }

        public Order PlaceOrder(string customerName, (MenuItem item, int qty)[] items)
        {
            orderCounter++;
            var order = new Order($"ORD-{orderCounter:D3}", customerName);
            foreach (var (item, qty) in items)
                order.AddItem(item, qty);
            Kitchen.ReceiveOrder(order);
            return order;
        }

        public void CompletePayment(Order order)
        {
            if (order.Status != "Served")
            {
                Console.WriteLine($"  [Error] Order {order.OrderId} is not served yet (status: '{order.Status}').");
                return;
            }
            order.IsPaid = true;
            Console.WriteLine($"  [OK] Payment of {order.CalculateTotal():F2} completed for order " +
                              $"{order.OrderId} by {order.CustomerName}.");
        }

        public void ShowMenu()
        {
            Console.WriteLine($"  === {Name} Menu ===");
            foreach (var item in menu)
                item.ShowInfo();
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        var restaurant = new Restaurant("Spice Garden");
        var biryani = new MenuItem("M-01", "Chicken Biryani", 350);
        var kebab = new MenuItem("M-02", "Seekh Kebab", 180);
        var naan = new MenuItem("M-03", "Butter Naan", 60);
        var lassi = new MenuItem("M-04", "Mango Lassi", 120);
        var dessert = new MenuItem("M-05", "Gulab Jamun", 80);

        foreach (var item in new[] { biryani, kebab, naan, lassi, dessert })
            restaurant.AddMenuItem(item);
        restaurant.ShowMenu();

        Console.WriteLine("=== Placing Orders ===");
        var order1 = restaurant.PlaceOrder("Tareq",
            new[] { (biryani, 2), (kebab, 3), (lassi, 2) });
        Console.WriteLine();
        var order2 = restaurant.PlaceOrder("Afsana",
            new[] { (naan, 4), (biryani, 1), (dessert, 2) });
        Console.WriteLine();

        Console.WriteLine("=== Kitchen Queue ===");
        restaurant.Kitchen.ShowQueue();

        Console.WriteLine("=== Kitchen Processing ===");
        restaurant.Kitchen.StartCooking("ORD-001");
        restaurant.Kitchen.StartCooking("ORD-002");
        Console.WriteLine();

        restaurant.Kitchen.ServeOrder("ORD-001");
        Console.WriteLine();

        Console.WriteLine("=== Kitchen Queue After Processing ===");
        restaurant.Kitchen.ShowQueue();

        Console.WriteLine("=== Payment ===");
        restaurant.CompletePayment(order1);
        restaurant.CompletePayment(order2); // Should fail
        Console.WriteLine();

        restaurant.Kitchen.ServeOrder("ORD-002");
        restaurant.CompletePayment(order2);
        Console.WriteLine();

        Console.WriteLine("=== Final Order Details ===");
        order1.ShowInfo();
        order2.ShowInfo();
    }
}
