/**
 * Practice 12: 1-Many Association — Doctor & Patients
 * Task: Manage doctors and patients. Doctors: schedule appointments,
 *       diagnose, prescribe treatments, discharge patients.
 *       Patients: view appointments, check treatments, track history.
 *
 * How to run:
 *   go run practice-12.go
 *
 * Key Concepts:
 *   - 1-Many Association (Doctor has many Patients)
 *   - Pointer receivers and slice management
 *   - Business logic across associated structs
 */

package main

import (
	"fmt"
	"time"
)

/** Patient with personal info, medical history, appointments, and treatments. */
type Patient12 struct {
	PatientId      string
	Name           string
	MedicalHistory []string
	Appointments   []string
	Treatments     []string
	IsAdmitted     bool
}

func NewPatient12(patientId, name string) *Patient12 {
	return &Patient12{
		PatientId:  patientId,
		Name:       name,
		IsAdmitted: true,
	}
}

func (p *Patient12) ViewAppointments() {
	if len(p.Appointments) == 0 {
		fmt.Printf("  %s has no appointments.\n", p.Name)
		return
	}
	fmt.Printf("  Appointments for %s:\n", p.Name)
	for _, appt := range p.Appointments {
		fmt.Println("    -", appt)
	}
}

func (p *Patient12) HasOngoingTreatments() bool {
	return len(p.Treatments) > 0
}

func (p *Patient12) ViewMedicalHistory() {
	if len(p.MedicalHistory) == 0 {
		fmt.Printf("  %s has no medical history records.\n", p.Name)
		return
	}
	fmt.Printf("  Medical history for %s:\n", p.Name)
	for _, record := range p.MedicalHistory {
		fmt.Println("    -", record)
	}
}

func (p *Patient12) ShowInfo() {
	status := "Admitted"
	if !p.IsAdmitted {
		status = "Discharged"
	}
	fmt.Println("  Patient ID    :", p.PatientId)
	fmt.Println("  Name          :", p.Name)
	fmt.Println("  Status        :", status)
	fmt.Println("  Diagnoses     :", len(p.MedicalHistory))
	fmt.Println("  Treatments    :", len(p.Treatments))
	fmt.Println("  Appointments  :", len(p.Appointments))
	fmt.Println()
}

/** Doctor who manages multiple patients (1-Many association). */
type Doctor12 struct {
	DoctorId       string
	Name           string
	Specialization string
	Patients       []*Patient12
}

func NewDoctor12(doctorId, name, specialization string) *Doctor12 {
	return &Doctor12{DoctorId: doctorId, Name: name, Specialization: specialization}
}

func (d *Doctor12) findPatient(patient *Patient12) int {
	for i, p := range d.Patients {
		if p.PatientId == patient.PatientId {
			return i
		}
	}
	return -1
}

func (d *Doctor12) AddPatient(patient *Patient12) {
	d.Patients = append(d.Patients, patient)
	fmt.Printf("  [OK] %s assigned to Dr. %s.\n", patient.Name, d.Name)
}

func (d *Doctor12) ScheduleAppointment(patient *Patient12, date time.Time, description string) {
	if d.findPatient(patient) == -1 {
		fmt.Printf("  [Error] %s is not under Dr. %s's care.\n", patient.Name, d.Name)
		return
	}
	appt := fmt.Sprintf("%s: %s", date.Format("2006-01-02"), description)
	patient.Appointments = append(patient.Appointments, appt)
	fmt.Printf("  [OK] Appointment scheduled for %s on %s: %s\n",
		patient.Name, date.Format("2006-01-02"), description)
}

func (d *Doctor12) Diagnose(patient *Patient12, diagnosis string) {
	if d.findPatient(patient) == -1 {
		fmt.Printf("  [Error] %s is not under Dr. %s's care.\n", patient.Name, d.Name)
		return
	}
	patient.MedicalHistory = append(patient.MedicalHistory, diagnosis)
	fmt.Printf("  [OK] Dr. %s diagnosed %s: %s\n", d.Name, patient.Name, diagnosis)
}

func (d *Doctor12) PrescribeTreatment(patient *Patient12, treatment string) {
	if d.findPatient(patient) == -1 {
		fmt.Printf("  [Error] %s is not under Dr. %s's care.\n", patient.Name, d.Name)
		return
	}
	if len(patient.MedicalHistory) == 0 {
		fmt.Printf("  [Error] %s has no diagnosis. Diagnose before prescribing.\n", patient.Name)
		return
	}
	patient.Treatments = append(patient.Treatments, treatment)
	fmt.Printf("  [OK] Dr. %s prescribed '%s' to %s.\n", d.Name, treatment, patient.Name)
}

func (d *Doctor12) DischargePatient(patient *Patient12) {
	idx := d.findPatient(patient)
	if idx == -1 {
		fmt.Printf("  [Error] %s is not under Dr. %s's care.\n", patient.Name, d.Name)
		return
	}
	patient.IsAdmitted = false
	patient.Treatments = nil
	d.Patients = append(d.Patients[:idx], d.Patients[idx+1:]...)
	fmt.Printf("  [OK] %s has been discharged by Dr. %s.\n", patient.Name, d.Name)
}

func (d *Doctor12) ShowInfo() {
	fmt.Println("  Doctor ID     :", d.DoctorId)
	fmt.Println("  Name          : Dr.", d.Name)
	fmt.Println("  Specialization:", d.Specialization)
	fmt.Println("  Patients      :", len(d.Patients))
	for _, p := range d.Patients {
		status := "Admitted"
		if !p.IsAdmitted {
			status = "Discharged"
		}
		fmt.Printf("    - %s (%s) [%s]\n", p.Name, p.PatientId, status)
	}
	fmt.Println()
}

func main() {
	doctor := NewDoctor12("DOC-001", "Karim", "General Medicine")
	patient1 := NewPatient12("PAT-001", "Tareq")
	patient2 := NewPatient12("PAT-002", "Afsana")
	patient3 := NewPatient12("PAT-003", "Imtiaz")

	fmt.Println("=== Assigning Patients ===")
	doctor.AddPatient(patient1)
	doctor.AddPatient(patient2)
	doctor.AddPatient(patient3)
	fmt.Println()

	fmt.Println("=== Doctor Info ===")
	doctor.ShowInfo()

	fmt.Println("=== Scheduling Appointments ===")
	doctor.ScheduleAppointment(patient1, time.Date(2026, 3, 10, 0, 0, 0, 0, time.Local), "General Checkup")
	doctor.ScheduleAppointment(patient2, time.Date(2026, 3, 12, 0, 0, 0, 0, time.Local), "Follow-up Visit")
	doctor.ScheduleAppointment(patient3, time.Date(2026, 3, 15, 0, 0, 0, 0, time.Local), "Blood Test Review")
	fmt.Println()

	fmt.Println("=== Diagnosing Patients ===")
	doctor.Diagnose(patient1, "Mild Fever")
	doctor.Diagnose(patient2, "Vitamin D Deficiency")
	doctor.Diagnose(patient3, "Fractured Wrist")
	fmt.Println()

	fmt.Println("=== Prescribing Treatments ===")
	doctor.PrescribeTreatment(patient1, "Paracetamol 500mg — 3 times daily")
	doctor.PrescribeTreatment(patient2, "Vitamin D supplements — once daily")
	doctor.PrescribeTreatment(patient3, "Cast and rest for 6 weeks")
	fmt.Println()

	fmt.Println("=== Patient Details ===")
	patient1.ShowInfo()
	patient1.ViewAppointments()
	patient1.ViewMedicalHistory()
	fmt.Println()

	patient2.ShowInfo()
	patient2.ViewAppointments()
	patient2.ViewMedicalHistory()
	fmt.Println()

	fmt.Println("=== Discharging Patient ===")
	doctor.DischargePatient(patient1)
	fmt.Println()

	fmt.Println("=== After Discharge ===")
	patient1.ShowInfo()
	doctor.ShowInfo()
}
