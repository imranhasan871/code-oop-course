# Python — Getting Started

## Prerequisites

- **Python 3.8+** — download from [python.org](https://www.python.org/downloads/)
- Verify installation: `python --version`

---

## How to Run

Navigate to the session folder and run the practice file directly:

```bash
# Session 01 — Introduction
cd Python/01.\ Introduction/
python practice_01.py
python practice_02.py
python practice_03.py

# Session 03 — Understanding Class & Object
cd Python/03.\ Understanding\ Class\ \&\ Object/
python practice_04.py
python practice_05.py
python practice_06.py
python practice_07.py
python practice_08.py
python practice_09.py

# Session 04 — Association Relationship
cd Python/04.\ Association\ Relationship/
python practice_10.py
python practice_11.py
python practice_12.py
python practice_13.py
python practice_14.py
python practice_15.py
python practice_16.py

# Session 05 — Inheritance Relationship
cd Python/05.\ Inheritance\ Relationship/
python practice_17.py
python practice_18.py
python practice_19.py

# Session 06 — Software Design Principles SOLID & More
cd Python/06.\ Software\ Design\ Principles\ SOLID\ \&\ More/
python practice_17_srp.py
python practice_18_dip.py
python practice_19_dry.py
```

> **Note:** Practice 02 and 03 read data files from `../../data/`, so make sure you run them from inside the session folder.

---

## Practice Files

### Session 01 — Introduction

| File | Practice | Description |
|------|----------|-------------|
| `practice_01.py` | 01 | Partial name search in a list |
| `practice_02.py` | 02 | Find min & max salary from CSV |
| `practice_03.py` | 03 | Generate a grocery voucher (console + file) |

### Session 03 — Understanding Class & Object

| File | Practice | Description |
|------|----------|-------------|
| `practice_04.py` | 04 | BankAccount class — deposit, withdraw, transfer |
| `practice_05.py` | 05 | CreditCard class — spending limits & cash rules |
| `practice_06.py` | 06 | Car Rental system (OOAD) |
| `practice_07.py` | 07 | School Homework system (OOAD) |
| `practice_08.py` | 08 | Movie Streaming platform (OOAD) |
| `practice_09.py` | 09 | Bank Account Collection — list of accounts |

### Session 04 — Association Relationship

| File | Practice | Description |
|------|----------|-------------|
| `practice_10.py` | 10 | Customer & Credit Card — 1-1 Association |
| `practice_11.py` | 11 | Car & License Plate — 1-1 Association |
| `practice_12.py` | 12 | Doctor & Patients — 1-Many Association |
| `practice_13.py` | 13 | Patient & Medications — Many-Many Association |
| `practice_14.py` | 14 | Smart Parking System |
| `practice_15.py` | 15 | Restaurant Order Management |
| `practice_16.py` | 16 | Library Management System |

### Session 05 — Inheritance Relationship

| File | Practice | Description |
|------|----------|-------------|
| `practice_17.py` | 17 | Vehicle Rental — Inheritance (IS-A) |
| `practice_18.py` | 18 | Method Overriding & Constructor Chaining |
| `practice_19.py` | 19 | Runtime Polymorphism |

### Session 06 — Software Design Principles SOLID & More

| File | Practice | Description |
|------|----------|-------------|
| `practice_17_srp.py` | 17 | SRP refactor for OnDemandAgentService |
| `practice_18_dip.py` | 18 | DIP-based payment processing system |
| `practice_19_dry.py` | 19 | DRY refactor for invoice and discount logic |

---

## Key Python Concepts Used

| Concept | Python Feature |
|---------|---------------|
| Class definition | `class ClassName:` |
| Constructor | `__init__(self, ...)` |
| Private attributes | `_attribute` (convention) |
| Properties | `@property` decorator |
| Lists | `list` (built-in) |
| Dictionaries | `dict` (built-in) |
| Sets | `set` (built-in) |
| String formatting | f-strings `f"..."` |
| File I/O | `open()`, `with` statement |
| Type hints | `def func(x: int) -> str:` |
