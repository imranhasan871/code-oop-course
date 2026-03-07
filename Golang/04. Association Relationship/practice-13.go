/**
 * Practice 13: Many-Many Association — Patient & Medications
 * Task: Patients prescribed multiple medications; each medication
 *       prescribed to many patients. Certain combinations forbidden.
 *
 *       Forbidden combinations:
 *         - Antibiotics & Statins
 *         - Muscle Relaxants & CNS Depressants
 *         - Anti-Inflammatories & Anticoagulants
 *
 * How to run:
 *   go run practice-13.go
 *
 * Key Concepts:
 *   - Many-Many Association via a Prescription junction struct
 *   - Conflict detection logic
 *   - Slice and map management
 */

package main

import (
	"fmt"
	"time"
)

/** Medication with name and category. */
type Medication13 struct {
	Name     string
	Category string
}

func (m *Medication13) ShowInfo() {
	fmt.Printf("  %-25s | Category: %s\n", m.Name, m.Category)
}

/** Patient who can be prescribed multiple medications. */
type Patient13 struct {
	PatientId string
	Name      string
}

/** Prescription links a patient to a medication. */
type Prescription13 struct {
	Patient        *Patient13
	Medication     *Medication13
	PrescribedDate time.Time
}

func (p *Prescription13) ShowInfo() {
	fmt.Printf("  %-12s | %-25s | Category: %-22s | Date: %s\n",
		p.Patient.Name, p.Medication.Name, p.Medication.Category,
		p.PrescribedDate.Format("2006-01-02"))
}

/** Forbidden medication category combinations. */
var forbiddenCombinations = [][2]string{
	{"Antibiotics", "Statins"},
	{"Muscle Relaxants", "CNS Depressants"},
	{"Anti-Inflammatories", "Anticoagulants"},
}

/** PrescriptionManager manages all prescriptions with conflict checking. */
type PrescriptionManager13 struct {
	Prescriptions []*Prescription13
}

func NewPrescriptionManager13() *PrescriptionManager13 {
	return &PrescriptionManager13{}
}

func (pm *PrescriptionManager13) getPatientCategories(patient *Patient13) map[string]bool {
	categories := make(map[string]bool)
	for _, p := range pm.Prescriptions {
		if p.Patient.PatientId == patient.PatientId {
			categories[p.Medication.Category] = true
		}
	}
	return categories
}

func (pm *PrescriptionManager13) checkConflict(existing map[string]bool, newCategory string) string {
	for _, pair := range forbiddenCombinations {
		if newCategory == pair[0] && existing[pair[1]] {
			return fmt.Sprintf("%s cannot be combined with %s", pair[0], pair[1])
		}
		if newCategory == pair[1] && existing[pair[0]] {
			return fmt.Sprintf("%s cannot be combined with %s", pair[1], pair[0])
		}
	}
	return ""
}

func (pm *PrescriptionManager13) Prescribe(patient *Patient13, medication *Medication13, date time.Time) {
	// Check duplicate
	for _, p := range pm.Prescriptions {
		if p.Patient.PatientId == patient.PatientId && p.Medication.Name == medication.Name {
			fmt.Printf("  [Error] %s is already prescribed to %s.\n", medication.Name, patient.Name)
			return
		}
	}
	// Check conflict
	existing := pm.getPatientCategories(patient)
	conflict := pm.checkConflict(existing, medication.Category)
	if conflict != "" {
		fmt.Printf("  [Error] Cannot prescribe %s to %s: %s.\n", medication.Name, patient.Name, conflict)
		return
	}
	pm.Prescriptions = append(pm.Prescriptions, &Prescription13{patient, medication, date})
	fmt.Printf("  [OK] %s prescribed to %s.\n", medication.Name, patient.Name)
}

func (pm *PrescriptionManager13) GetPatientPrescriptions(patient *Patient13) []*Prescription13 {
	var result []*Prescription13
	for _, p := range pm.Prescriptions {
		if p.Patient.PatientId == patient.PatientId {
			result = append(result, p)
		}
	}
	return result
}

func (pm *PrescriptionManager13) GetMedicationPatients(medication *Medication13) []*Prescription13 {
	var result []*Prescription13
	for _, p := range pm.Prescriptions {
		if p.Medication.Name == medication.Name {
			result = append(result, p)
		}
	}
	return result
}

func (pm *PrescriptionManager13) ShowAll() {
	if len(pm.Prescriptions) == 0 {
		fmt.Println("  No prescriptions recorded.")
		return
	}
	fmt.Printf("  %-12s | %-25s | %-22s | Date\n", "Patient", "Medication", "Category")
	fmt.Println("  --------------------------------------------------------------------------------")
	for _, p := range pm.Prescriptions {
		p.ShowInfo()
	}
	fmt.Println()
}

func main() {
	amoxicillin := &Medication13{"Amoxicillin", "Antibiotics"}
	ibuprofen := &Medication13{"Ibuprofen", "Anti-Inflammatories"}
	diazepam := &Medication13{"Diazepam", "Muscle Relaxants"}
	coughSyrup := &Medication13{"Dextromethorphan", "Cough Syrup"}
	insulin := &Medication13{"Insulin", "Insulin"}
	warfarin := &Medication13{"Warfarin", "Anticoagulants"}
	phenobarbital := &Medication13{"Phenobarbital", "CNS Depressants"}
	atorvastatin := &Medication13{"Atorvastatin", "Statins"}

	fmt.Println("=== Available Medications ===")
	for _, m := range []*Medication13{amoxicillin, ibuprofen, diazepam, coughSyrup,
		insulin, warfarin, phenobarbital, atorvastatin} {
		m.ShowInfo()
	}
	fmt.Println()

	patient1 := &Patient13{"PAT-001", "Tareq"}
	patient2 := &Patient13{"PAT-002", "Afsana"}
	patient3 := &Patient13{"PAT-003", "Imtiaz"}

	manager := NewPrescriptionManager13()
	today := time.Now()

	fmt.Println("=== Prescribing Medications ===")
	manager.Prescribe(patient1, amoxicillin, today)
	manager.Prescribe(patient1, coughSyrup, today)
	manager.Prescribe(patient1, atorvastatin, today) // CONFLICT
	fmt.Println()

	manager.Prescribe(patient2, diazepam, today)
	manager.Prescribe(patient2, insulin, today)
	manager.Prescribe(patient2, phenobarbital, today) // CONFLICT
	fmt.Println()

	manager.Prescribe(patient3, ibuprofen, today)
	manager.Prescribe(patient3, coughSyrup, today)
	manager.Prescribe(patient3, warfarin, today) // CONFLICT
	fmt.Println()

	fmt.Println("=== All Prescriptions ===")
	manager.ShowAll()

	for _, patient := range []*Patient13{patient1, patient2, patient3} {
		fmt.Printf("=== Prescriptions for %s ===\n", patient.Name)
		for _, p := range manager.GetPatientPrescriptions(patient) {
			p.ShowInfo()
		}
		fmt.Println()
	}

	fmt.Println("=== Patients on Dextromethorphan (Cough Syrup) ===")
	for _, p := range manager.GetMedicationPatients(coughSyrup) {
		fmt.Printf("  %s (prescribed: %s)\n", p.Patient.Name, p.PrescribedDate.Format("2006-01-02"))
	}
	fmt.Println()
}
