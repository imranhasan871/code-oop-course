/**
 * Practice 07: School Homework System (OOAD)
 * Task: Model a school homework system where teachers create homework,
 *       students submit answers, and teachers grade the submissions.
 *
 * How to compile and run:
 *   javac Practice07.java
 *   java Practice07
 *
 * Key Concepts:
 *   - Multiple classes collaborating (Teacher, Student, Homework, Submission)
 *   - Methods that create and return objects
 *   - State tracking (graded/ungraded)
 */

public class Practice07 {

    /** Homework represents an assignment created by a teacher. */
    static class Homework {
        private String title;
        private String subject;
        private String dueDate;
        private String description;
        private int maxScore;

        /** Creates a new Homework assignment. */
        public Homework(String title, String subject, String dueDate, String description, int maxScore) {
            this.title = title;
            this.subject = subject;
            this.dueDate = dueDate;
            this.description = description;
            this.maxScore = maxScore;
        }

        /** Prints the homework details. */
        public void showInfo() {
            System.out.printf("  Homework    : %s%n", title);
            System.out.printf("  Subject     : %s%n", subject);
            System.out.printf("  Due Date    : %s%n", dueDate);
            System.out.printf("  Description : %s%n", description);
            System.out.printf("  Max Score   : %d%n", maxScore);
            System.out.println();
        }
    }

    /** Submission represents a student's submitted answer for a homework. */
    static class Submission {
        private Student student;
        private Homework homework;
        private String answer;
        private int score;
        private boolean isGraded;

        /** Creates a new Submission. */
        public Submission(Student student, Homework homework, String answer) {
            this.student = student;
            this.homework = homework;
            this.answer = answer;
            this.score = 0;
            this.isGraded = false;
        }

        /** Prints the submission details. */
        public void showInfo() {
            System.out.printf("  Student     : %s%n", student.name);
            System.out.printf("  Homework    : %s%n", homework.title);
            System.out.printf("  Answer      : %s%n", answer);
            if (isGraded) {
                System.out.printf("  Score       : %d / %d%n", score, homework.maxScore);
            } else {
                System.out.printf("  Score       : Not graded yet%n");
            }
            System.out.println();
        }
    }

    /** Teacher represents a teacher who creates homework and grades submissions. */
    static class Teacher {
        private String name;
        private String subject;

        /** Creates a new Teacher. */
        public Teacher(String name, String subject) {
            this.name = name;
            this.subject = subject;
        }

        /** Creates a new Homework assignment. */
        public Homework createHomework(String title, String dueDate, String description, int maxScore) {
            System.out.printf("  [OK] %s created homework: \"%s\"%n", name, title);
            return new Homework(title, subject, dueDate, description, maxScore);
        }

        /** Assigns a score to a student's submission. */
        public void gradeSubmission(Submission sub, int score) {
            if (score < 0 || score > sub.homework.maxScore) {
                System.out.printf("  [Error] Score must be between 0 and %d.%n", sub.homework.maxScore);
                return;
            }
            sub.score = score;
            sub.isGraded = true;
            System.out.printf("  [OK] %s graded %s's submission: %d / %d%n",
                    name, sub.student.name, score, sub.homework.maxScore);
        }

        /** Prints teacher details. */
        public void showInfo() {
            System.out.printf("  Teacher: %s | Subject: %s%n", name, subject);
        }
    }

    /** Student represents a student who submits homework. */
    static class Student {
        private String name;
        private String grade;

        /** Creates a new Student. */
        public Student(String name, String grade) {
            this.name = name;
            this.grade = grade;
        }

        /** Creates a submission for the given homework. */
        public Submission submitHomework(Homework hw, String answer) {
            System.out.printf("  [OK] %s submitted answer for \"%s\"%n", name, hw.title);
            return new Submission(this, hw, answer);
        }

        /** Prints student details. */
        public void showInfo() {
            System.out.printf("  Student: %s | Grade: %s%n", name, grade);
        }
    }

    public static void main(String[] args) {
        // --- Create teacher and students ---
        Teacher teacher = new Teacher("Mr. Tiemoon", "Mathematics");
        Student student1 = new Student("Tareq", "Grade 10");
        Student student2 = new Student("Afsana", "Grade 10");

        System.out.println("=== Participants ===");
        teacher.showInfo();
        student1.showInfo();
        student2.showInfo();
        System.out.println();

        // --- Teacher creates homework ---
        System.out.println("=== Creating Homework ===");
        Homework hw = teacher.createHomework(
                "Quadratic Equations",
                "2026-03-15",
                "Solve 10 quadratic equations using the quadratic formula.",
                100
        );
        System.out.println();

        System.out.println("=== Homework Details ===");
        hw.showInfo();

        // --- Students submit homework ---
        System.out.println("=== Student Submissions ===");
        Submission sub1 = student1.submitHomework(hw, "x = (-b +/- sqrt(b^2-4ac)) / 2a ... [10 solutions]");
        Submission sub2 = student2.submitHomework(hw, "Used factoring method ... [10 solutions]");
        System.out.println();

        // --- Teacher grades submissions ---
        System.out.println("=== Grading ===");
        teacher.gradeSubmission(sub1, 92);
        teacher.gradeSubmission(sub2, 85);
        System.out.println();

        // --- Show final results ---
        System.out.println("=== Submission Results ===");
        sub1.showInfo();
        sub2.showInfo();
    }
}
