/**
 * Practice 12: 1-Many Association — Doctor & Patients
 * Task: Manage doctors and patients. Doctors: schedule appointments,
 *       diagnose, prescribe treatments, discharge patients.
 *       Patients: view appointments, check treatments, track history.
 *
 * How to compile and run:
 *   javac Practice12.java
 *   java Practice12
 *
 * Key Concepts:
 *   - 1-Many Association (Doctor has many Patients)
 *   - Object references and collection management
 *   - Business logic across associated objects
 */

import java.time.LocalDate;
import java.util.ArrayList;
import java.util.List;

public class Practice12 {

    /** Patient with personal info, medical history, appointments, and treatments. */
    static class Patient {
        private String patientId;
        private String name;
        private List<String> medicalHistory;
        private List<String> appointments; // "date: description"
        private List<String> treatments;
        private boolean isAdmitted;

        public Patient(String patientId, String name) {
            this.patientId = patientId;
            this.name = name;
            this.medicalHistory = new ArrayList<>();
            this.appointments = new ArrayList<>();
            this.treatments = new ArrayList<>();
            this.isAdmitted = true;
        }

        public String getName() { return name; }
        public String getPatientId() { return patientId; }

        public void viewAppointments() {
            if (appointments.isEmpty()) {
                System.out.println("  " + name + " has no appointments.");
                return;
            }
            System.out.println("  Appointments for " + name + ":");
            for (String appt : appointments) {
                System.out.println("    - " + appt);
            }
        }

        public boolean hasOngoingTreatments() {
            return !treatments.isEmpty();
        }

        public void viewMedicalHistory() {
            if (medicalHistory.isEmpty()) {
                System.out.println("  " + name + " has no medical history records.");
                return;
            }
            System.out.println("  Medical history for " + name + ":");
            for (String record : medicalHistory) {
                System.out.println("    - " + record);
            }
        }

        public void showInfo() {
            String status = isAdmitted ? "Admitted" : "Discharged";
            System.out.println("  Patient ID    : " + patientId);
            System.out.println("  Name          : " + name);
            System.out.println("  Status        : " + status);
            System.out.println("  Diagnoses     : " + medicalHistory.size());
            System.out.println("  Treatments    : " + treatments.size());
            System.out.println("  Appointments  : " + appointments.size());
            System.out.println();
        }
    }

    /** Doctor who manages multiple patients (1-Many association). */
    static class Doctor {
        private String doctorId;
        private String name;
        private String specialization;
        private List<Patient> patients;

        public Doctor(String doctorId, String name, String specialization) {
            this.doctorId = doctorId;
            this.name = name;
            this.specialization = specialization;
            this.patients = new ArrayList<>();
        }

        public void addPatient(Patient patient) {
            patients.add(patient);
            System.out.println("  [OK] " + patient.getName() + " assigned to Dr. " + name + ".");
        }

        public void scheduleAppointment(Patient patient, LocalDate date, String description) {
            if (!patients.contains(patient)) {
                System.out.println("  [Error] " + patient.getName() + " is not under Dr. " + name + "'s care.");
                return;
            }
            patient.appointments.add(date + ": " + description);
            System.out.println("  [OK] Appointment scheduled for " + patient.getName() + " on " + date + ": " + description);
        }

        public void diagnose(Patient patient, String diagnosis) {
            if (!patients.contains(patient)) {
                System.out.println("  [Error] " + patient.getName() + " is not under Dr. " + name + "'s care.");
                return;
            }
            patient.medicalHistory.add(diagnosis);
            System.out.println("  [OK] Dr. " + name + " diagnosed " + patient.getName() + ": " + diagnosis);
        }

        public void prescribeTreatment(Patient patient, String treatment) {
            if (!patients.contains(patient)) {
                System.out.println("  [Error] " + patient.getName() + " is not under Dr. " + name + "'s care.");
                return;
            }
            if (patient.medicalHistory.isEmpty()) {
                System.out.println("  [Error] " + patient.getName() + " has no diagnosis. Diagnose before prescribing.");
                return;
            }
            patient.treatments.add(treatment);
            System.out.println("  [OK] Dr. " + name + " prescribed '" + treatment + "' to " + patient.getName() + ".");
        }

        public void dischargePatient(Patient patient) {
            if (!patients.contains(patient)) {
                System.out.println("  [Error] " + patient.getName() + " is not under Dr. " + name + "'s care.");
                return;
            }
            patient.isAdmitted = false;
            patient.treatments.clear();
            patients.remove(patient);
            System.out.println("  [OK] " + patient.getName() + " has been discharged by Dr. " + name + ".");
        }

        public void showInfo() {
            System.out.println("  Doctor ID     : " + doctorId);
            System.out.println("  Name          : Dr. " + name);
            System.out.println("  Specialization: " + specialization);
            System.out.println("  Patients      : " + patients.size());
            for (Patient p : patients) {
                String status = p.isAdmitted ? "Admitted" : "Discharged";
                System.out.println("    - " + p.getName() + " (" + p.getPatientId() + ") [" + status + "]");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        Doctor doctor = new Doctor("DOC-001", "Karim", "General Medicine");
        Patient patient1 = new Patient("PAT-001", "Tareq");
        Patient patient2 = new Patient("PAT-002", "Afsana");
        Patient patient3 = new Patient("PAT-003", "Imtiaz");

        System.out.println("=== Assigning Patients ===");
        doctor.addPatient(patient1);
        doctor.addPatient(patient2);
        doctor.addPatient(patient3);
        System.out.println();

        System.out.println("=== Doctor Info ===");
        doctor.showInfo();

        System.out.println("=== Scheduling Appointments ===");
        doctor.scheduleAppointment(patient1, LocalDate.of(2026, 3, 10), "General Checkup");
        doctor.scheduleAppointment(patient2, LocalDate.of(2026, 3, 12), "Follow-up Visit");
        doctor.scheduleAppointment(patient3, LocalDate.of(2026, 3, 15), "Blood Test Review");
        System.out.println();

        System.out.println("=== Diagnosing Patients ===");
        doctor.diagnose(patient1, "Mild Fever");
        doctor.diagnose(patient2, "Vitamin D Deficiency");
        doctor.diagnose(patient3, "Fractured Wrist");
        System.out.println();

        System.out.println("=== Prescribing Treatments ===");
        doctor.prescribeTreatment(patient1, "Paracetamol 500mg — 3 times daily");
        doctor.prescribeTreatment(patient2, "Vitamin D supplements — once daily");
        doctor.prescribeTreatment(patient3, "Cast and rest for 6 weeks");
        System.out.println();

        System.out.println("=== Patient Details ===");
        patient1.showInfo();
        patient1.viewAppointments();
        patient1.viewMedicalHistory();
        System.out.println();

        patient2.showInfo();
        patient2.viewAppointments();
        patient2.viewMedicalHistory();
        System.out.println();

        System.out.println("=== Discharging Patient ===");
        doctor.dischargePatient(patient1);
        System.out.println();

        System.out.println("=== After Discharge ===");
        patient1.showInfo();
        doctor.showInfo();
    }
}
