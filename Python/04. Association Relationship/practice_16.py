"""
Practice 16: Library Management System
Task: A library with books (each having copies). Members borrow books.
      Regular members: max 3 books. Premium members: max 5 books.
      Late return fine: 10 taka/day (Regular), 5 taka/day (Premium).

How to run:
  python practice_16.py

Key Concepts:
  - Association between Library, Book, and Member
  - Membership-based borrowing rules
  - Fine calculation based on late returns
"""

from datetime import date, timedelta


class Book:
    """Book with title, author, and copy tracking."""

    def __init__(self, book_id: str, title: str, author: str, total_copies: int):
        """Creates a new Book."""
        self.book_id = book_id
        self.title = title
        self.author = author
        self.total_copies = total_copies
        self.available_copies = total_copies

    def is_available(self) -> bool:
        """Returns True if at least one copy is available."""
        return self.available_copies > 0

    def show_info(self):
        """Prints book details."""
        print(f"  {self.book_id:<8} | {self.title:<30} | {self.author:<20} | "
              f"Available: {self.available_copies}/{self.total_copies}")


class BorrowRecord:
    """Tracks a single book borrowing by a member."""

    def __init__(self, book: Book, borrow_date: date, due_date: date):
        """Creates a new BorrowRecord."""
        self.book = book
        self.borrow_date = borrow_date
        self.due_date = due_date
        self.return_date = None

    def is_returned(self) -> bool:
        """Returns True if the book has been returned."""
        return self.return_date is not None


class Member:
    """Library member who can borrow books."""

    BORROW_LIMITS = {"Regular": 3, "Premium": 5}
    FINE_RATES = {"Regular": 10.0, "Premium": 5.0}

    def __init__(self, member_id: str, name: str, membership_type: str):
        """Creates a new Member. membership_type: 'Regular' or 'Premium'."""
        if membership_type not in self.BORROW_LIMITS:
            raise ValueError(f"Invalid membership type: {membership_type}")
        self.member_id = member_id
        self.name = name
        self.membership_type = membership_type
        self.borrow_records = []
        self.total_fines = 0.0

    def borrow_limit(self) -> int:
        """Returns the maximum books this member can borrow."""
        return self.BORROW_LIMITS[self.membership_type]

    def fine_rate(self) -> float:
        """Returns the fine rate per day for this member."""
        return self.FINE_RATES[self.membership_type]

    def current_borrows(self) -> int:
        """Returns the number of books currently borrowed (not returned)."""
        return sum(1 for r in self.borrow_records if not r.is_returned())

    def show_info(self):
        """Prints member details."""
        print(f"  Member ID  : {self.member_id}")
        print(f"  Name       : {self.name}")
        print(f"  Membership : {self.membership_type}")
        print(f"  Books Out  : {self.current_borrows()} / {self.borrow_limit()}")
        print(f"  Total Fines: {self.total_fines:.2f}")
        print()


class Library:
    """Library that manages books and member borrowing."""

    def __init__(self, name: str):
        """Creates a new Library."""
        self.name = name
        self.books = []
        self.members = []

    def add_book(self, book: Book):
        """Adds a book to the library."""
        self.books.append(book)

    def add_member(self, member: Member):
        """Registers a member with the library."""
        self.members.append(member)

    def borrow_book(self, member: Member, book: Book, borrow_date: date, loan_days: int = 14):
        """Lends a book to a member."""
        if not book.is_available():
            print(f"  [Error] '{book.title}' has no available copies.")
            return
        if member.current_borrows() >= member.borrow_limit():
            print(f"  [Error] {member.name} has reached the borrowing limit "
                  f"({member.borrow_limit()} books for {member.membership_type} members).")
            return
        book.available_copies -= 1
        due_date = borrow_date + timedelta(days=loan_days)
        record = BorrowRecord(book, borrow_date, due_date)
        member.borrow_records.append(record)
        print(f"  [OK] {member.name} borrowed '{book.title}'. Due: {due_date}")

    def return_book(self, member: Member, book: Book, return_date: date):
        """Returns a book and calculates fine if late."""
        for record in member.borrow_records:
            if record.book.book_id == book.book_id and not record.is_returned():
                record.return_date = return_date
                book.available_copies += 1

                fine = 0.0
                if return_date > record.due_date:
                    late_days = (return_date - record.due_date).days
                    fine = late_days * member.fine_rate()
                    member.total_fines += fine
                    print(f"  [OK] {member.name} returned '{book.title}'. "
                          f"Late by {late_days} day(s). Fine: {fine:.2f} taka.")
                else:
                    print(f"  [OK] {member.name} returned '{book.title}' on time.")
                return

        print(f"  [Error] {member.name} has no active borrow record for '{book.title}'.")

    def show_catalog(self):
        """Prints all books in the library."""
        print(f"  === {self.name} Catalog ===")
        print(f"  {'ID':<8} | {'Title':<30} | {'Author':<20} | Availability")
        print("  " + "-" * 80)
        for book in self.books:
            book.show_info()
        print()


def main():
    # --- Setup library ---
    library = Library("City Public Library")

    book1 = Book("B-001", "Clean Code", "Robert C. Martin", 3)
    book2 = Book("B-002", "Design Patterns", "Gang of Four", 2)
    book3 = Book("B-003", "The Pragmatic Programmer", "Hunt & Thomas", 1)
    book4 = Book("B-004", "Refactoring", "Martin Fowler", 2)

    for b in [book1, book2, book3, book4]:
        library.add_book(b)

    # --- Register members ---
    regular = Member("M-001", "Tareq", "Regular")    # Max 3 books, 10 taka/day fine
    premium = Member("M-002", "Afsana", "Premium")   # Max 5 books, 5 taka/day fine

    library.add_member(regular)
    library.add_member(premium)

    print("=== Library Catalog ===")
    library.show_catalog()

    print("=== Member Info ===")
    regular.show_info()
    premium.show_info()

    # --- Borrow books ---
    print("=== Borrowing Books ===")
    borrow_date = date(2026, 2, 1)
    library.borrow_book(regular, book1, borrow_date)
    library.borrow_book(regular, book2, borrow_date)
    library.borrow_book(regular, book3, borrow_date)
    library.borrow_book(regular, book4, borrow_date)  # Should fail — limit reached
    print()

    library.borrow_book(premium, book1, borrow_date)
    library.borrow_book(premium, book2, borrow_date)
    library.borrow_book(premium, book4, borrow_date)
    print()

    print("=== After Borrowing ===")
    library.show_catalog()
    regular.show_info()
    premium.show_info()

    # --- Return books ---
    print("=== Returning Books ===")
    library.return_book(regular, book1, date(2026, 2, 10))   # On time (due Feb 15)
    library.return_book(regular, book2, date(2026, 2, 20))   # 5 days late
    library.return_book(premium, book1, date(2026, 2, 25))   # 10 days late (premium rate)
    print()

    print("=== After Returns ===")
    library.show_catalog()
    regular.show_info()
    premium.show_info()


if __name__ == "__main__":
    main()
