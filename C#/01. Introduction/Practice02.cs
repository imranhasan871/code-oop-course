/**
 * Practice 02: Find Min & Max Salary
 * Task: Read employee salary data from a CSV file and display the employee
 *       with the highest and lowest salary.
 *
 * How to run:
 *   Create a Console App, paste this code into Program.cs, and run.
 *   OR use: dotnet-script Practice02.cs
 *
 * Expected output:
 *   Max: Hasan   56000
 *   Min: Rafi    34000
 *
 * NOTE: The data file is at: ../../data/salarysheet.csv
 *       Make sure the working directory is C#/01. Introduction/ when running.
 */

using System;
using System.Collections.Generic;
using System.IO;

class Practice02
{
    /**
     * Inner class: represents one employee with name and salary.
     */
    class Employee
    {
        public string Name { get; set; }
        public double Salary { get; set; }

        public Employee(string name, double salary)
        {
            Name = name;
            Salary = salary;
        }
    }

    static void Main(string[] args)
    {
        string fileName = "../../data/salarysheet.csv";

        // Step 1: Read data from file
        List<Employee> salaryList = ReadDataFromCSVFile(fileName);

        if (salaryList == null || salaryList.Count == 0)
        {
            Console.WriteLine("No employee data found in the file.");
            return;
        }

        // Step 2: Find max and min salary employees
        Employee maxEmployee = FindMaxSalaryWithEmployeeName(salaryList);
        Employee minEmployee = FindMinSalaryWithEmployeeName(salaryList);

        // Step 3: Display the results
        Console.WriteLine($"Max: {maxEmployee.Name,-10} {maxEmployee.Salary:F0}");
        Console.WriteLine($"Min: {minEmployee.Name,-10} {minEmployee.Salary:F0}");
    }

    /**
     * Reads employee names and salaries from a tab-separated file.
     * Each line should have: EmployeeName TAB Salary
     */
    static List<Employee> ReadDataFromCSVFile(string fileName)
    {
        List<Employee> employees = new List<Employee>();

        try
        {
            string[] lines = File.ReadAllLines(fileName);
            foreach (string line in lines)
            {
                string trimmed = line.Trim();
                if (string.IsNullOrEmpty(trimmed)) continue; // skip blank lines

                // Split by tab first, then fallback to comma
                string[] parts = trimmed.Split('\t');
                if (parts.Length < 2)
                    parts = trimmed.Split(',');
                if (parts.Length < 2) continue; // skip malformed lines

                string name = parts[0].Trim();
                double salary = double.Parse(parts[1].Trim());
                employees.Add(new Employee(name, salary));
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error reading file: {ex.Message}");
        }

        return employees;
    }

    /**
     * Returns the employee with the highest salary.
     */
    static Employee FindMaxSalaryWithEmployeeName(List<Employee> employees)
    {
        Employee max = employees[0];
        foreach (Employee emp in employees)
        {
            if (emp.Salary > max.Salary)
                max = emp;
        }
        return max;
    }

    /**
     * Returns the employee with the lowest salary.
     */
    static Employee FindMinSalaryWithEmployeeName(List<Employee> employees)
    {
        Employee min = employees[0];
        foreach (Employee emp in employees)
        {
            if (emp.Salary < min.Salary)
                min = emp;
        }
        return min;
    }
}
