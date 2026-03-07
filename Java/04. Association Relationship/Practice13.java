/**
 * Practice 13: Many-Many Association — Patient & Medications
 * Task: Patients are prescribed multiple medications, and each medication
 *       can be prescribed to many patients. Certain combinations are forbidden.
 *
 *       Forbidden combinations:
 *         - Antibiotics & Statins
 *         - Muscle Relaxants & CNS Depressants
 *         - Anti-Inflammatories & Anticoagulants
 *
 * How to compile and run:
 *   javac Practice13.java
 *   java Practice13
 *
 * Key Concepts:
 *   - Many-Many Association via a Prescription junction class
 *   - Conflict detection logic
 *   - Collection management
 */

import java.time.LocalDate;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Practice13 {

    /** Medication with name and category. */
    static class Medication {
        private String name;
        private String category;

        public Medication(String name, String category) {
            this.name = name;
            this.category = category;
        }

        public String getName() { return name; }
        public String getCategory() { return category; }

        public void showInfo() {
            System.out.printf("  %-25s | Category: %s%n", name, category);
        }
    }

    /** Patient who can be prescribed multiple medications. */
    static class Patient {
        private String patientId;
        private String name;

        public Patient(String patientId, String name) {
            this.patientId = patientId;
            this.name = name;
        }

        public String getPatientId() { return patientId; }
        public String getName() { return name; }

        public void showInfo() {
            System.out.println("  Patient ID : " + patientId);
            System.out.println("  Name       : " + name);
        }
    }

    /** Represents a single prescription linking a patient to a medication. */
    static class Prescription {
        private Patient patient;
        private Medication medication;
        private LocalDate prescribedDate;

        public Prescription(Patient patient, Medication medication, LocalDate prescribedDate) {
            this.patient = patient;
            this.medication = medication;
            this.prescribedDate = prescribedDate;
        }

        public Patient getPatient() { return patient; }
        public Medication getMedication() { return medication; }

        public void showInfo() {
            System.out.printf("  %-12s | %-25s | Category: %-22s | Date: %s%n",
                    patient.getName(), medication.getName(), medication.getCategory(), prescribedDate);
        }
    }

    /** Forbidden medication category combinations. */
    static final String[][] FORBIDDEN_COMBINATIONS = {
        {"Antibiotics", "Statins"},
        {"Muscle Relaxants", "CNS Depressants"},
        {"Anti-Inflammatories", "Anticoagulants"},
    };

    /** Manages all prescriptions with conflict checking. */
    static class PrescriptionManager {
        private List<Prescription> prescriptions = new ArrayList<>();

        private Set<String> getPatientCategories(Patient patient) {
            Set<String> categories = new HashSet<>();
            for (Prescription p : prescriptions) {
                if (p.getPatient().getPatientId().equals(patient.getPatientId())) {
                    categories.add(p.getMedication().getCategory());
                }
            }
            return categories;
        }

        private String checkConflict(Set<String> existingCategories, String newCategory) {
            for (String[] pair : FORBIDDEN_COMBINATIONS) {
                if (newCategory.equals(pair[0]) && existingCategories.contains(pair[1])) {
                    return pair[0] + " cannot be combined with " + pair[1];
                }
                if (newCategory.equals(pair[1]) && existingCategories.contains(pair[0])) {
                    return pair[1] + " cannot be combined with " + pair[0];
                }
            }
            return "";
        }

        public void prescribe(Patient patient, Medication medication, LocalDate date) {
            // Check for duplicate
            for (Prescription p : prescriptions) {
                if (p.getPatient().getPatientId().equals(patient.getPatientId())
                        && p.getMedication().getName().equals(medication.getName())) {
                    System.out.println("  [Error] " + medication.getName() + " is already prescribed to " + patient.getName() + ".");
                    return;
                }
            }
            // Check for conflicts
            Set<String> existingCategories = getPatientCategories(patient);
            String conflict = checkConflict(existingCategories, medication.getCategory());
            if (!conflict.isEmpty()) {
                System.out.println("  [Error] Cannot prescribe " + medication.getName()
                        + " to " + patient.getName() + ": " + conflict + ".");
                return;
            }
            prescriptions.add(new Prescription(patient, medication, date));
            System.out.println("  [OK] " + medication.getName() + " prescribed to " + patient.getName() + ".");
        }

        public List<Prescription> getPatientPrescriptions(Patient patient) {
            List<Prescription> result = new ArrayList<>();
            for (Prescription p : prescriptions) {
                if (p.getPatient().getPatientId().equals(patient.getPatientId())) {
                    result.add(p);
                }
            }
            return result;
        }

        public List<Prescription> getMedicationPatients(Medication medication) {
            List<Prescription> result = new ArrayList<>();
            for (Prescription p : prescriptions) {
                if (p.getMedication().getName().equals(medication.getName())) {
                    result.add(p);
                }
            }
            return result;
        }

        public void showAll() {
            if (prescriptions.isEmpty()) {
                System.out.println("  No prescriptions recorded.");
                return;
            }
            System.out.printf("  %-12s | %-25s | %-22s | Date%n", "Patient", "Medication", "Category");
            System.out.println("  " + "-".repeat(80));
            for (Prescription p : prescriptions) {
                p.showInfo();
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // --- Create medications ---
        Medication amoxicillin = new Medication("Amoxicillin", "Antibiotics");
        Medication ibuprofen = new Medication("Ibuprofen", "Anti-Inflammatories");
        Medication diazepam = new Medication("Diazepam", "Muscle Relaxants");
        Medication coughSyrup = new Medication("Dextromethorphan", "Cough Syrup");
        Medication insulin = new Medication("Insulin", "Insulin");
        Medication warfarin = new Medication("Warfarin", "Anticoagulants");
        Medication phenobarbital = new Medication("Phenobarbital", "CNS Depressants");
        Medication atorvastatin = new Medication("Atorvastatin", "Statins");

        System.out.println("=== Available Medications ===");
        Medication[] meds = {amoxicillin, ibuprofen, diazepam, coughSyrup,
                insulin, warfarin, phenobarbital, atorvastatin};
        for (Medication m : meds) {
            m.showInfo();
        }
        System.out.println();

        // --- Create patients ---
        Patient patient1 = new Patient("PAT-001", "Tareq");
        Patient patient2 = new Patient("PAT-002", "Afsana");
        Patient patient3 = new Patient("PAT-003", "Imtiaz");

        PrescriptionManager manager = new PrescriptionManager();
        LocalDate today = LocalDate.now();

        System.out.println("=== Prescribing Medications ===");
        manager.prescribe(patient1, amoxicillin, today);
        manager.prescribe(patient1, coughSyrup, today);
        manager.prescribe(patient1, atorvastatin, today); // CONFLICT
        System.out.println();

        manager.prescribe(patient2, diazepam, today);
        manager.prescribe(patient2, insulin, today);
        manager.prescribe(patient2, phenobarbital, today); // CONFLICT
        System.out.println();

        manager.prescribe(patient3, ibuprofen, today);
        manager.prescribe(patient3, coughSyrup, today);
        manager.prescribe(patient3, warfarin, today); // CONFLICT
        System.out.println();

        System.out.println("=== All Prescriptions ===");
        manager.showAll();

        Patient[] patients = {patient1, patient2, patient3};
        for (Patient p : patients) {
            System.out.println("=== Prescriptions for " + p.getName() + " ===");
            for (Prescription pr : manager.getPatientPrescriptions(p)) {
                pr.showInfo();
            }
            System.out.println();
        }

        System.out.println("=== Patients on Dextromethorphan (Cough Syrup) ===");
        for (Prescription p : manager.getMedicationPatients(coughSyrup)) {
            System.out.println("  " + p.getPatient().getName() + " (prescribed: " + p.prescribedDate + ")");
        }
        System.out.println();
    }
}
