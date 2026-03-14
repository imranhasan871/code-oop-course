/**
 * Practice 19: Applying DRY (Don't Repeat Yourself)
 * Task: Refactor duplicated discount and invoice logic into reusable components.
 *
 * How to run:
 *   go run practice-19.go
 */

package main

import "fmt"

type OrderItem struct {
	Name      string
	UnitPrice float64
	Quantity  int
}

func (i *OrderItem) LineTotal() float64 {
	return i.UnitPrice * float64(i.Quantity)
}

type DiscountPolicy struct{}

func (d *DiscountPolicy) Calculate(subtotal float64, customerType string) float64 {
	rate := 0.0
	switch customerType {
	case "gold":
		rate = 0.10
	case "silver":
		rate = 0.05
	default:
		rate = 0.00
	}
	return subtotal * rate
}

type InvoiceSummary struct {
	Subtotal float64
	Discount float64
	Vat      float64
	Total    float64
}

type InvoiceCalculator struct {
	DiscountPolicy *DiscountPolicy
}

func (c *InvoiceCalculator) Summarize(items []OrderItem, customerType string) InvoiceSummary {
	subtotal := 0.0
	for _, item := range items {
		subtotal += item.LineTotal()
	}

	discount := c.DiscountPolicy.Calculate(subtotal, customerType)
	taxable := subtotal - discount
	vat := taxable * 0.05

	return InvoiceSummary{
		Subtotal: subtotal,
		Discount: discount,
		Vat:      vat,
		Total:    taxable + vat,
	}
}

func printInvoice(customerName string, customerType string, items []OrderItem) {
	calculator := &InvoiceCalculator{DiscountPolicy: &DiscountPolicy{}}
	summary := calculator.Summarize(items, customerType)

	fmt.Printf("Customer: %s (%s)\n", customerName, customerType)
	for _, item := range items {
		fmt.Printf("  - %-10s x%-2d @ $%6.2f = $%7.2f\n",
			item.Name, item.Quantity, item.UnitPrice, item.LineTotal())
	}

	fmt.Printf("Subtotal : $%.2f\n", summary.Subtotal)
	fmt.Printf("Discount : $%.2f\n", summary.Discount)
	fmt.Printf("VAT (5%%) : $%.2f\n", summary.Vat)
	fmt.Printf("Total    : $%.2f\n", summary.Total)
	fmt.Println()
}

func main() {
	items := []OrderItem{
		{Name: "Keyboard", UnitPrice: 45.0, Quantity: 2},
		{Name: "Mouse", UnitPrice: 20.0, Quantity: 1},
		{Name: "USBC", UnitPrice: 12.5, Quantity: 3},
	}

	printInvoice("Samia", "regular", items)
	printInvoice("Afsana", "silver", items)
	printInvoice("Hasan", "gold", items)
}
