/**
 * Practice 01: Working with Array/ArrayList/List
 * Task: Create a list of employee names and find partial matches with a user input.
 *
 * How to run:
 *   go run practice-01.go
 *
 * Try changing the value of userInput to test different searches:
 *   "pulok"  → Pulok
 *   "sa"     → Afsana, Samia
 *   "n"      → Afsana, Robin
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	// --- Data Setup ---
	// Create a list (slice) of employee names
	employeeNames := []string{
		"Tareq",
		"Afsana",
		"Imtiaz",
		"Pulok",
		"Robin",
		"Samia",
		"Rupok",
	}

	// --- User Input (hard-coded for practice) ---
	// Change this value to test different search terms
	userInput := "sa"

	// --- Find Matches ---
	results := findPartialMatch(employeeNames, userInput)

	// --- Show Result ---
	showResult(userInput, results)
}

/**
 * findPartialMatch searches the names list for any name that contains the input string.
 * The search is case-insensitive (e.g., "sa" matches "Afsana" and "Samia").
 */
func findPartialMatch(names []string, input string) []string {
	var matches []string
	inputLower := strings.ToLower(input) // normalize input to lowercase

	for _, name := range names {
		// Compare both in lowercase so the match is not case-sensitive
		if strings.Contains(strings.ToLower(name), inputLower) {
			matches = append(matches, name)
		}
	}

	return matches
}

/**
 * showResult prints the search input and the list of matching names.
 */
func showResult(input string, matches []string) {
	fmt.Printf("Search input : %q\n", input)

	if len(matches) == 0 {
		fmt.Println("Output       : No matching names found.")
		return
	}

	fmt.Printf("Output       : %s\n", strings.Join(matches, ", "))
}
