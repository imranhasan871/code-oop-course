/**
 * Practice 16: Library Management System
 * Task: Library with books (each having copies). Members borrow books.
 *       Regular members: max 3 books. Premium members: max 5 books.
 *       Late return fine: 10 taka/day (Regular), 5 taka/day (Premium).
 *
 * How to run:
 *   go run practice-16.go
 *
 * Key Concepts:
 *   - Association between Library, Book, and Member
 *   - Membership-based borrowing rules
 *   - Fine calculation based on late returns
 */

package main

import (
	"fmt"
	"time"
)

/** Book with title, author, and copy tracking. */
type Book16 struct {
	BookId          string
	Title           string
	Author          string
	TotalCopies     int
	AvailableCopies int
}

func NewBook16(bookId, title, author string, totalCopies int) *Book16 {
	return &Book16{bookId, title, author, totalCopies, totalCopies}
}

func (b *Book16) IsAvailable() bool {
	return b.AvailableCopies > 0
}

func (b *Book16) ShowInfo() {
	fmt.Printf("  %-8s | %-30s | %-20s | Available: %d/%d\n",
		b.BookId, b.Title, b.Author, b.AvailableCopies, b.TotalCopies)
}

/** BorrowRecord tracks a single book borrowing. */
type BorrowRecord16 struct {
	Book       *Book16
	BorrowDate time.Time
	DueDate    time.Time
	ReturnDate *time.Time
}

func (r *BorrowRecord16) IsReturned() bool {
	return r.ReturnDate != nil
}

/** Member — library member who can borrow books. */
type Member16 struct {
	MemberId       string
	Name           string
	MembershipType string // "Regular" or "Premium"
	BorrowRecords  []*BorrowRecord16
	TotalFines     float64
}

func NewMember16(memberId, name, membershipType string) *Member16 {
	return &Member16{MemberId: memberId, Name: name, MembershipType: membershipType}
}

func (m *Member16) BorrowLimit() int {
	if m.MembershipType == "Premium" {
		return 5
	}
	return 3
}

func (m *Member16) FineRate() float64 {
	if m.MembershipType == "Premium" {
		return 5.0
	}
	return 10.0
}

func (m *Member16) CurrentBorrows() int {
	count := 0
	for _, r := range m.BorrowRecords {
		if !r.IsReturned() {
			count++
		}
	}
	return count
}

func (m *Member16) ShowInfo() {
	fmt.Println("  Member ID  :", m.MemberId)
	fmt.Println("  Name       :", m.Name)
	fmt.Println("  Membership :", m.MembershipType)
	fmt.Printf("  Books Out  : %d / %d\n", m.CurrentBorrows(), m.BorrowLimit())
	fmt.Printf("  Total Fines: %.2f\n", m.TotalFines)
	fmt.Println()
}

/** Library manages books and member borrowing. */
type Library16 struct {
	Name    string
	Books   []*Book16
	Members []*Member16
}

func NewLibrary16(name string) *Library16 {
	return &Library16{Name: name}
}

func (lib *Library16) AddBook(book *Book16) {
	lib.Books = append(lib.Books, book)
}

func (lib *Library16) AddMember(member *Member16) {
	lib.Members = append(lib.Members, member)
}

func (lib *Library16) BorrowBook(member *Member16, book *Book16, borrowDate time.Time, loanDays int) {
	if !book.IsAvailable() {
		fmt.Printf("  [Error] '%s' has no available copies.\n", book.Title)
		return
	}
	if member.CurrentBorrows() >= member.BorrowLimit() {
		fmt.Printf("  [Error] %s has reached the borrowing limit (%d books for %s members).\n",
			member.Name, member.BorrowLimit(), member.MembershipType)
		return
	}
	book.AvailableCopies--
	dueDate := borrowDate.AddDate(0, 0, loanDays)
	record := &BorrowRecord16{Book: book, BorrowDate: borrowDate, DueDate: dueDate}
	member.BorrowRecords = append(member.BorrowRecords, record)
	fmt.Printf("  [OK] %s borrowed '%s'. Due: %s\n", member.Name, book.Title, dueDate.Format("2006-01-02"))
}

func (lib *Library16) ReturnBook(member *Member16, book *Book16, returnDate time.Time) {
	for _, record := range member.BorrowRecords {
		if record.Book.BookId == book.BookId && !record.IsReturned() {
			record.ReturnDate = &returnDate
			book.AvailableCopies++

			if returnDate.After(record.DueDate) {
				lateDays := int(returnDate.Sub(record.DueDate).Hours() / 24)
				fine := float64(lateDays) * member.FineRate()
				member.TotalFines += fine
				fmt.Printf("  [OK] %s returned '%s'. Late by %d day(s). Fine: %.2f taka.\n",
					member.Name, book.Title, lateDays, fine)
			} else {
				fmt.Printf("  [OK] %s returned '%s' on time.\n", member.Name, book.Title)
			}
			return
		}
	}
	fmt.Printf("  [Error] %s has no active borrow record for '%s'.\n", member.Name, book.Title)
}

func (lib *Library16) ShowCatalog() {
	fmt.Printf("  === %s Catalog ===\n", lib.Name)
	fmt.Printf("  %-8s | %-30s | %-20s | Availability\n", "ID", "Title", "Author")
	fmt.Println("  --------------------------------------------------------------------------------")
	for _, book := range lib.Books {
		book.ShowInfo()
	}
	fmt.Println()
}

func main() {
	library := NewLibrary16("City Public Library")

	book1 := NewBook16("B-001", "Clean Code", "Robert C. Martin", 3)
	book2 := NewBook16("B-002", "Design Patterns", "Gang of Four", 2)
	book3 := NewBook16("B-003", "The Pragmatic Programmer", "Hunt & Thomas", 1)
	book4 := NewBook16("B-004", "Refactoring", "Martin Fowler", 2)

	for _, b := range []*Book16{book1, book2, book3, book4} {
		library.AddBook(b)
	}

	regular := NewMember16("M-001", "Tareq", "Regular")
	premium := NewMember16("M-002", "Afsana", "Premium")
	library.AddMember(regular)
	library.AddMember(premium)

	fmt.Println("=== Library Catalog ===")
	library.ShowCatalog()

	fmt.Println("=== Member Info ===")
	regular.ShowInfo()
	premium.ShowInfo()

	fmt.Println("=== Borrowing Books ===")
	borrowDate := time.Date(2026, 2, 1, 0, 0, 0, 0, time.Local)
	library.BorrowBook(regular, book1, borrowDate, 14)
	library.BorrowBook(regular, book2, borrowDate, 14)
	library.BorrowBook(regular, book3, borrowDate, 14)
	library.BorrowBook(regular, book4, borrowDate, 14) // Should fail
	fmt.Println()

	library.BorrowBook(premium, book1, borrowDate, 14)
	library.BorrowBook(premium, book2, borrowDate, 14)
	library.BorrowBook(premium, book4, borrowDate, 14)
	fmt.Println()

	fmt.Println("=== After Borrowing ===")
	library.ShowCatalog()
	regular.ShowInfo()
	premium.ShowInfo()

	fmt.Println("=== Returning Books ===")
	library.ReturnBook(regular, book1, time.Date(2026, 2, 10, 0, 0, 0, 0, time.Local))
	library.ReturnBook(regular, book2, time.Date(2026, 2, 20, 0, 0, 0, 0, time.Local))
	library.ReturnBook(premium, book1, time.Date(2026, 2, 25, 0, 0, 0, 0, time.Local))
	fmt.Println()

	fmt.Println("=== After Returns ===")
	library.ShowCatalog()
	regular.ShowInfo()
	premium.ShowInfo()
}
