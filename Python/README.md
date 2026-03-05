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
