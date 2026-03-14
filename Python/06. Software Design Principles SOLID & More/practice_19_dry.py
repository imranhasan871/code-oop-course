"""
Practice 19: Applying DRY (Don't Repeat Yourself)
Task: Refactor duplicated discount and invoice logic into reusable components.

How to run:
  python practice_19_dry.py
"""

from dataclasses import dataclass


@dataclass
class OrderItem:
    name: str
    unit_price: float
    quantity: int


class DiscountPolicy:
    """Single source of truth for discount logic."""

    @staticmethod
    def calculate(subtotal: float, customer_type: str) -> float:
        rates = {
            "regular": 0.00,
            "silver": 0.05,
            "gold": 0.10,
        }
        rate = rates.get(customer_type.lower(), 0.00)
        return subtotal * rate


class InvoiceCalculator:
    """Centralized invoice calculation to avoid repeated formulas."""

    VAT_RATE = 0.05

    def summarize(self, items: list[OrderItem], customer_type: str) -> dict:
        subtotal = sum(item.unit_price * item.quantity for item in items)
        discount = DiscountPolicy.calculate(subtotal, customer_type)
        taxable_amount = subtotal - discount
        vat = taxable_amount * self.VAT_RATE
        total = taxable_amount + vat

        return {
            "subtotal": subtotal,
            "discount": discount,
            "vat": vat,
            "total": total,
        }


def print_invoice(customer_name: str, customer_type: str, items: list[OrderItem]) -> None:
    calculator = InvoiceCalculator()
    summary = calculator.summarize(items, customer_type)

    print(f"Customer: {customer_name} ({customer_type.title()})")
    for item in items:
        line_total = item.unit_price * item.quantity
        print(f"  - {item.name:<10} x{item.quantity:<2} @ ${item.unit_price:>6.2f} = ${line_total:>7.2f}")

    print(f"Subtotal : ${summary['subtotal']:.2f}")
    print(f"Discount : ${summary['discount']:.2f}")
    print(f"VAT (5%) : ${summary['vat']:.2f}")
    print(f"Total    : ${summary['total']:.2f}")
    print()


def main() -> None:
    items = [
        OrderItem("Keyboard", 45.0, 2),
        OrderItem("Mouse", 20.0, 1),
        OrderItem("USB-C", 12.5, 3),
    ]

    print_invoice("Samia", "regular", items)
    print_invoice("Afsana", "silver", items)
    print_invoice("Hasan", "gold", items)


if __name__ == "__main__":
    main()
