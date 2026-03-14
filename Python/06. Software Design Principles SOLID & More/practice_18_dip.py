"""
Practice 18: Applying DIP (Dependency Inversion Principle)
Task: Build a payment system where PaymentService depends on abstraction,
      not concrete payment providers.

How to run:
  python practice_18_dip.py
"""

from abc import ABC, abstractmethod


class PaymentMethod(ABC):
    @abstractmethod
    def pay(self, amount: float) -> str:
        raise NotImplementedError


class CreditCardPayment(PaymentMethod):
    def pay(self, amount: float) -> str:
        return f"Paid ${amount:.2f} using Credit Card"


class BkashPayment(PaymentMethod):
    def pay(self, amount: float) -> str:
        return f"Paid ${amount:.2f} using bKash"


class PaypalPayment(PaymentMethod):
    def pay(self, amount: float) -> str:
        return f"Paid ${amount:.2f} using PayPal"


class BankTransferPayment(PaymentMethod):
    def pay(self, amount: float) -> str:
        return f"Paid ${amount:.2f} using Bank Transfer"


class PaymentService:
    """High-level module depending only on the PaymentMethod abstraction."""

    def process_payment(self, payment_method: PaymentMethod, amount: float) -> None:
        if amount <= 0:
            raise ValueError("Amount must be greater than 0.")
        result = payment_method.pay(amount)
        print(f"[PaymentService] {result}")


def main() -> None:
    payment_service = PaymentService()

    methods: list[PaymentMethod] = [
        CreditCardPayment(),
        BkashPayment(),
        PaypalPayment(),
        BankTransferPayment(),
    ]

    for method in methods:
        payment_service.process_payment(method, 1500.0)


if __name__ == "__main__":
    main()
