/**
 * Practice 07: School Homework System (OOAD)
 * Task: Model a school homework system where teachers create homework,
 *       students submit answers, and teachers grade the submissions.
 *
 * How to run:
 *   go run practice-07.go
 *
 * Key Concepts:
 *   - Multiple structs collaborating (Teacher, Student, Homework, Submission)
 *   - Struct methods that create and return other structs
 *   - State tracking (graded/ungraded)
 */

package main

import "fmt"

/** Homework represents an assignment created by a teacher. */
type Homework struct {
	Title       string
	Subject     string
	DueDate     string
	Description string
	MaxScore    int
}

/** ShowInfo prints the homework details. */
func (h *Homework) ShowInfo() {
	fmt.Printf("  Homework    : %s\n", h.Title)
	fmt.Printf("  Subject     : %s\n", h.Subject)
	fmt.Printf("  Due Date    : %s\n", h.DueDate)
	fmt.Printf("  Description : %s\n", h.Description)
	fmt.Printf("  Max Score   : %d\n", h.MaxScore)
	fmt.Println()
}

/** Submission represents a student's submitted answer for a homework. */
type Submission struct {
	Student  *Student
	Homework *Homework
	Answer   string
	Score    int
	IsGraded bool
}

/** ShowInfo prints the submission details. */
func (s *Submission) ShowInfo() {
	fmt.Printf("  Student     : %s\n", s.Student.Name)
	fmt.Printf("  Homework    : %s\n", s.Homework.Title)
	fmt.Printf("  Answer      : %s\n", s.Answer)
	if s.IsGraded {
		fmt.Printf("  Score       : %d / %d\n", s.Score, s.Homework.MaxScore)
	} else {
		fmt.Printf("  Score       : Not graded yet\n")
	}
	fmt.Println()
}

/** Teacher represents a teacher who creates homework and grades submissions. */
type Teacher struct {
	Name    string
	Subject string
}

/** NewTeacher creates a new Teacher. */
func NewTeacher(name, subject string) Teacher {
	return Teacher{Name: name, Subject: subject}
}

/** CreateHomework creates a new Homework assignment. */
func (t *Teacher) CreateHomework(title, dueDate, description string, maxScore int) Homework {
	fmt.Printf("  [OK] %s created homework: \"%s\"\n", t.Name, title)
	return Homework{
		Title:       title,
		Subject:     t.Subject,
		DueDate:     dueDate,
		Description: description,
		MaxScore:    maxScore,
	}
}

/** GradeSubmission assigns a score to a student's submission. */
func (t *Teacher) GradeSubmission(sub *Submission, score int) {
	if score < 0 || score > sub.Homework.MaxScore {
		fmt.Printf("  [Error] Score must be between 0 and %d.\n", sub.Homework.MaxScore)
		return
	}
	sub.Score = score
	sub.IsGraded = true
	fmt.Printf("  [OK] %s graded %s's submission: %d / %d\n",
		t.Name, sub.Student.Name, score, sub.Homework.MaxScore)
}

/** ShowInfo prints teacher details. */
func (t *Teacher) ShowInfo() {
	fmt.Printf("  Teacher: %s | Subject: %s\n", t.Name, t.Subject)
}

/** Student represents a student who submits homework. */
type Student struct {
	Name  string
	Grade string
}

/** NewStudent creates a new Student. */
func NewStudent(name, grade string) Student {
	return Student{Name: name, Grade: grade}
}

/** SubmitHomework creates a submission for the given homework. */
func (st *Student) SubmitHomework(hw *Homework, answer string) Submission {
	fmt.Printf("  [OK] %s submitted answer for \"%s\"\n", st.Name, hw.Title)
	return Submission{
		Student:  st,
		Homework: hw,
		Answer:   answer,
		Score:    0,
		IsGraded: false,
	}
}

/** ShowInfo prints student details. */
func (st *Student) ShowInfo() {
	fmt.Printf("  Student: %s | Grade: %s\n", st.Name, st.Grade)
}

func main() {
	// --- Create teacher and students ---
	teacher := NewTeacher("Mr. Tiemoon", "Mathematics")
	student1 := NewStudent("Tareq", "Grade 10")
	student2 := NewStudent("Afsana", "Grade 10")

	fmt.Println("=== Participants ===")
	teacher.ShowInfo()
	student1.ShowInfo()
	student2.ShowInfo()
	fmt.Println()

	// --- Teacher creates homework ---
	fmt.Println("=== Creating Homework ===")
	hw := teacher.CreateHomework(
		"Quadratic Equations",
		"2026-03-15",
		"Solve 10 quadratic equations using the quadratic formula.",
		100,
	)
	fmt.Println()

	fmt.Println("=== Homework Details ===")
	hw.ShowInfo()

	// --- Students submit homework ---
	fmt.Println("=== Student Submissions ===")
	sub1 := student1.SubmitHomework(&hw, "x = (-b ± sqrt(b²-4ac)) / 2a ... [10 solutions]")
	sub2 := student2.SubmitHomework(&hw, "Used factoring method ... [10 solutions]")
	fmt.Println()

	// --- Teacher grades submissions ---
	fmt.Println("=== Grading ===")
	teacher.GradeSubmission(&sub1, 92)
	teacher.GradeSubmission(&sub2, 85)
	fmt.Println()

	// --- Show final results ---
	fmt.Println("=== Submission Results ===")
	sub1.ShowInfo()
	sub2.ShowInfo()
}
