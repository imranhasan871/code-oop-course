/**
 * Practice 03: Generate a Voucher
 * Task: Read grocery items from a CSV file, calculate totals, and generate
 *       a formatted voucher. Save the voucher to a text file.
 *       If there are duplicate item IDs in the file, show an error and stop.
 *
 * How to run:
 *   Create a Console App, paste this code into Program.cs, and run.
 *   OR use: dotnet-script Practice03.cs
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

using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

class Practice03
{
    /**
     * Inner class: represents one grocery line item.
     */
    class GroceryItem
    {
        public string ItemID { get; set; }
        public int Quantity { get; set; }
        public double UnitPrice { get; set; }
        public double Total => Quantity * UnitPrice; // Calculated property

        public GroceryItem(string itemID, int quantity, double unitPrice)
        {
            ItemID = itemID;
            Quantity = quantity;
            UnitPrice = unitPrice;
        }
    }

    static void Main(string[] args)
    {
        string inputFile = "../../data/groceryitems.csv";
        string outputFile = "voucher.txt";

        // Step 1: Read items from file
        List<GroceryItem> items = ReadGroceryItems(inputFile);

        if (items == null || items.Count == 0)
        {
            Console.WriteLine("No grocery items found in the file.");
            return;
        }

        // Step 2: Check for duplicate item IDs before processing
        List<string> duplicates = FindDuplicateItemIDs(items);
        if (duplicates.Count > 0)
        {
            Console.WriteLine("Error: Cannot generate voucher. Duplicate item IDs found:");
            foreach (string id in duplicates)
            {
                Console.WriteLine($"  - {id}");
            }
            Console.WriteLine("Please fix the data file and try again.");
            return;
        }

        // Step 3: Generate the formatted voucher string
        string voucher = GenerateVoucher(items);

        // Step 4: Display voucher on console
        Console.Write(voucher);

        // Step 5: Save voucher to a text file
        SaveToFile(outputFile, voucher);
        Console.WriteLine($"\nVoucher saved to: {outputFile}");
    }

    /**
     * Reads grocery item data from a whitespace-separated file.
     * Each line: ItemID  Quantity  UnitPrice
     */
    static List<GroceryItem> ReadGroceryItems(string fileName)
    {
        List<GroceryItem> items = new List<GroceryItem>();

        try
        {
            string[] lines = File.ReadAllLines(fileName);
            foreach (string line in lines)
            {
                string trimmed = line.Trim();
                if (string.IsNullOrEmpty(trimmed)) continue; // skip blank lines

                // Split by any whitespace (tab or spaces), removing empty entries
                string[] parts = trimmed.Split(
                    new char[] { ' ', '\t' },
                    StringSplitOptions.RemoveEmptyEntries
                );
                if (parts.Length < 3) continue; // skip malformed lines

                string itemID = parts[0];
                int qty = int.Parse(parts[1]);
                double unitPrice = double.Parse(parts[2]);

                items.Add(new GroceryItem(itemID, qty, unitPrice));
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error reading file: {ex.Message}");
        }

        return items;
    }

    /**
     * Returns a list of item IDs that appear more than once.
     */
    static List<string> FindDuplicateItemIDs(List<GroceryItem> items)
    {
        Dictionary<string, int> count = new Dictionary<string, int>();
        foreach (GroceryItem item in items)
        {
            if (count.ContainsKey(item.ItemID))
                count[item.ItemID]++;
            else
                count[item.ItemID] = 1;
        }

        List<string> duplicates = new List<string>();
        foreach (KeyValuePair<string, int> entry in count)
        {
            if (entry.Value > 1)
                duplicates.Add(entry.Key);
        }
        return duplicates;
    }

    /**
     * Builds and returns the formatted voucher as a string.
     */
    static string GenerateVoucher(List<GroceryItem> items)
    {
        StringBuilder sb = new StringBuilder();

        // Header row
        sb.AppendLine($"{"Item id",-12}\t{"Qty",-6}\t{"unit price",-12}\ttotal");

        double grandTotal = 0;
        int totalQty = 0;

        // Item rows
        foreach (GroceryItem item in items)
        {
            sb.AppendLine($"{item.ItemID,-12}\t{item.Quantity,-6}\t{item.UnitPrice,-12:F1}\t{item.Total:F1}");
            grandTotal += item.Total;
            totalQty += item.Quantity;
        }

        // Summary rows
        double vat = grandTotal * 0.05;
        double netTotal = grandTotal - vat;

        sb.AppendLine();
        sb.AppendLine($"{"Total",-12}\t{totalQty,-6}\t{"Grand total",-12}\t{grandTotal:F1}");
        sb.AppendLine($"VAT (5%)\t{vat:F3}");
        sb.AppendLine($"Net total\t{netTotal:F3}");

        return sb.ToString();
    }

    /**
     * Writes the given content to a text file.
     */
    static void SaveToFile(string fileName, string content)
    {
        try
        {
            File.WriteAllText(fileName, content);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error saving file: {ex.Message}");
        }
    }
}
