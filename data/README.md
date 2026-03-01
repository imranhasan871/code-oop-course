# Data Files

This folder contains the shared input data files used by the practice programs.

---

## salarysheet.csv

Used by **Practice 02** across all three languages.

**Format:** Tab-separated — `EmployeeName <TAB> Salary`

```
Dolon    45000
Rafi     34000
Hasan    56000
Salma    45000
```

**Expected output:**
```
Max: Hasan    56000
Min: Rafi     34000
```

---

## groceryitems.csv

Used by **Practice 03** across all three languages.

**Format:** Tab-separated — `ItemID <TAB> Quantity <TAB> UnitPrice`

```
item-937    12    230.5
item-432    3     120.0
item-431    1     1230.0
item-098    5     12.5
item-133    8     30.0
```

**Expected voucher output:**
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

> To test the **duplicate detection** feature of Practice 03, add a repeated `item-id` in `groceryitems.csv` and re-run the program.
