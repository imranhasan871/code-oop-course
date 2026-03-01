# Go — Getting Started

## Prerequisites

Install Go from the official website: https://go.dev/dl/

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
