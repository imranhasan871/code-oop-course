"""
Practice 03: Generate a Voucher
Task: Read grocery items from a CSV file, calculate totals, and generate
      a formatted voucher. Save the voucher to a text file.
      If there are duplicate item IDs in the file, show an error and stop.

How to run:
  python practice_03.py

Expected output on console:
  Item id      Qty    unit price    total
  item-937     12     230.5         2766.0
  ...
  Total        29     Grand total   4658.5
  VAT (5%)     232.925
  Net total    4425.575

NOTE: The data file is at: ../../data/groceryitems.csv
      The voucher will be saved to: voucher.txt (in the same directory)

Course: Professional OOP — by Zohirul Alam Tiemoon
"""


class GroceryItem:
    """Holds the details for one grocery line item."""

    def __init__(self, item_id: str, quantity: int, unit_price: float):
        self.item_id = item_id
        self.quantity = quantity
        self.unit_price = unit_price
        self.total = quantity * unit_price


def read_grocery_items(file_name: str) -> list:
    """
    Reads grocery item data from a whitespace-separated file.
    Each line: ItemID  Quantity  UnitPrice
    """
    items = []

    with open(file_name, "r") as file:
        for line in file:
            line = line.strip()
            if not line:
                continue  # skip blank lines

            parts = line.split()
            if len(parts) < 3:
                continue  # skip malformed lines

            item_id = parts[0]
            try:
                qty = int(parts[1])
                unit_price = float(parts[2])
            except ValueError:
                continue

            items.append(GroceryItem(item_id, qty, unit_price))

    return items


def find_duplicate_item_ids(items: list) -> list:
    """Returns a list of item IDs that appear more than once."""
    count = {}
    for item in items:
        count[item.item_id] = count.get(item.item_id, 0) + 1

    return [item_id for item_id, n in count.items() if n > 1]


def generate_voucher(items: list) -> str:
    """Builds and returns the formatted voucher as a string."""
    lines = []

    # Header row
    lines.append(f"{'Item id':<12}\t{'Qty':<6}\t{'unit price':<12}\t{'total'}")

    grand_total = 0.0
    total_qty = 0

    # Item rows
    for item in items:
        lines.append(f"{item.item_id:<12}\t{item.quantity:<6}\t"
                      f"{item.unit_price:<12.1f}\t{item.total:.1f}")
        grand_total += item.total
        total_qty += item.quantity

    # Summary rows
    vat = grand_total * 0.05
    net_total = grand_total - vat

    lines.append("")
    lines.append(f"{'Total':<12}\t{total_qty:<6}\t{'Grand total':<12}\t{grand_total:.1f}")
    lines.append(f"VAT (5%)\t{vat:.3f}")
    lines.append(f"Net total\t{net_total:.3f}")

    return "\n".join(lines) + "\n"


def save_to_file(file_name: str, content: str):
    """Writes the given content string to a text file."""
    with open(file_name, "w") as file:
        file.write(content)


if __name__ == "__main__":
    input_file = "../../data/groceryitems.csv"
    output_file = "voucher.txt"

    # Step 1: Read items from file
    items = read_grocery_items(input_file)

    if not items:
        print("No grocery items found in the file.")
    else:
        # Step 2: Check for duplicate item IDs before processing
        duplicates = find_duplicate_item_ids(items)
        if duplicates:
            print("Error: Cannot generate voucher. Duplicate item IDs found:")
            for dup_id in duplicates:
                print(f"  - {dup_id}")
            print("Please fix the data file and try again.")
        else:
            # Step 3: Generate the formatted voucher string
            voucher = generate_voucher(items)

            # Step 4: Display voucher on console
            print(voucher, end="")

            # Step 5: Save voucher to a text file
            save_to_file(output_file, voucher)
            print(f"\nVoucher saved to: {output_file}")
