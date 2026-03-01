/**
 * Practice 01: Working with Array/ArrayList/List
 * Task: Create a list of employee names and find partial matches with a user input.
 *
 * How to compile and run:
 *   javac Practice01.java
 *   java Practice01
 *
 * Try changing the value of userInput in main() to test different searches:
 *   "pulok"  → Pulok
 *   "sa"     → Afsana, Samia
 *   "n"      → Afsana, Robin
 */

import java.util.ArrayList;
import java.util.List;

public class Practice01 {

    public static void main(String[] args) {
        // --- Data Setup ---
        // Create a list (ArrayList) of employee names
        List<String> employeeNames = new ArrayList<>();
        employeeNames.add("Tareq");
        employeeNames.add("Afsana");
        employeeNames.add("Imtiaz");
        employeeNames.add("Pulok");
        employeeNames.add("Robin");
        employeeNames.add("Samia");
        employeeNames.add("Rupok");

        // --- User Input (hard-coded for practice) ---
        // Change this value to test different search terms
        String userInput = "sa";

        // --- Find Matches ---
        List<String> results = findPartialMatch(employeeNames, userInput);

        // --- Show Result ---
        showResult(userInput, results);
    }

    /**
     * Searches the names list for any name that contains the input string.
     * The search is case-insensitive (e.g., "sa" matches "Afsana" and "Samia").
     */
    public static List<String> findPartialMatch(List<String> names, String input) {
        List<String> matches = new ArrayList<>();
        String inputLower = input.toLowerCase(); // normalize input to lowercase

        for (String name : names) {
            // Compare both in lowercase so the match is not case-sensitive
            if (name.toLowerCase().contains(inputLower)) {
                matches.add(name);
            }
        }

        return matches;
    }

    /**
     * Prints the search input and the list of matching names.
     */
    public static void showResult(String input, List<String> matches) {
        System.out.println("Search input : \"" + input + "\"");

        if (matches.isEmpty()) {
            System.out.println("Output       : No matching names found.");
            return;
        }

        System.out.println("Output       : " + String.join(", ", matches));
    }
}
