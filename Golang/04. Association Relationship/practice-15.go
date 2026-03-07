/**
 * Practice 15: Restaurant Order Management System
 * Task: Customers place orders with multiple menu items.
 *       Calculate total bill. Kitchen updates order status.
 *       Payment completion after confirmation.
 *
 * How to run:
 *   go run practice-15.go
 *
 * Key Concepts:
 *   - Association between Customer, Order, MenuItem, and Kitchen
 *   - Order status management (Pending, Cooking, Served)
 *   - Bill calculation
 */

package main

import "fmt"

/** MenuItem — a single item on the restaurant menu. */
type MenuItem15 struct {
	ItemId string
	Name   string
	Price  float64
}

func (m *MenuItem15) ShowInfo() {
	fmt.Printf("  %-8s | %-25s | Price: %.2f\n", m.ItemId, m.Name, m.Price)
}

/** OrderItem holds a menu item and its quantity. */
type OrderItem15 struct {
	Item     *MenuItem15
	Quantity int
}

/** Order placed by a customer. */
type Order15 struct {
	OrderId      string
	CustomerName string
	Items        []OrderItem15
	Status       string
	IsPaid       bool
}

func NewOrder15(orderId, customerName string) *Order15 {
	return &Order15{OrderId: orderId, CustomerName: customerName, Status: "Pending"}
}

func (o *Order15) AddItem(item *MenuItem15, quantity int) {
	if o.Status != "Pending" {
		fmt.Printf("  [Error] Cannot modify order %s — status is '%s'.\n", o.OrderId, o.Status)
		return
	}
	o.Items = append(o.Items, OrderItem15{item, quantity})
	fmt.Printf("  [OK] Added %dx %s to order %s.\n", quantity, item.Name, o.OrderId)
}

func (o *Order15) CalculateTotal() float64 {
	total := 0.0
	for _, oi := range o.Items {
		total += oi.Item.Price * float64(oi.Quantity)
	}
	return total
}

func (o *Order15) ShowInfo() {
	paid := "No"
	if o.IsPaid {
		paid = "Yes"
	}
	fmt.Println("  Order ID  :", o.OrderId)
	fmt.Println("  Customer  :", o.CustomerName)
	fmt.Println("  Status    :", o.Status)
	fmt.Println("  Paid      :", paid)
	if len(o.Items) > 0 {
		fmt.Printf("  %-25s | %3s | %10s | %10s\n", "Item", "Qty", "Unit Price", "Total")
		fmt.Println("  ------------------------------------------------------------")
		for _, oi := range o.Items {
			fmt.Printf("  %-25s | %3d | %10.2f | %10.2f\n",
				oi.Item.Name, oi.Quantity, oi.Item.Price, oi.Item.Price*float64(oi.Quantity))
		}
		fmt.Println("  ------------------------------------------------------------")
		fmt.Printf("  %-25s | %3s | %10s | %10.2f\n", "Total Bill", "", "", o.CalculateTotal())
	}
	fmt.Println()
}

/** Kitchen processes orders and updates their status. */
type Kitchen15 struct {
	Orders []*Order15
}

func (k *Kitchen15) ReceiveOrder(order *Order15) {
	k.Orders = append(k.Orders, order)
	fmt.Printf("  [OK] Kitchen received order %s from %s.\n", order.OrderId, order.CustomerName)
}

func (k *Kitchen15) findOrder(orderId string) *Order15 {
	for _, o := range k.Orders {
		if o.OrderId == orderId {
			return o
		}
	}
	fmt.Printf("  [Error] Order %s not found in kitchen.\n", orderId)
	return nil
}

func (k *Kitchen15) StartCooking(orderId string) {
	order := k.findOrder(orderId)
	if order == nil {
		return
	}
	if order.Status != "Pending" {
		fmt.Printf("  [Error] Order %s is already '%s'.\n", orderId, order.Status)
		return
	}
	order.Status = "Cooking"
	fmt.Printf("  [OK] Order %s is now being cooked.\n", orderId)
}

func (k *Kitchen15) ServeOrder(orderId string) {
	order := k.findOrder(orderId)
	if order == nil {
		return
	}
	if order.Status != "Cooking" {
		fmt.Printf("  [Error] Order %s must be 'Cooking' before serving (currently: '%s').\n",
			orderId, order.Status)
		return
	}
	order.Status = "Served"
	fmt.Printf("  [OK] Order %s has been served.\n", orderId)
}

func (k *Kitchen15) ShowQueue() {
	if len(k.Orders) == 0 {
		fmt.Println("  Kitchen queue is empty.")
		return
	}
	fmt.Printf("  %-10s | %-15s | %-10s | %10s\n", "Order ID", "Customer", "Status", "Total")
	fmt.Println("  -------------------------------------------------------")
	for _, o := range k.Orders {
		fmt.Printf("  %-10s | %-15s | %-10s | %10.2f\n",
			o.OrderId, o.CustomerName, o.Status, o.CalculateTotal())
	}
	fmt.Println()
}

/** Restaurant with menu, kitchen, and order management. */
type Restaurant15 struct {
	Name         string
	Menu         []*MenuItem15
	Kitchen      *Kitchen15
	OrderCounter int
}

func NewRestaurant15(name string) *Restaurant15 {
	return &Restaurant15{Name: name, Kitchen: &Kitchen15{}}
}

func (r *Restaurant15) AddMenuItem(item *MenuItem15) {
	r.Menu = append(r.Menu, item)
}

func (r *Restaurant15) PlaceOrder(customerName string, items []*MenuItem15, quantities []int) *Order15 {
	r.OrderCounter++
	order := NewOrder15(fmt.Sprintf("ORD-%03d", r.OrderCounter), customerName)
	for i := range items {
		order.AddItem(items[i], quantities[i])
	}
	r.Kitchen.ReceiveOrder(order)
	return order
}

func (r *Restaurant15) CompletePayment(order *Order15) {
	if order.Status != "Served" {
		fmt.Printf("  [Error] Order %s is not served yet (status: '%s').\n", order.OrderId, order.Status)
		return
	}
	order.IsPaid = true
	fmt.Printf("  [OK] Payment of %.2f completed for order %s by %s.\n",
		order.CalculateTotal(), order.OrderId, order.CustomerName)
}

func (r *Restaurant15) ShowMenu() {
	fmt.Printf("  === %s Menu ===\n", r.Name)
	for _, item := range r.Menu {
		item.ShowInfo()
	}
	fmt.Println()
}

func main() {
	restaurant := NewRestaurant15("Spice Garden")
	biryani := &MenuItem15{"M-01", "Chicken Biryani", 350}
	kebab := &MenuItem15{"M-02", "Seekh Kebab", 180}
	naan := &MenuItem15{"M-03", "Butter Naan", 60}
	lassi := &MenuItem15{"M-04", "Mango Lassi", 120}
	dessert := &MenuItem15{"M-05", "Gulab Jamun", 80}

	for _, item := range []*MenuItem15{biryani, kebab, naan, lassi, dessert} {
		restaurant.AddMenuItem(item)
	}
	restaurant.ShowMenu()

	fmt.Println("=== Placing Orders ===")
	order1 := restaurant.PlaceOrder("Tareq",
		[]*MenuItem15{biryani, kebab, lassi}, []int{2, 3, 2})
	fmt.Println()
	order2 := restaurant.PlaceOrder("Afsana",
		[]*MenuItem15{naan, biryani, dessert}, []int{4, 1, 2})
	fmt.Println()

	fmt.Println("=== Kitchen Queue ===")
	restaurant.Kitchen.ShowQueue()

	fmt.Println("=== Kitchen Processing ===")
	restaurant.Kitchen.StartCooking("ORD-001")
	restaurant.Kitchen.StartCooking("ORD-002")
	fmt.Println()

	restaurant.Kitchen.ServeOrder("ORD-001")
	fmt.Println()

	fmt.Println("=== Kitchen Queue After Processing ===")
	restaurant.Kitchen.ShowQueue()

	fmt.Println("=== Payment ===")
	restaurant.CompletePayment(order1)
	restaurant.CompletePayment(order2) // Should fail
	fmt.Println()

	restaurant.Kitchen.ServeOrder("ORD-002")
	restaurant.CompletePayment(order2)
	fmt.Println()

	fmt.Println("=== Final Order Details ===")
	order1.ShowInfo()
	order2.ShowInfo()
}
