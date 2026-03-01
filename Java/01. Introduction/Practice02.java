/**
 * Practice 02: Find Min & Max Salary
 * Task: Read employee salary data from a CSV file and display the employee
 *       with the highest and lowest salary.
 *
 * How to compile and run:
 *   javac Practice02.java
 *   java Practice02
 *
 * Expected output:
 *   Max: Hasan   56000
 *   Min: Rafi    34000
 *
 * NOTE: The data file is at: ../../data/salarysheet.csv
 *       Make sure you run this program from the Java/01. Introduction/ directory.
 */

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Practice02 {

    /**
     * Inner class: represents one employee with name and salary.
     */
    static class Employee {
        String name;
        double salary;

        Employee(String name, double salary) {
            this.name = name;
            this.salary = salary;
        }
    }

    public static void main(String[] args) {
        String fileName = "../../data/salarysheet.csv";

        // Step 1: Read data from file
        List<Employee> salaryList = readDataFromCSVFile(fileName);

        if (salaryList == null || salaryList.isEmpty()) {
            System.out.println("No employee data found in the file.");
            return;
        }

        // Step 2: Find max and min salary employees
        Employee maxEmployee = findMaxSalaryWithEmployeeName(salaryList);
        Employee minEmployee = findMinSalaryWithEmployeeName(salaryList);

        // Step 3: Display the results
        System.out.printf("Max: %-10s %.0f%n", maxEmployee.name, maxEmployee.salary);
        System.out.printf("Min: %-10s %.0f%n", minEmployee.name, minEmployee.salary);
    }

    /**
     * Reads employee names and salaries from a tab-separated file.
     * Each line should have: EmployeeName <tab> Salary
     */
    public static List<Employee> readDataFromCSVFile(String fileName) {
        List<Employee> employees = new ArrayList<>();

        try (BufferedReader reader = new BufferedReader(new FileReader(fileName))) {
            String line;
            while ((line = reader.readLine()) != null) {
                line = line.trim();
                if (line.isEmpty()) continue; // skip blank lines

                // Split by tab first, then fallback to comma
                String[] parts = line.split("\t");
                if (parts.length < 2) {
                    parts = line.split(",");
                }
                if (parts.length < 2) continue; // skip malformed lines

                String name = parts[0].trim();
                double salary = Double.parseDouble(parts[1].trim());
                employees.add(new Employee(name, salary));
            }
        } catch (IOException e) {
            System.out.println("Error reading file: " + e.getMessage());
        }

        return employees;
    }

    /**
     * Returns the employee with the highest salary.
     */
    public static Employee findMaxSalaryWithEmployeeName(List<Employee> employees) {
        Employee max = employees.get(0);
        for (Employee emp : employees) {
            if (emp.salary > max.salary) {
                max = emp;
            }
        }
        return max;
    }

    /**
     * Returns the employee with the lowest salary.
     */
    public static Employee findMinSalaryWithEmployeeName(List<Employee> employees) {
        Employee min = employees.get(0);
        for (Employee emp : employees) {
            if (emp.salary < min.salary) {
                min = emp;
            }
        }
        return min;
    }
}
