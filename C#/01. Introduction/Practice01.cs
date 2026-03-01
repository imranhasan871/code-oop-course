/**
 * Practice 01: Working with Array/ArrayList/List
 * Task: Create a list of employee names and find partial matches with a user input.
 *
 * How to run:
 *   Create a Console App, paste this code into Program.cs, and run.
 *   OR use: dotnet-script Practice01.cs
 *
 * Try changing the value of userInput in Main() to test different searches:
 *   "pulok"  → Pulok
 *   "sa"     → Afsana, Samia
 *   "n"      → Afsana, Robin
 */

using System;
using System.Collections.Generic;

class Practice01
{
    static void Main(string[] args)
    {
        // --- Data Setup ---
        // Create a List of employee names
        List<string> employeeNames = new List<string>
        {
            "Tareq",
            "Afsana",
            "Imtiaz",
            "Pulok",
            "Robin",
            "Samia",
            "Rupok"
        };

        // --- User Input (hard-coded for practice) ---
        // Change this value to test different search terms
        string userInput = "sa";

        // --- Find Matches ---
        List<string> results = FindPartialMatch(employeeNames, userInput);

        // --- Show Result ---
        ShowResult(userInput, results);
    }

    /**
     * Searches the names list for any name that contains the input string.
     * The search is case-insensitive (e.g., "sa" matches "Afsana" and "Samia").
     */
    static List<string> FindPartialMatch(List<string> names, string input)
    {
        List<string> matches = new List<string>();
        string inputLower = input.ToLower(); // normalize input to lowercase

        foreach (string name in names)
        {
            // Compare both in lowercase so the match is not case-sensitive
            if (name.ToLower().Contains(inputLower))
            {
                matches.Add(name);
            }
        }

        return matches;
    }

    /**
     * Prints the search input and the list of matching names.
     */
    static void ShowResult(string input, List<string> matches)
    {
        Console.WriteLine($"Search input : \"{input}\"");

        if (matches.Count == 0)
        {
            Console.WriteLine("Output       : No matching names found.");
            return;
        }

        Console.WriteLine($"Output       : {string.Join(", ", matches)}");
    }
}
