# Java — Getting Started

## Prerequisites

Install the Java Development Kit (JDK 11 or higher): <https://adoptium.net/>

Verify installation:

```bash
java -version
javac -version
```

---

## How to Run

Open a terminal and navigate to the practice folder:

```bash
cd "Java/01. Introduction"
```

Compile and run any practice:

```bash
# Practice 01
javac Practice01.java
java Practice01

# Practice 02
javac Practice02.java
java Practice02

# Practice 03
javac Practice03.java
java Practice03
```

> **Important:** Practices 02 and 03 read from the `data/` folder using a relative path `../../data/`.
> Always run them **from inside** `Java/01. Introduction/` so the path resolves correctly.

---

## Practice Files

| File | Topic |
|------|-------|
| [`Practice01.java`](01.%20Introduction/Practice01.java) | Partial name search in an ArrayList |
| [`Practice02.java`](01.%20Introduction/Practice02.java) | Read CSV → find Min & Max salary |
| [`Practice03.java`](01.%20Introduction/Practice03.java) | Generate & save a formatted voucher |

### Session 03 — Understanding Class & Object

| File | Topic |
|------|-------|
| [`Practice04.java`](03.%20Understanding%20Class%20%26%20Object/Practice04.java) | BankAccount class — deposit, withdraw, transfer |
| [`Practice05.java`](03.%20Understanding%20Class%20%26%20Object/Practice05.java) | CreditCard class — spending limits & cash rules |
| [`Practice06.java`](03.%20Understanding%20Class%20%26%20Object/Practice06.java) | Car Rental system (OOAD) |
| [`Practice07.java`](03.%20Understanding%20Class%20%26%20Object/Practice07.java) | School Homework system (OOAD) |
| [`Practice08.java`](03.%20Understanding%20Class%20%26%20Object/Practice08.java) | Movie Streaming platform (OOAD) |
| [`Practice09.java`](03.%20Understanding%20Class%20%26%20Object/Practice09.java) | Bank Account Collection — list of accounts |

### Session 04 — Association Relationship

| File | Topic |
|------|-------|
| [`Practice10.java`](04.%20Association%20Relationship/Practice10.java) | Customer & Credit Card — 1-1 Association |
| [`Practice11.java`](04.%20Association%20Relationship/Practice11.java) | Car & License Plate — 1-1 Association |
| [`Practice12.java`](04.%20Association%20Relationship/Practice12.java) | Doctor & Patients — 1-Many Association |
| [`Practice13.java`](04.%20Association%20Relationship/Practice13.java) | Patient & Medications — Many-Many Association |
| [`Practice14.java`](04.%20Association%20Relationship/Practice14.java) | Smart Parking System |
| [`Practice15.java`](04.%20Association%20Relationship/Practice15.java) | Restaurant Order Management |
| [`Practice16.java`](04.%20Association%20Relationship/Practice16.java) | Library Management System |

### Session 05 — Inheritance Relationship

| File | Topic |
|------|-------|
| [`Practice17.java`](05.%20Inheritance%20Relationship/Practice17.java) | Vehicle Rental — Inheritance (IS-A) |
| [`Practice18.java`](05.%20Inheritance%20Relationship/Practice18.java) | Method Overriding & Constructor Chaining |
| [`Practice19.java`](05.%20Inheritance%20Relationship/Practice19.java) | Runtime Polymorphism |

### Session 06 — Software Design Principles SOLID & More

| File | Topic |
|------|-------|
| [`Practice17.java`](06.%20Software%20Design%20Principles%20SOLID%20%26%20More/Practice17.java) | Applying SRP (Single Responsibility Principle) |
| [`Practice18.java`](06.%20Software%20Design%20Principles%20SOLID%20%26%20More/Practice18.java) | Applying DIP (Dependency Inversion Principle) |
| [`Practice19.java`](06.%20Software%20Design%20Principles%20SOLID%20%26%20More/Practice19.java) | Applying DRY (Don't Repeat Yourself) |

---

## Key Java Concepts Used

| Concept | Used In |
|---------|---------|
| `ArrayList<String>`, `List<T>` | Practice 01, 02, 03 |
| Static inner class | Practice 02, 03 |
| Enhanced `for` loop | All |
| `String.toLowerCase()`, `.contains()` | Practice 01 |
| `BufferedReader`, `FileReader` | Practice 02, 03 |
| `String.split("\\s+")` | Practice 03 |
| `LinkedHashMap` (duplicate detection) | Practice 03 |
| `StringBuilder` | Practice 03 |
| `PrintWriter`, `FileWriter` | Practice 03 |
