# Java — Getting Started

## Prerequisites

Install the Java Development Kit (JDK 11 or higher): https://adoptium.net/

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
