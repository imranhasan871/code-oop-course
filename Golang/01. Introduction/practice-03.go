/**
 * Practice 03: Generate a Voucher
 * Task: Read grocery items from a CSV file, calculate totals, and generate
 *       a formatted voucher. Save the voucher to a text file.
 *       If there are duplicate item IDs in the file, show an error and stop.
 *
 * How to run:
 *   go run practice-03.go
 *
 * Expected output on console:
 *   Item id      Qty    unit price    total
 *   item-937     12     230.5         2766.0
 *   ...
 *   Total        29     Grand total   4658.5
 *   VAT (5%)     232.925
 *   Net total    4425.575
 *
 * NOTE: The data file is at: ../../data/groceryitems.csv
 *       The voucher will be saved to: voucher.txt (in the same directory)
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/** GroceryItem holds the details for one grocery line item. */
type GroceryItem struct {
	ItemID    string
	Quantity  int
	UnitPrice float64
	Total     float64 // Calculated: Quantity × UnitPrice
}

func main() {
	inputFile := "../../data/groceryitems.csv"
	outputFile := "voucher.txt"

	// Step 1: Read items from file
	items, err := readGroceryItems(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	if len(items) == 0 {
		fmt.Println("No grocery items found in the file.")
		return
	}

	// Step 2: Check for duplicate item IDs before processing
	duplicates := findDuplicateItemIDs(items)
	if len(duplicates) > 0 {
		fmt.Println("Error: Cannot generate voucher. Duplicate item IDs found:")
		for _, id := range duplicates {
			fmt.Printf("  - %s\n", id)
		}
		fmt.Println("Please fix the data file and try again.")
		return
	}

	// Step 3: Calculate total for each item
	for i := range items {
		items[i].Total = float64(items[i].Quantity) * items[i].UnitPrice
	}

	// Step 4: Generate the formatted voucher string
	voucher := generateVoucher(items)

	// Step 5: Display voucher on console
	fmt.Print(voucher)

	// Step 6: Save voucher to a text file
	if err := saveToFile(outputFile, voucher); err != nil {
		fmt.Printf("Error saving voucher: %v\n", err)
		return
	}
	fmt.Printf("\nVoucher saved to: %s\n", outputFile)
}

/**
 * readGroceryItems reads grocery item data from a whitespace-separated file.
 * Each line: ItemID  Quantity  UnitPrice
 */
func readGroceryItems(fileName string) ([]GroceryItem, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []GroceryItem

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // skip blank lines
		}

		// strings.Fields splits by any whitespace (tab or spaces)
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue // skip malformed lines
		}

		itemID := parts[0]
		qty, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		unitPrice, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			continue
		}

		items = append(items, GroceryItem{
			ItemID:    itemID,
			Quantity:  qty,
			UnitPrice: unitPrice,
		})
	}

	return items, scanner.Err()
}

/**
 * findDuplicateItemIDs returns a slice of item IDs that appear more than once.
 */
func findDuplicateItemIDs(items []GroceryItem) []string {
	count := make(map[string]int)
	for _, item := range items {
		count[item.ItemID]++
	}

	var duplicates []string
	for id, n := range count {
		if n > 1 {
			duplicates = append(duplicates, id)
		}
	}
	return duplicates
}

/**
 * generateVoucher builds the formatted voucher as a string.
 */
func generateVoucher(items []GroceryItem) string {
	var sb strings.Builder

	// Header row
	sb.WriteString(fmt.Sprintf("%-12s\t%-6s\t%-12s\t%s\n",
		"Item id", "Qty", "unit price", "total"))

	var grandTotal float64
	var totalQty int

	// Item rows
	for _, item := range items {
		sb.WriteString(fmt.Sprintf("%-12s\t%-6d\t%-12.1f\t%.1f\n",
			item.ItemID, item.Quantity, item.UnitPrice, item.Total))
		grandTotal += item.Total
		totalQty += item.Quantity
	}

	// Summary rows
	vat := grandTotal * 0.05
	netTotal := grandTotal - vat

	sb.WriteString(fmt.Sprintf("\n%-12s\t%-6d\t%-12s\t%.1f\n",
		"Total", totalQty, "Grand total", grandTotal))
	sb.WriteString(fmt.Sprintf("VAT (5%%)\t%.3f\n", vat))
	sb.WriteString(fmt.Sprintf("Net total\t%.3f\n", netTotal))

	return sb.String()
}

/**
 * saveToFile writes the given content string to a text file.
 */
func saveToFile(fileName string, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
