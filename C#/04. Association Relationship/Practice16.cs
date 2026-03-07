/**
 * Practice 16: Library Management System
 * Task: Library with books (each having copies). Members borrow books.
 *       Regular members: max 3 books. Premium members: max 5 books.
 *       Late return fine: 10 taka/day (Regular), 5 taka/day (Premium).
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice16.cs
 *
 * Key Concepts:
 *   - Association between Library, Book, and Member
 *   - Membership-based borrowing rules
 *   - Fine calculation based on late returns
 */

using System;
using System.Collections.Generic;
using System.Linq;

class Practice16
{
    class Book
    {
        public string BookId { get; }
        public string Title { get; }
        public string Author { get; }
        public int TotalCopies { get; }
        public int AvailableCopies { get; set; }

        public Book(string bookId, string title, string author, int totalCopies)
        {
            BookId = bookId;
            Title = title;
            Author = author;
            TotalCopies = totalCopies;
            AvailableCopies = totalCopies;
        }

        public bool IsAvailable() => AvailableCopies > 0;

        public void ShowInfo()
        {
            Console.WriteLine($"  {BookId,-8} | {Title,-30} | {Author,-20} | Available: {AvailableCopies}/{TotalCopies}");
        }
    }

    class BorrowRecord
    {
        public Book Book { get; }
        public DateTime BorrowDate { get; }
        public DateTime DueDate { get; }
        public DateTime? ReturnDate { get; set; }

        public BorrowRecord(Book book, DateTime borrowDate, DateTime dueDate)
        {
            Book = book;
            BorrowDate = borrowDate;
            DueDate = dueDate;
        }

        public bool IsReturned() => ReturnDate != null;
    }

    class Member
    {
        public string MemberId { get; }
        public string Name { get; }
        public string MembershipType { get; }
        public List<BorrowRecord> BorrowRecords { get; } = new List<BorrowRecord>();
        public double TotalFines { get; set; } = 0;

        public Member(string memberId, string name, string membershipType)
        {
            MemberId = memberId;
            Name = name;
            MembershipType = membershipType;
        }

        public int BorrowLimit() => MembershipType == "Premium" ? 5 : 3;
        public double FineRate() => MembershipType == "Premium" ? 5.0 : 10.0;
        public int CurrentBorrows() => BorrowRecords.Count(r => !r.IsReturned());

        public void ShowInfo()
        {
            Console.WriteLine($"  Member ID  : {MemberId}");
            Console.WriteLine($"  Name       : {Name}");
            Console.WriteLine($"  Membership : {MembershipType}");
            Console.WriteLine($"  Books Out  : {CurrentBorrows()} / {BorrowLimit()}");
            Console.WriteLine($"  Total Fines: {TotalFines:F2}");
            Console.WriteLine();
        }
    }

    class Library
    {
        public string Name { get; }
        private List<Book> books = new List<Book>();
        private List<Member> members = new List<Member>();

        public Library(string name) { Name = name; }

        public void AddBook(Book book) { books.Add(book); }
        public void AddMember(Member member) { members.Add(member); }

        public void BorrowBook(Member member, Book book, DateTime borrowDate, int loanDays = 14)
        {
            if (!book.IsAvailable())
            {
                Console.WriteLine($"  [Error] '{book.Title}' has no available copies.");
                return;
            }
            if (member.CurrentBorrows() >= member.BorrowLimit())
            {
                Console.WriteLine($"  [Error] {member.Name} has reached the borrowing limit " +
                                  $"({member.BorrowLimit()} books for {member.MembershipType} members).");
                return;
            }
            book.AvailableCopies--;
            var dueDate = borrowDate.AddDays(loanDays);
            var record = new BorrowRecord(book, borrowDate, dueDate);
            member.BorrowRecords.Add(record);
            Console.WriteLine($"  [OK] {member.Name} borrowed '{book.Title}'. Due: {dueDate:yyyy-MM-dd}");
        }

        public void ReturnBook(Member member, Book book, DateTime returnDate)
        {
            foreach (var record in member.BorrowRecords)
            {
                if (record.Book.BookId == book.BookId && !record.IsReturned())
                {
                    record.ReturnDate = returnDate;
                    book.AvailableCopies++;

                    if (returnDate > record.DueDate)
                    {
                        int lateDays = (returnDate - record.DueDate).Days;
                        double fine = lateDays * member.FineRate();
                        member.TotalFines += fine;
                        Console.WriteLine($"  [OK] {member.Name} returned '{book.Title}'. " +
                                          $"Late by {lateDays} day(s). Fine: {fine:F2} taka.");
                    }
                    else
                    {
                        Console.WriteLine($"  [OK] {member.Name} returned '{book.Title}' on time.");
                    }
                    return;
                }
            }
            Console.WriteLine($"  [Error] {member.Name} has no active borrow record for '{book.Title}'.");
        }

        public void ShowCatalog()
        {
            Console.WriteLine($"  === {Name} Catalog ===");
            Console.WriteLine($"  {"ID",-8} | {"Title",-30} | {"Author",-20} | Availability");
            Console.WriteLine("  " + new string('-', 80));
            foreach (var book in books)
                book.ShowInfo();
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        var library = new Library("City Public Library");

        var book1 = new Book("B-001", "Clean Code", "Robert C. Martin", 3);
        var book2 = new Book("B-002", "Design Patterns", "Gang of Four", 2);
        var book3 = new Book("B-003", "The Pragmatic Programmer", "Hunt & Thomas", 1);
        var book4 = new Book("B-004", "Refactoring", "Martin Fowler", 2);

        foreach (var b in new[] { book1, book2, book3, book4 })
            library.AddBook(b);

        var regular = new Member("M-001", "Tareq", "Regular");
        var premium = new Member("M-002", "Afsana", "Premium");
        library.AddMember(regular);
        library.AddMember(premium);

        Console.WriteLine("=== Library Catalog ===");
        library.ShowCatalog();

        Console.WriteLine("=== Member Info ===");
        regular.ShowInfo();
        premium.ShowInfo();

        Console.WriteLine("=== Borrowing Books ===");
        var borrowDate = new DateTime(2026, 2, 1);
        library.BorrowBook(regular, book1, borrowDate);
        library.BorrowBook(regular, book2, borrowDate);
        library.BorrowBook(regular, book3, borrowDate);
        library.BorrowBook(regular, book4, borrowDate); // Should fail
        Console.WriteLine();

        library.BorrowBook(premium, book1, borrowDate);
        library.BorrowBook(premium, book2, borrowDate);
        library.BorrowBook(premium, book4, borrowDate);
        Console.WriteLine();

        Console.WriteLine("=== After Borrowing ===");
        library.ShowCatalog();
        regular.ShowInfo();
        premium.ShowInfo();

        Console.WriteLine("=== Returning Books ===");
        library.ReturnBook(regular, book1, new DateTime(2026, 2, 10));  // On time
        library.ReturnBook(regular, book2, new DateTime(2026, 2, 20));  // 5 days late
        library.ReturnBook(premium, book1, new DateTime(2026, 2, 25));  // 10 days late
        Console.WriteLine();

        Console.WriteLine("=== After Returns ===");
        library.ShowCatalog();
        regular.ShowInfo();
        premium.ShowInfo();
    }
}
