/**
 * Practice 07: School Homework System (OOAD)
 * Task: Model a school homework system where teachers create homework,
 *       students submit answers, and teachers grade the submissions.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice07.cs
 *
 * Or create a Console App and copy contents into Program.cs:
 *   dotnet new console -n Practice07
 *   dotnet run --project Practice07
 *
 * Key Concepts:
 *   - Multiple classes collaborating (Teacher, Student, Homework, Submission)
 *   - Methods that create and return objects
 *   - State tracking (graded/ungraded)
 */

using System;

class Practice07
{
    /** Homework represents an assignment created by a teacher. */
    class Homework
    {
        public string Title;
        public string Subject;
        public string DueDate;
        public string Description;
        public int MaxScore;

        /** Creates a new Homework assignment. */
        public Homework(string title, string subject, string dueDate, string description, int maxScore)
        {
            Title = title;
            Subject = subject;
            DueDate = dueDate;
            Description = description;
            MaxScore = maxScore;
        }

        /** Prints the homework details. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Homework    : {Title}");
            Console.WriteLine($"  Subject     : {Subject}");
            Console.WriteLine($"  Due Date    : {DueDate}");
            Console.WriteLine($"  Description : {Description}");
            Console.WriteLine($"  Max Score   : {MaxScore}");
            Console.WriteLine();
        }
    }

    /** Submission represents a student's submitted answer for a homework. */
    class Submission
    {
        public Student StudentRef;
        public Homework HomeworkRef;
        public string Answer;
        public int Score;
        public bool IsGraded;

        /** Creates a new Submission. */
        public Submission(Student student, Homework homework, string answer)
        {
            StudentRef = student;
            HomeworkRef = homework;
            Answer = answer;
            Score = 0;
            IsGraded = false;
        }

        /** Prints the submission details. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Student     : {StudentRef.Name}");
            Console.WriteLine($"  Homework    : {HomeworkRef.Title}");
            Console.WriteLine($"  Answer      : {Answer}");
            if (IsGraded)
                Console.WriteLine($"  Score       : {Score} / {HomeworkRef.MaxScore}");
            else
                Console.WriteLine("  Score       : Not graded yet");
            Console.WriteLine();
        }
    }

    /** Teacher represents a teacher who creates homework and grades submissions. */
    class Teacher
    {
        public string Name;
        public string Subject;

        /** Creates a new Teacher. */
        public Teacher(string name, string subject)
        {
            Name = name;
            Subject = subject;
        }

        /** Creates a new Homework assignment. */
        public Homework CreateHomework(string title, string dueDate, string description, int maxScore)
        {
            Console.WriteLine($"  [OK] {Name} created homework: \"{title}\"");
            return new Homework(title, Subject, dueDate, description, maxScore);
        }

        /** Assigns a score to a student's submission. */
        public void GradeSubmission(Submission sub, int score)
        {
            if (score < 0 || score > sub.HomeworkRef.MaxScore)
            {
                Console.WriteLine($"  [Error] Score must be between 0 and {sub.HomeworkRef.MaxScore}.");
                return;
            }
            sub.Score = score;
            sub.IsGraded = true;
            Console.WriteLine($"  [OK] {Name} graded {sub.StudentRef.Name}'s submission: {score} / {sub.HomeworkRef.MaxScore}");
        }

        /** Prints teacher details. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Teacher: {Name} | Subject: {Subject}");
        }
    }

    /** Student represents a student who submits homework. */
    class Student
    {
        public string Name;
        public string Grade;

        /** Creates a new Student. */
        public Student(string name, string grade)
        {
            Name = name;
            Grade = grade;
        }

        /** Creates a submission for the given homework. */
        public Submission SubmitHomework(Homework hw, string answer)
        {
            Console.WriteLine($"  [OK] {Name} submitted answer for \"{hw.Title}\"");
            return new Submission(this, hw, answer);
        }

        /** Prints student details. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Student: {Name} | Grade: {Grade}");
        }
    }

    static void Main(string[] args)
    {
        // --- Create teacher and students ---
        Teacher teacher = new Teacher("Mr. Tiemoon", "Mathematics");
        Student student1 = new Student("Tareq", "Grade 10");
        Student student2 = new Student("Afsana", "Grade 10");

        Console.WriteLine("=== Participants ===");
        teacher.ShowInfo();
        student1.ShowInfo();
        student2.ShowInfo();
        Console.WriteLine();

        // --- Teacher creates homework ---
        Console.WriteLine("=== Creating Homework ===");
        Homework hw = teacher.CreateHomework(
            "Quadratic Equations",
            "2026-03-15",
            "Solve 10 quadratic equations using the quadratic formula.",
            100
        );
        Console.WriteLine();

        Console.WriteLine("=== Homework Details ===");
        hw.ShowInfo();

        // --- Students submit homework ---
        Console.WriteLine("=== Student Submissions ===");
        Submission sub1 = student1.SubmitHomework(hw, "x = (-b +/- sqrt(b^2-4ac)) / 2a ... [10 solutions]");
        Submission sub2 = student2.SubmitHomework(hw, "Used factoring method ... [10 solutions]");
        Console.WriteLine();

        // --- Teacher grades submissions ---
        Console.WriteLine("=== Grading ===");
        teacher.GradeSubmission(sub1, 92);
        teacher.GradeSubmission(sub2, 85);
        Console.WriteLine();

        // --- Show final results ---
        Console.WriteLine("=== Submission Results ===");
        sub1.ShowInfo();
        sub2.ShowInfo();
    }
}
