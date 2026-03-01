# Professional OOP Training Program
### Trainer: Zohirul Alam Tiemoon
[learnwithtiemoon.com/professional-oop](https://learnwithtiemoon.com/professional-oop)

---

## About This Repository

This repository contains the **practice exercises** for the Professional OOP Training Program conducted by **Zohirul Alam Tiemoon**, Co-founder & CEO of Nerd Castle Limited.

Each practice is implemented in **three languages** side by side so students can compare syntax and structure while learning the same concept:

| Language | Folder |
|----------|--------|
| Go       | [`Golang/`](Golang/) |
| Java     | [`Java/`](Java/) |
| C#       | [`C#/`](C#/) |

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

## Project Structure

```
code-oop-course/
├── data/
│   ├── salarysheet.csv       # Input data for Practice 02
│   └── groceryitems.csv      # Input data for Practice 03
│
├── Golang/
│   └── 01. Introduction/
│       ├── practice-01.go
│       ├── practice-02.go
│       └── practice-03.go
│
├── Java/
│   └── 01. Introduction/
│       ├── Practice01.java
│       ├── Practice02.java
│       └── Practice03.java
│
└── C#/
    └── 01. Introduction/
        ├── Practice01.cs
        ├── Practice02.cs
        └── Practice03.cs
```

---

## How to Run

See the individual language README for setup and run instructions:

- [Go — Getting Started](Golang/README.md)
- [Java — Getting Started](Java/README.md)
- [C# — Getting Started](C#/README.md)

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
