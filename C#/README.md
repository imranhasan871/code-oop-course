# C# — Getting Started

## Prerequisites

Install the .NET SDK (version 6 or higher): https://dotnet.microsoft.com/download

Verify installation:
```bash
dotnet --version
```

---

## How to Run

### Option A — Console App (Recommended)

1. Create a new Console App project:
   ```bash
   dotnet new console -n Practice01
   cd Practice01
   ```
2. Replace the contents of `Program.cs` with the code from the practice file.
3. Run:
   ```bash
   dotnet run
   ```

### Option B — dotnet-script (Run `.cs` files directly)

Install `dotnet-script` globally:
```bash
dotnet tool install -g dotnet-script
```

Then run any `.cs` file directly from `C#/01. Introduction/`:
```bash
cd "C#/01. Introduction"
dotnet-script Practice01.cs
dotnet-script Practice02.cs
dotnet-script Practice03.cs
```

> **Important:** Practices 02 and 03 read from the `data/` folder using a relative path `../../data/`.
> Always run them **from inside** `C#/01. Introduction/` so the path resolves correctly.

---

## Practice Files

| File | Topic |
|------|-------|
| [`Practice01.cs`](01.%20Introduction/Practice01.cs) | Partial name search in a List |
| [`Practice02.cs`](01.%20Introduction/Practice02.cs) | Read CSV → find Min & Max salary |
| [`Practice03.cs`](01.%20Introduction/Practice03.cs) | Generate & save a formatted voucher |

---

## Key C# Concepts Used

| Concept | Used In |
|---------|---------|
| `List<string>`, `List<T>` | Practice 01, 02, 03 |
| Nested class with properties | Practice 02, 03 |
| `foreach` loop | All |
| `string.ToLower()`, `.Contains()` | Practice 01 |
| `File.ReadAllLines()` | Practice 02, 03 |
| `string.Split()` with `RemoveEmptyEntries` | Practice 03 |
| `Dictionary<string, int>` (duplicate detection) | Practice 03 |
| `StringBuilder` | Practice 03 |
| `File.WriteAllText()` | Practice 03 |
| Expression-bodied property (`=>`) | Practice 03 |
