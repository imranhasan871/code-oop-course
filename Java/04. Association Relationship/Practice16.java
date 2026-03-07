/**
 * Practice 16: Library Management System
 * Task: Library with books (each having copies). Members borrow books.
 *       Regular members: max 3 books. Premium members: max 5 books.
 *       Late return fine: 10 taka/day (Regular), 5 taka/day (Premium).
 *
 * How to compile and run:
 *   javac Practice16.java
 *   java Practice16
 *
 * Key Concepts:
 *   - Association between Library, Book, and Member
 *   - Membership-based borrowing rules
 *   - Fine calculation based on late returns
 */

import java.time.LocalDate;
import java.time.temporal.ChronoUnit;
import java.util.ArrayList;
import java.util.List;

public class Practice16 {

    /** Book with title, author, and copy tracking. */
    static class Book {
        private String bookId;
        private String title;
        private String author;
        private int totalCopies;
        private int availableCopies;

        public Book(String bookId, String title, String author, int totalCopies) {
            this.bookId = bookId;
            this.title = title;
            this.author = author;
            this.totalCopies = totalCopies;
            this.availableCopies = totalCopies;
        }

        public String getBookId() { return bookId; }
        public String getTitle() { return title; }
        public boolean isAvailable() { return availableCopies > 0; }

        public void showInfo() {
            System.out.printf("  %-8s | %-30s | %-20s | Available: %d/%d%n",
                    bookId, title, author, availableCopies, totalCopies);
        }
    }

    /** Tracks a single book borrowing by a member. */
    static class BorrowRecord {
        private Book book;
        private LocalDate borrowDate;
        private LocalDate dueDate;
        private LocalDate returnDate;

        public BorrowRecord(Book book, LocalDate borrowDate, LocalDate dueDate) {
            this.book = book;
            this.borrowDate = borrowDate;
            this.dueDate = dueDate;
        }

        public boolean isReturned() { return returnDate != null; }
        public Book getBook() { return book; }
        public LocalDate getDueDate() { return dueDate; }
        public void setReturnDate(LocalDate date) { this.returnDate = date; }
    }

    /** Library member who can borrow books. */
    static class Member {
        private String memberId;
        private String name;
        private String membershipType; // "Regular" or "Premium"
        private List<BorrowRecord> borrowRecords = new ArrayList<>();
        private double totalFines = 0;

        public Member(String memberId, String name, String membershipType) {
            this.memberId = memberId;
            this.name = name;
            this.membershipType = membershipType;
        }

        public String getName() { return name; }
        public String getMembershipType() { return membershipType; }

        public int borrowLimit() {
            return membershipType.equals("Premium") ? 5 : 3;
        }

        public double fineRate() {
            return membershipType.equals("Premium") ? 5.0 : 10.0;
        }

        public int currentBorrows() {
            int count = 0;
            for (BorrowRecord r : borrowRecords) {
                if (!r.isReturned()) count++;
            }
            return count;
        }

        public void showInfo() {
            System.out.println("  Member ID  : " + memberId);
            System.out.println("  Name       : " + name);
            System.out.println("  Membership : " + membershipType);
            System.out.printf("  Books Out  : %d / %d%n", currentBorrows(), borrowLimit());
            System.out.printf("  Total Fines: %.2f%n", totalFines);
            System.out.println();
        }
    }

    /** Library that manages books and member borrowing. */
    static class Library {
        private String name;
        private List<Book> books = new ArrayList<>();
        private List<Member> members = new ArrayList<>();

        public Library(String name) {
            this.name = name;
        }

        public void addBook(Book book) { books.add(book); }
        public void addMember(Member member) { members.add(member); }

        public void borrowBook(Member member, Book book, LocalDate borrowDate, int loanDays) {
            if (!book.isAvailable()) {
                System.out.printf("  [Error] '%s' has no available copies.%n", book.getTitle());
                return;
            }
            if (member.currentBorrows() >= member.borrowLimit()) {
                System.out.printf("  [Error] %s has reached the borrowing limit (%d books for %s members).%n",
                        member.getName(), member.borrowLimit(), member.getMembershipType());
                return;
            }
            book.availableCopies--;
            LocalDate dueDate = borrowDate.plusDays(loanDays);
            BorrowRecord record = new BorrowRecord(book, borrowDate, dueDate);
            member.borrowRecords.add(record);
            System.out.printf("  [OK] %s borrowed '%s'. Due: %s%n", member.getName(), book.getTitle(), dueDate);
        }

        public void returnBook(Member member, Book book, LocalDate returnDate) {
            for (BorrowRecord record : member.borrowRecords) {
                if (record.getBook().getBookId().equals(book.getBookId()) && !record.isReturned()) {
                    record.setReturnDate(returnDate);
                    book.availableCopies++;

                    if (returnDate.isAfter(record.getDueDate())) {
                        long lateDays = ChronoUnit.DAYS.between(record.getDueDate(), returnDate);
                        double fine = lateDays * member.fineRate();
                        member.totalFines += fine;
                        System.out.printf("  [OK] %s returned '%s'. Late by %d day(s). Fine: %.2f taka.%n",
                                member.getName(), book.getTitle(), lateDays, fine);
                    } else {
                        System.out.printf("  [OK] %s returned '%s' on time.%n",
                                member.getName(), book.getTitle());
                    }
                    return;
                }
            }
            System.out.printf("  [Error] %s has no active borrow record for '%s'.%n",
                    member.getName(), book.getTitle());
        }

        public void showCatalog() {
            System.out.println("  === " + name + " Catalog ===");
            System.out.printf("  %-8s | %-30s | %-20s | Availability%n", "ID", "Title", "Author");
            System.out.println("  " + "-".repeat(80));
            for (Book book : books) {
                book.showInfo();
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        Library library = new Library("City Public Library");

        Book book1 = new Book("B-001", "Clean Code", "Robert C. Martin", 3);
        Book book2 = new Book("B-002", "Design Patterns", "Gang of Four", 2);
        Book book3 = new Book("B-003", "The Pragmatic Programmer", "Hunt & Thomas", 1);
        Book book4 = new Book("B-004", "Refactoring", "Martin Fowler", 2);

        for (Book b : new Book[]{book1, book2, book3, book4}) {
            library.addBook(b);
        }

        Member regular = new Member("M-001", "Tareq", "Regular");
        Member premium = new Member("M-002", "Afsana", "Premium");
        library.addMember(regular);
        library.addMember(premium);

        System.out.println("=== Library Catalog ===");
        library.showCatalog();

        System.out.println("=== Member Info ===");
        regular.showInfo();
        premium.showInfo();

        System.out.println("=== Borrowing Books ===");
        LocalDate borrowDate = LocalDate.of(2026, 2, 1);
        library.borrowBook(regular, book1, borrowDate, 14);
        library.borrowBook(regular, book2, borrowDate, 14);
        library.borrowBook(regular, book3, borrowDate, 14);
        library.borrowBook(regular, book4, borrowDate, 14); // Should fail
        System.out.println();

        library.borrowBook(premium, book1, borrowDate, 14);
        library.borrowBook(premium, book2, borrowDate, 14);
        library.borrowBook(premium, book4, borrowDate, 14);
        System.out.println();

        System.out.println("=== After Borrowing ===");
        library.showCatalog();
        regular.showInfo();
        premium.showInfo();

        System.out.println("=== Returning Books ===");
        library.returnBook(regular, book1, LocalDate.of(2026, 2, 10));  // On time
        library.returnBook(regular, book2, LocalDate.of(2026, 2, 20));  // 5 days late
        library.returnBook(premium, book1, LocalDate.of(2026, 2, 25));  // 10 days late
        System.out.println();

        System.out.println("=== After Returns ===");
        library.showCatalog();
        regular.showInfo();
        premium.showInfo();
    }
}
