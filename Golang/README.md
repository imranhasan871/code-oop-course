# Go — Getting Started

## Prerequisites

Install Go from the official website: <https://go.dev/dl/>

Verify installation:

```bash
go version
# go version go1.22.x ...
```

---

## How to Run

Open a terminal and navigate to the practice folder:

```bash
cd "Golang/01. Introduction"
```

Then run any practice file directly:

```bash
go run practice-01.go
go run practice-02.go
go run practice-03.go
```

> **Important:** Practices 02 and 03 read from the `data/` folder using a relative path `../../data/`.
> Always run them **from inside** `Golang/01. Introduction/` so the path resolves correctly.

---

## Practice Files

| File | Topic |
|------|-------|
| [`practice-01.go`](01.%20Introduction/practice-01.go) | Partial name search in a List (slice) |
| [`practice-02.go`](01.%20Introduction/practice-02.go) | Read CSV → find Min & Max salary |
| [`practice-03.go`](01.%20Introduction/practice-03.go) | Generate & save a formatted voucher |

### Session 03 — Understanding Class & Object

| File | Topic |
|------|-------|
| [`practice-04.go`](03.%20Understanding%20Class%20%26%20Object/practice-04.go) | BankAccount class — deposit, withdraw, transfer |
| [`practice-05.go`](03.%20Understanding%20Class%20%26%20Object/practice-05.go) | CreditCard class — spending limits & cash rules |
| [`practice-06.go`](03.%20Understanding%20Class%20%26%20Object/practice-06.go) | Car Rental system (OOAD) |
| [`practice-07.go`](03.%20Understanding%20Class%20%26%20Object/practice-07.go) | School Homework system (OOAD) |
| [`practice-08.go`](03.%20Understanding%20Class%20%26%20Object/practice-08.go) | Movie Streaming platform (OOAD) |
| [`practice-09.go`](03.%20Understanding%20Class%20%26%20Object/practice-09.go) | Bank Account Collection — list of accounts |

### Session 04 — Association Relationship

| File | Topic |
|------|-------|
| [`practice-10.go`](04.%20Association%20Relationship/practice-10.go) | Customer & Credit Card — 1-1 Association |
| [`practice-11.go`](04.%20Association%20Relationship/practice-11.go) | Car & License Plate — 1-1 Association |
| [`practice-12.go`](04.%20Association%20Relationship/practice-12.go) | Doctor & Patients — 1-Many Association |
| [`practice-13.go`](04.%20Association%20Relationship/practice-13.go) | Patient & Medications — Many-Many Association |
| [`practice-14.go`](04.%20Association%20Relationship/practice-14.go) | Smart Parking System |
| [`practice-15.go`](04.%20Association%20Relationship/practice-15.go) | Restaurant Order Management |
| [`practice-16.go`](04.%20Association%20Relationship/practice-16.go) | Library Management System |

### Session 05 — Inheritance Relationship

| File | Topic |
|------|-------|
| [`practice-17.go`](05.%20Inheritance%20Relationship/practice-17.go) | Vehicle Rental — Inheritance (IS-A) |
| [`practice-18.go`](05.%20Inheritance%20Relationship/practice-18.go) | Method Overriding & Constructor Chaining |
| [`practice-19.go`](05.%20Inheritance%20Relationship/practice-19.go) | Runtime Polymorphism |

### Session 06 — Software Design Principles SOLID & More

| File | Topic |
|------|-------|
| [`practice-17.go`](06.%20Software%20Design%20Principles%20SOLID%20%26%20More/practice-17.go) | Applying SRP (Single Responsibility Principle) |
| [`practice-18.go`](06.%20Software%20Design%20Principles%20SOLID%20%26%20More/practice-18.go) | Applying DIP (Dependency Inversion Principle) |
| [`practice-19.go`](06.%20Software%20Design%20Principles%20SOLID%20%26%20More/practice-19.go) | Applying DRY (Don't Repeat Yourself) |

---

## Key Go Concepts Used

| Concept | Used In |
|---------|---------|
| Slice (`[]string`, `[]Employee`) | Practice 01, 02, 03 |
| Struct | Practice 02, 03 |
| `for range` loop | All |
| `strings.ToLower`, `strings.Contains` | Practice 01 |
| `os.Open`, `bufio.Scanner` | Practice 02, 03 |
| `strings.Fields` (split by whitespace) | Practice 03 |
| `map[string]int` (duplicate detection) | Practice 03 |
| `strings.Builder` (build output string) | Practice 03 |
| `os.Create`, file write | Practice 03 |
