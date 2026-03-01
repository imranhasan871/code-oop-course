/**
 * Practice 03: Generate a Voucher
 * Task: Read grocery items from a CSV file, calculate totals, and generate
 *       a formatted voucher. Save the voucher to a text file.
 *       If there are duplicate item IDs in the file, show an error and stop.
 *
 * How to compile and run:
 *   javac Practice03.java
 *   java Practice03
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

import java.io.*;
import java.util.*;

public class Practice03 {

    /**
     * Inner class: represents one grocery line item.
     */
    static class GroceryItem {
        String itemID;
        int quantity;
        double unitPrice;
        double total; // Calculated: quantity × unitPrice

        GroceryItem(String itemID, int quantity, double unitPrice) {
            this.itemID = itemID;
            this.quantity = quantity;
            this.unitPrice = unitPrice;
            this.total = quantity * unitPrice;
        }
    }

    public static void main(String[] args) {
        String inputFile = "../../data/groceryitems.csv";
        String outputFile = "voucher.txt";

        // Step 1: Read items from file
        List<GroceryItem> items = readGroceryItems(inputFile);

        if (items == null || items.isEmpty()) {
            System.out.println("No grocery items found in the file.");
            return;
        }

        // Step 2: Check for duplicate item IDs before processing
        List<String> duplicates = findDuplicateItemIDs(items);
        if (!duplicates.isEmpty()) {
            System.out.println("Error: Cannot generate voucher. Duplicate item IDs found:");
            for (String id : duplicates) {
                System.out.println("  - " + id);
            }
            System.out.println("Please fix the data file and try again.");
            return;
        }

        // Step 3: Generate the formatted voucher string
        String voucher = generateVoucher(items);

        // Step 4: Display voucher on console
        System.out.print(voucher);

        // Step 5: Save voucher to a text file
        saveToFile(outputFile, voucher);
        System.out.println("\nVoucher saved to: " + outputFile);
    }

    /**
     * Reads grocery item data from a whitespace-separated file.
     * Each line: ItemID  Quantity  UnitPrice
     */
    public static List<GroceryItem> readGroceryItems(String fileName) {
        List<GroceryItem> items = new ArrayList<>();

        try (BufferedReader reader = new BufferedReader(new FileReader(fileName))) {
            String line;
            while ((line = reader.readLine()) != null) {
                line = line.trim();
                if (line.isEmpty()) continue; // skip blank lines

                // split("\\s+") splits by any whitespace (tab or spaces)
                String[] parts = line.split("\\s+");
                if (parts.length < 3) continue; // skip malformed lines

                String itemID = parts[0];
                int qty = Integer.parseInt(parts[1]);
                double unitPrice = Double.parseDouble(parts[2]);

                items.add(new GroceryItem(itemID, qty, unitPrice));
            }
        } catch (IOException e) {
            System.out.println("Error reading file: " + e.getMessage());
        }

        return items;
    }

    /**
     * Returns a list of item IDs that appear more than once.
     */
    public static List<String> findDuplicateItemIDs(List<GroceryItem> items) {
        Map<String, Integer> count = new LinkedHashMap<>();
        for (GroceryItem item : items) {
            count.put(item.itemID, count.getOrDefault(item.itemID, 0) + 1);
        }

        List<String> duplicates = new ArrayList<>();
        for (Map.Entry<String, Integer> entry : count.entrySet()) {
            if (entry.getValue() > 1) {
                duplicates.add(entry.getKey());
            }
        }
        return duplicates;
    }

    /**
     * Builds and returns the formatted voucher as a String.
     */
    public static String generateVoucher(List<GroceryItem> items) {
        StringBuilder sb = new StringBuilder();

        // Header row
        sb.append(String.format("%-12s\t%-6s\t%-12s\t%s%n",
                "Item id", "Qty", "unit price", "total"));

        double grandTotal = 0;
        int totalQty = 0;

        // Item rows
        for (GroceryItem item : items) {
            sb.append(String.format("%-12s\t%-6d\t%-12.1f\t%.1f%n",
                    item.itemID, item.quantity, item.unitPrice, item.total));
            grandTotal += item.total;
            totalQty += item.quantity;
        }

        // Summary rows
        double vat = grandTotal * 0.05;
        double netTotal = grandTotal - vat;

        sb.append(String.format("%n%-12s\t%-6d\t%-12s\t%.1f%n",
                "Total", totalQty, "Grand total", grandTotal));
        sb.append(String.format("VAT (5%%)  \t%.3f%n", vat));
        sb.append(String.format("Net total \t%.3f%n", netTotal));

        return sb.toString();
    }

    /**
     * Writes the given content to a text file.
     */
    public static void saveToFile(String fileName, String content) {
        try (PrintWriter writer = new PrintWriter(new FileWriter(fileName))) {
            writer.print(content);
        } catch (IOException e) {
            System.out.println("Error saving file: " + e.getMessage());
        }
    }
}
