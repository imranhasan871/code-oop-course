"""
Practice 07: School Homework System (OOAD)
Task: Model a school homework system where teachers create homework,
      students submit answers, and teachers grade the submissions.

How to run:
  python practice_07.py

Key Concepts:
  - Multiple classes collaborating (Teacher, Student, Homework, Submission)
  - Methods that create and return objects
  - State tracking (graded/ungraded)
"""


class Homework:
    """Homework represents an assignment created by a teacher."""

    def __init__(self, title: str, subject: str, due_date: str, description: str, max_score: int):
        """Creates a new Homework assignment."""
        self.title = title
        self.subject = subject
        self.due_date = due_date
        self.description = description
        self.max_score = max_score

    def show_info(self):
        """Prints the homework details."""
        print(f"  Homework    : {self.title}")
        print(f"  Subject     : {self.subject}")
        print(f"  Due Date    : {self.due_date}")
        print(f"  Description : {self.description}")
        print(f"  Max Score   : {self.max_score}")
        print()


class Submission:
    """Submission represents a student's submitted answer for a homework."""

    def __init__(self, student, homework: Homework, answer: str):
        """Creates a new Submission."""
        self.student = student
        self.homework = homework
        self.answer = answer
        self.score = 0
        self.is_graded = False

    def show_info(self):
        """Prints the submission details."""
        print(f"  Student     : {self.student.name}")
        print(f"  Homework    : {self.homework.title}")
        print(f"  Answer      : {self.answer}")
        if self.is_graded:
            print(f"  Score       : {self.score} / {self.homework.max_score}")
        else:
            print("  Score       : Not graded yet")
        print()


class Teacher:
    """Teacher represents a teacher who creates homework and grades submissions."""

    def __init__(self, name: str, subject: str):
        """Creates a new Teacher."""
        self.name = name
        self.subject = subject

    def create_homework(self, title: str, due_date: str, description: str, max_score: int) -> Homework:
        """Creates a new Homework assignment."""
        print(f'  [OK] {self.name} created homework: "{title}"')
        return Homework(title, self.subject, due_date, description, max_score)

    def grade_submission(self, sub: Submission, score: int):
        """Assigns a score to a student's submission."""
        if score < 0 or score > sub.homework.max_score:
            print(f"  [Error] Score must be between 0 and {sub.homework.max_score}.")
            return
        sub.score = score
        sub.is_graded = True
        print(f"  [OK] {self.name} graded {sub.student.name}'s submission: "
              f"{score} / {sub.homework.max_score}")

    def show_info(self):
        """Prints teacher details."""
        print(f"  Teacher: {self.name} | Subject: {self.subject}")


class Student:
    """Student represents a student who submits homework."""

    def __init__(self, name: str, grade: str):
        """Creates a new Student."""
        self.name = name
        self.grade = grade

    def submit_homework(self, hw: Homework, answer: str) -> Submission:
        """Creates a submission for the given homework."""
        print(f'  [OK] {self.name} submitted answer for "{hw.title}"')
        return Submission(self, hw, answer)

    def show_info(self):
        """Prints student details."""
        print(f"  Student: {self.name} | Grade: {self.grade}")


def main():
    # --- Create teacher and students ---
    teacher = Teacher("Mr. Tiemoon", "Mathematics")
    student1 = Student("Tareq", "Grade 10")
    student2 = Student("Afsana", "Grade 10")

    print("=== Participants ===")
    teacher.show_info()
    student1.show_info()
    student2.show_info()
    print()

    # --- Teacher creates homework ---
    print("=== Creating Homework ===")
    hw = teacher.create_homework(
        "Quadratic Equations",
        "2026-03-15",
        "Solve 10 quadratic equations using the quadratic formula.",
        100,
    )
    print()

    print("=== Homework Details ===")
    hw.show_info()

    # --- Students submit homework ---
    print("=== Student Submissions ===")
    sub1 = student1.submit_homework(hw, "x = (-b +/- sqrt(b^2-4ac)) / 2a ... [10 solutions]")
    sub2 = student2.submit_homework(hw, "Used factoring method ... [10 solutions]")
    print()

    # --- Teacher grades submissions ---
    print("=== Grading ===")
    teacher.grade_submission(sub1, 92)
    teacher.grade_submission(sub2, 85)
    print()

    # --- Show final results ---
    print("=== Submission Results ===")
    sub1.show_info()
    sub2.show_info()


if __name__ == "__main__":
    main()
