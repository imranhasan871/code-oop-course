"""
Practice 01: Working with Array/ArrayList/List
Task: Create a list of employee names and find partial matches with a user input.

How to run:
  python practice_01.py

Try changing the value of user_input to test different searches:
  "pulok"  -> Pulok
  "sa"     -> Afsana, Samia
  "n"      -> Afsana, Robin

Course: Professional OOP — by Zohirul Alam Tiemoon
"""


def find_partial_match(names: list, user_input: str) -> list:
    """
    Searches the names list for any name that contains the input string.
    The search is case-insensitive (e.g., "sa" matches "Afsana" and "Samia").
    """
    matches = []
    input_lower = user_input.lower()

    for name in names:
        if input_lower in name.lower():
            matches.append(name)

    return matches


def show_result(user_input: str, matches: list):
    """Prints the search input and the list of matching names."""
    print(f'Search input : "{user_input}"')

    if not matches:
        print("Output       : No matching names found.")
        return

    print(f"Output       : {', '.join(matches)}")


if __name__ == "__main__":
    # --- Data Setup ---
    # Create a list of employee names
    employee_names = [
        "Tareq",
        "Afsana",
        "Imtiaz",
        "Pulok",
        "Robin",
        "Samia",
        "Rupok",
    ]

    # --- User Input (hard-coded for practice) ---
    # Change this value to test different search terms
    user_input = "sa"

    # --- Find Matches ---
    results = find_partial_match(employee_names, user_input)

    # --- Show Result ---
    show_result(user_input, results)
