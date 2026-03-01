/**
 * Practice 02: Find Min & Max Salary
 * Task: Read employee salary data from a CSV file and display the employee
 *       with the highest and lowest salary.
 *
 * How to run:
 *   go run practice-02.go
 *
 * Expected output:
 *   Max: Hasan   56000
 *   Min: Rafi    34000
 *
 * NOTE: The data file is at: ../../data/salarysheet.csv
 *       Make sure you run this program from the Golang/01. Introduction/ directory.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/** Employee holds the name and salary of a single employee. */
type Employee struct {
	Name   string
	Salary float64
}

func main() {
	fileName := "../../data/salarysheet.csv"

	// Step 1: Read data from file
	salaryList, err := readDataFromCSVFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if len(salaryList) == 0 {
		fmt.Println("No employee data found in the file.")
		return
	}

	// Step 2: Find max and min salary employees
	maxEmployee := findMaxSalaryWithEmployeeName(salaryList)
	minEmployee := findMinSalaryWithEmployeeName(salaryList)

	// Step 3: Display the results
	fmt.Printf("Max: %-10s %.0f\n", maxEmployee.Name, maxEmployee.Salary)
	fmt.Printf("Min: %-10s %.0f\n", minEmployee.Name, minEmployee.Salary)
}

/**
 * readDataFromCSVFile reads employee names and salaries from a tab-separated file.
 * Each line should have: EmployeeName <tab> Salary
 */
func readDataFromCSVFile(fileName string) ([]Employee, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var employees []Employee

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // skip blank lines
		}

		// Split by tab first, then fallback to comma
		parts := strings.Split(line, "\t")
		if len(parts) < 2 {
			parts = strings.Split(line, ",")
		}
		if len(parts) < 2 {
			continue // skip malformed lines
		}

		name := strings.TrimSpace(parts[0])
		salary, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err != nil {
			continue // skip lines with invalid salary
		}

		employees = append(employees, Employee{Name: name, Salary: salary})
	}

	return employees, scanner.Err()
}

/**
 * findMaxSalaryWithEmployeeName returns the employee with the highest salary.
 */
func findMaxSalaryWithEmployeeName(employees []Employee) Employee {
	max := employees[0]
	for _, emp := range employees[1:] {
		if emp.Salary > max.Salary {
			max = emp
		}
	}
	return max
}

/**
 * findMinSalaryWithEmployeeName returns the employee with the lowest salary.
 */
func findMinSalaryWithEmployeeName(employees []Employee) Employee {
	min := employees[0]
	for _, emp := range employees[1:] {
		if emp.Salary < min.Salary {
			min = emp
		}
	}
	return min
}
