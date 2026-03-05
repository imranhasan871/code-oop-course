"""
Practice 02: Find Min & Max Salary
Task: Read employee salary data from a CSV file and display the employee
      with the highest and lowest salary.

How to run:
  python practice_02.py

Expected output:
  Max: Hasan      56000
  Min: Rafi       34000

NOTE: The data file is at: ../../data/salarysheet.csv
      Make sure you run this program from the Python/01. Introduction/ directory.

Course: Professional OOP — by Zohirul Alam Tiemoon
"""


class Employee:
    """Holds the name and salary of a single employee."""

    def __init__(self, name: str, salary: float):
        self.name = name
        self.salary = salary


def read_data_from_csv_file(file_name: str) -> list:
    """
    Reads employee names and salaries from a tab-separated file.
    Each line should have: EmployeeName <tab> Salary
    """
    employees = []

    with open(file_name, "r") as file:
        for line in file:
            line = line.strip()
            if not line:
                continue  # skip blank lines

            # Split by tab first, then fallback to comma
            parts = line.split("\t")
            if len(parts) < 2:
                parts = line.split(",")
            if len(parts) < 2:
                continue  # skip malformed lines

            name = parts[0].strip()
            try:
                salary = float(parts[1].strip())
            except ValueError:
                continue  # skip lines with invalid salary

            employees.append(Employee(name, salary))

    return employees


def find_max_salary_with_employee_name(employees: list) -> Employee:
    """Returns the employee with the highest salary."""
    max_emp = employees[0]
    for emp in employees[1:]:
        if emp.salary > max_emp.salary:
            max_emp = emp
    return max_emp


def find_min_salary_with_employee_name(employees: list) -> Employee:
    """Returns the employee with the lowest salary."""
    min_emp = employees[0]
    for emp in employees[1:]:
        if emp.salary < min_emp.salary:
            min_emp = emp
    return min_emp


if __name__ == "__main__":
    file_name = "../../data/salarysheet.csv"

    # Step 1: Read data from file
    salary_list = read_data_from_csv_file(file_name)

    if not salary_list:
        print("No employee data found in the file.")
    else:
        # Step 2: Find max and min salary employees
        max_employee = find_max_salary_with_employee_name(salary_list)
        min_employee = find_min_salary_with_employee_name(salary_list)

        # Step 3: Display the results
        print(f"Max: {max_employee.name:<10} {max_employee.salary:.0f}")
        print(f"Min: {min_employee.name:<10} {min_employee.salary:.0f}")
