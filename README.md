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

## Session 04 — Association Relationship

> **Goal:** Understand and implement different types of associations — 1-to-1, 1-to-Many, and Many-to-Many — along with real-world system designs that combine multiple association patterns.

### Practice 10 — Customer & Credit Card (1-1 Association)

Model a bank customer who has exactly one credit card. Card number must be digits only, expiration date must be validated, 500K credit limit, track available credit after purchases. Customer must be at least 18 years old.

**Key concepts:** 1-to-1 Association, input validation, date-based expiration

---

### Practice 11 — Car & License Plate (1-1 Association)

Model a car with its license plate. Validate plate expiration. Track owner, manufacturer, model, year. Car age must be ≤ 20 years for registration renewal.

**Key concepts:** 1-to-1 Association, age-based eligibility, date handling

---

### Practice 12 — Doctor & Patients (1-Many Association)

Model a hospital system where a doctor manages multiple patients. Doctors can schedule appointments, diagnose, prescribe treatments, and discharge patients. Patients can view appointments and track medical history.

**Key concepts:** 1-to-Many Association, list management, business workflow

---

### Practice 13 — Patient & Medications (Many-Many Association)

Model a prescription system where patients can be prescribed multiple medications and each medication can be prescribed to multiple patients. A `Prescription` junction class links them. Forbidden medication combinations must be enforced:

- Antibiotics & Statins
- Muscle Relaxants & CNS Depressants
- Anti-Inflammatories & Anticoagulants

**Key concepts:** Many-to-Many Association, junction class, conflict detection

---

### Practice 14 — Smart Parking System

Manage vehicle entry, slot assignment, duration-based charges, and exit. Track real-time slot availability. Charge based on parking duration and vehicle type.

**Key concepts:** Multiple associations (Vehicle, ParkingSlot, ParkingLot), time-based calculations

---

### Practice 15 — Restaurant Order Management

Customers place orders with multiple menu items. Calculate total bill. Kitchen updates order status through Pending → Cooking → Served. Payment after order is served.

**Key concepts:** Multiple associations (MenuItem, Order, Kitchen, Restaurant), state management

---

### Practice 16 — Library Management System

Library with books (each having copies). Members borrow books — Regular members max 3, Premium max 5. Late return fines: 10 taka/day (Regular), 5 taka/day (Premium).

**Key concepts:** Multiple associations (Book, Member, BorrowRecord, Library), membership-based rules, fine calculation

---

## Session 05 — Inheritance Relationship

> **Goal:** Learn inheritance (IS-A relationship), method overriding, constructor chaining, upcasting, downcasting, and runtime polymorphism.

### Practice 17 — Vehicle Rental System (Inheritance)

**Duration:** 30 min

A vehicle rental company rents out Cars, Bikes, and Trucks. All vehicles share Brand, Model, and Year. Each type has different rental pricing:

- **Car:** $50/day. 10% discount if older than 5 years.
- **Bike:** $15/day. 15% discount if rental period > 7 days.
- **Truck:** $100/day + $100/day additional maintenance fee.

**Key concepts:** Inheritance, IS-A relationship, base class, subclass

---

### Practice 18 — Method Overriding & Constructor Chaining

**Duration:** 20 min

Apply method overriding and constructor chaining to the vehicle rental system from Practice 17:

- Constructor chaining between base class and subclasses (`super()` / `base()`)
- Override `calculateRentalCost()` in each subclass
- Override `toString()` / `__str__()` for string representation

**Key concepts:** Method overriding, constructor chaining, `@Override` / `virtual` + `override`, `toString()`

---

### Practice 19 — Runtime Polymorphism

**Duration:** 15 min

Display the brand, model, year of manufacture, and rental price for one car, two bikes, and one truck — each rented separately for 10 days. Use a single collection of Vehicle references.

**Key concepts:** Upcasting, downcasting, runtime polymorphism, method dispatch

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
│   ├── 04. Association Relationship/
│   │   ├── practice-10.go
│   │   ├── practice-11.go
│   │   ├── practice-12.go
│   │   ├── practice-13.go
│   │   ├── practice-14.go
│   │   ├── practice-15.go
│   │   └── practice-16.go
│   └── 05. Inheritance Relationship/
│       ├── practice-17.go
│       ├── practice-18.go
│       └── practice-19.go
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
│   ├── 04. Association Relationship/
│   │   ├── Practice10.java
│   │   ├── Practice11.java
│   │   ├── Practice12.java
│   │   ├── Practice13.java
│   │   ├── Practice14.java
│   │   ├── Practice15.java
│   │   └── Practice16.java
│   └── 05. Inheritance Relationship/
│       ├── Practice17.java
│       ├── Practice18.java
│       └── Practice19.java
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
│   ├── 04. Association Relationship/
│   │   ├── Practice10.cs
│   │   ├── Practice11.cs
│   │   ├── Practice12.cs
│   │   ├── Practice13.cs
│   │   ├── Practice14.cs
│   │   ├── Practice15.cs
│   │   └── Practice16.cs
│   └── 05. Inheritance Relationship/
│       ├── Practice17.cs
│       ├── Practice18.cs
│       └── Practice19.cs
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
    ├── 04. Association Relationship/
    │   ├── practice_10.py
    │   ├── practice_11.py
    │   ├── practice_12.py
    │   ├── practice_13.py
    │   ├── practice_14.py
    │   ├── practice_15.py
    │   └── practice_16.py
    └── 05. Inheritance Relationship/
        ├── practice_17.py
        ├── practice_18.py
        └── practice_19.py
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
