# Professional OOP Training Program
### Trainer: Zohirul Alam Tiemoon
[learnwithtiemoon.com/professional-oop](https://learnwithtiemoon.com/professional-oop)

---

## About This Repository

This repository contains the **practice exercises** for the Professional OOP Training Program conducted by **Zohirul Alam Tiemoon**, Co-founder & CEO of Nerd Castle Limited.

Each practice is implemented in **four languages** side by side so students can compare syntax and structure while learning the same concept:

| Language | Folder |
|----------|--------|
| Go       | [`Golang/`](Golang/) |
| Java     | [`Java/`](Java/) |
| C#       | [`C#/`](C#/) |
| Python   | [`Python/`](Python/) |

---

## Session 01 — Language Basics

> **Goal:** Reinforce fundamentals before diving into OOP — data structures, file I/O, and iteration.

### Practice 01 — Working with List / ArrayList
**Duration:** 15 min

Create a list of employee names and find all names that **partially match** a given search input (case-insensitive).

| Input | Output |
|-------|--------|
| `pulok` | Pulok |
| `sa` | Afsana, Samia |
| `n` | Afsana, Robin |

**Key concepts:** `List`, iteration, `String` contains, case-insensitive comparison

---

### Practice 02 — Find Min & Max Salary
**Duration:** 15 min

Read employee salary data from [`data/salarysheet.csv`](data/salarysheet.csv) and display the **employee with the highest and lowest salary**.

```
Max: Hasan    56000
Min: Rafi     34000
```

**Key concepts:** File I/O, CSV parsing, iteration, comparison logic

---

### Practice 03 — Generate a Voucher
**Duration:** 1 hour (Homework)

Read grocery items from [`data/groceryitems.csv`](data/groceryitems.csv), calculate totals, and generate a formatted voucher — displayed on console **and** saved to a `.txt` file.

```
Item id      Qty    unit price    total
item-937     12     230.5         2766.0
item-432     3      120.0         360.0
item-431     1      1230.0        1230.0
item-098     5      12.5          62.5
item-133     8      30.0          240.0

Total        29     Grand total   4658.5
VAT (5%)     232.925
Net total    4425.575
```

> If **duplicate item IDs** are found in the file, the program must display an error message naming the duplicates and refuse to generate the voucher.

**Key concepts:** File I/O, data validation, duplicate detection, formatted output, writing to file

---

## Session 03 — Understanding Class & Object

> **Goal:** Learn how to model real-world entities as classes, apply encapsulation, and think in an object-oriented way.

### Practice 04 — Bank Account
**Duration:** 10 min

Create a `BankAccount` class with properties for account number, account name, and balance. Methods: **deposit**, **withdraw**, and **transfer**. No negative balance allowed.

**Key concepts:** Class/struct, constructor, encapsulation, business rules

---

### Practice 05 — Credit Card
**Duration:** 15 min

Design a `CreditCard` class with a max limit of **500K**. Cash withdrawal: daily limit **100K**, per-transaction limit **20K**. Bill payments: no per-transaction limit, total spending must not exceed the max limit.

**Key concepts:** Private fields, multiple business rules, cumulative tracking

---

### Practice 06 — Car Rental System (OOAD)

Map a real-world car rental scenario to code: `Car`, `Customer`, `RentalCompany` classes. Track fleet availability, rental rates, and maintenance schedules.

**Key concepts:** OOAD, multiple classes collaborating, object collections, constructors

---

### Practice 07 — School Homework System (OOAD)

Map a school homework scenario: `Homework`, `Student`, `Submission`, `Teacher` classes. Teachers create & assign homework, students submit, teachers grade.

**Key concepts:** OOAD, encapsulation, object references, collection management

---

### Practice 08 — Movie Streaming Platform (OOAD)

Map a streaming platform: `Movie`, `User`, `StreamingPlatform` classes. Users browse movies, manage watchlists, rate films, and get genre-based recommendations.

**Key concepts:** OOAD, sets/maps for lookups, business logic (recommendations)

---

### Practice 09 — Bank Account Collection

Extended Practice 04 — create a `Bank` class that owns a **list of BankAccounts**. Perform transactions across multiple accounts and calculate the **total balance** of the bank.

**Key concepts:** Object collection (List of objects), aggregation, iteration

---

## Project Structure

```
code-oop-course/
├── data/
│   ├── salarysheet.csv       # Input data for Practice 02
│   └── groceryitems.csv      # Input data for Practice 03
│
├── Golang/
│   ├── 01. Introduction/
│   │   ├── practice-01.go
│   │   ├── practice-02.go
│   │   └── practice-03.go
│   └── 03. Understanding Class & Object/
│       ├── practice-04.go
│       ├── practice-05.go
│       ├── practice-06.go
│       ├── practice-07.go
│       ├── practice-08.go
│       └── practice-09.go
│
├── Java/
│   ├── 01. Introduction/
│   │   ├── Practice01.java
│   │   ├── Practice02.java
│   │   └── Practice03.java
│   └── 03. Understanding Class & Object/
│       ├── Practice04.java
│       ├── Practice05.java
│       ├── Practice06.java
│       ├── Practice07.java
│       ├── Practice08.java
│       └── Practice09.java
│
├── C#/
│   ├── 01. Introduction/
│   │   ├── Practice01.cs
│   │   ├── Practice02.cs
│   │   └── Practice03.cs
│   └── 03. Understanding Class & Object/
│       ├── Practice04.cs
│       ├── Practice05.cs
│       ├── Practice06.cs
│       ├── Practice07.cs
│       ├── Practice08.cs
│       └── Practice09.cs
│
└── Python/
    ├── 01. Introduction/
    │   ├── practice_01.py
    │   ├── practice_02.py
    │   └── practice_03.py
    └── 03. Understanding Class & Object/
        ├── practice_04.py
        ├── practice_05.py
        ├── practice_06.py
        ├── practice_07.py
        ├── practice_08.py
        └── practice_09.py
```

---

## How to Run

See the individual language README for setup and run instructions:

- [Go — Getting Started](Golang/README.md)
- [Java — Getting Started](Java/README.md)
- [C# — Getting Started](C#/README.md)
- [Python — Getting Started](Python/README.md)

---

## About the Trainer

**Zohirul Alam Tiemoon**
Co-founder & CEO, Nerd Castle Limited

- 24+ years of experience in software development
- Former SVP (Technology) Goama Pte Ltd & CTO of Databiz Software Ltd
- Trained 13,000+ software engineers
- Specialist in OOP, Design Principles & Patterns, and scalable system design
- LinkedIn: [linkedin.com/in/tiemoon](https://linkedin.com/in/tiemoon)
- Website: [learnwithtiemoon.com](https://learnwithtiemoon.com)
