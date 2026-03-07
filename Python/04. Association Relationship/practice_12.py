"""
Practice 12: 1-Many Association — Doctor & Patients
Task: Manage doctors and the patients under their care.
      Doctors: schedule appointments, diagnose, prescribe treatments, discharge.
      Patients: view appointments, check treatments, track medical history.

How to run:
  python practice_12.py

Key Concepts:
  - 1-Many Association (Doctor has many Patients)
  - Object references and collection management
  - Business logic across associated objects
"""

from datetime import date


class Patient:
    """Patient with personal info, medical history, appointments, and treatments."""

    def __init__(self, patient_id: str, name: str):
        """Creates a new Patient."""
        self.patient_id = patient_id
        self.name = name
        self.medical_history = []     # list of diagnosis strings
        self.appointments = []        # list of (date, description) tuples
        self.treatments = []          # list of treatment strings
        self.is_admitted = True

    def view_appointments(self):
        """Prints upcoming appointments."""
        upcoming = [(d, desc) for d, desc in self.appointments if d >= date.today()]
        if not upcoming:
            print(f"  {self.name} has no upcoming appointments.")
            return
        print(f"  Upcoming appointments for {self.name}:")
        for appt_date, desc in upcoming:
            print(f"    - {appt_date}: {desc}")

    def has_ongoing_treatments(self) -> bool:
        """Returns True if the patient has any ongoing treatments."""
        return len(self.treatments) > 0

    def view_medical_history(self):
        """Prints the patient's medical history."""
        if not self.medical_history:
            print(f"  {self.name} has no medical history records.")
            return
        print(f"  Medical history for {self.name}:")
        for record in self.medical_history:
            print(f"    - {record}")

    def show_info(self):
        """Prints patient details."""
        status = "Admitted" if self.is_admitted else "Discharged"
        print(f"  Patient ID    : {self.patient_id}")
        print(f"  Name          : {self.name}")
        print(f"  Status        : {status}")
        print(f"  Diagnoses     : {len(self.medical_history)}")
        print(f"  Treatments    : {len(self.treatments)}")
        print(f"  Appointments  : {len(self.appointments)}")
        print()


class Doctor:
    """Doctor who manages multiple patients (1-Many association)."""

    def __init__(self, doctor_id: str, name: str, specialization: str):
        """Creates a new Doctor."""
        self.doctor_id = doctor_id
        self.name = name
        self.specialization = specialization
        self.patients = []  # list of Patient objects

    def add_patient(self, patient: Patient):
        """Adds a patient under this doctor's care."""
        self.patients.append(patient)
        print(f"  [OK] {patient.name} assigned to Dr. {self.name}.")

    def schedule_appointment(self, patient: Patient, appt_date: date, description: str):
        """Schedules an appointment for a patient."""
        if patient not in self.patients:
            print(f"  [Error] {patient.name} is not under Dr. {self.name}'s care.")
            return
        patient.appointments.append((appt_date, description))
        print(f"  [OK] Appointment scheduled for {patient.name} on {appt_date}: {description}")

    def diagnose(self, patient: Patient, diagnosis: str):
        """Adds a diagnosis to a patient's medical history."""
        if patient not in self.patients:
            print(f"  [Error] {patient.name} is not under Dr. {self.name}'s care.")
            return
        patient.medical_history.append(diagnosis)
        print(f"  [OK] Dr. {self.name} diagnosed {patient.name}: {diagnosis}")

    def prescribe_treatment(self, patient: Patient, treatment: str):
        """Prescribes a treatment to a patient based on diagnosis."""
        if patient not in self.patients:
            print(f"  [Error] {patient.name} is not under Dr. {self.name}'s care.")
            return
        if not patient.medical_history:
            print(f"  [Error] {patient.name} has no diagnosis. Diagnose before prescribing.")
            return
        patient.treatments.append(treatment)
        print(f"  [OK] Dr. {self.name} prescribed '{treatment}' to {patient.name}.")

    def discharge_patient(self, patient: Patient):
        """Discharges a patient once they have recovered."""
        if patient not in self.patients:
            print(f"  [Error] {patient.name} is not under Dr. {self.name}'s care.")
            return
        patient.is_admitted = False
        patient.treatments.clear()
        self.patients.remove(patient)
        print(f"  [OK] {patient.name} has been discharged by Dr. {self.name}.")

    def show_info(self):
        """Prints doctor details and patient list."""
        print(f"  Doctor ID     : {self.doctor_id}")
        print(f"  Name          : Dr. {self.name}")
        print(f"  Specialization: {self.specialization}")
        print(f"  Patients      : {len(self.patients)}")
        for p in self.patients:
            status = "Admitted" if p.is_admitted else "Discharged"
            print(f"    - {p.name} ({p.patient_id}) [{status}]")
        print()


def main():
    # --- Create doctor and patients ---
    doctor = Doctor("DOC-001", "Karim", "General Medicine")
    patient1 = Patient("PAT-001", "Tareq")
    patient2 = Patient("PAT-002", "Afsana")
    patient3 = Patient("PAT-003", "Imtiaz")

    print("=== Assigning Patients ===")
    doctor.add_patient(patient1)
    doctor.add_patient(patient2)
    doctor.add_patient(patient3)
    print()

    print("=== Doctor Info ===")
    doctor.show_info()

    # --- Schedule appointments ---
    print("=== Scheduling Appointments ===")
    doctor.schedule_appointment(patient1, date(2026, 3, 10), "General Checkup")
    doctor.schedule_appointment(patient2, date(2026, 3, 12), "Follow-up Visit")
    doctor.schedule_appointment(patient3, date(2026, 3, 15), "Blood Test Review")
    print()

    # --- Diagnose patients ---
    print("=== Diagnosing Patients ===")
    doctor.diagnose(patient1, "Mild Fever")
    doctor.diagnose(patient2, "Vitamin D Deficiency")
    doctor.diagnose(patient3, "Fractured Wrist")
    print()

    # --- Prescribe treatments ---
    print("=== Prescribing Treatments ===")
    doctor.prescribe_treatment(patient1, "Paracetamol 500mg — 3 times daily")
    doctor.prescribe_treatment(patient2, "Vitamin D supplements — once daily")
    doctor.prescribe_treatment(patient3, "Cast and rest for 6 weeks")
    print()

    # --- View patient details ---
    print("=== Patient Details ===")
    patient1.show_info()
    patient1.view_appointments()
    patient1.view_medical_history()
    print()

    patient2.show_info()
    patient2.view_appointments()
    patient2.view_medical_history()
    print()

    # --- Discharge a patient ---
    print("=== Discharging Patient ===")
    doctor.discharge_patient(patient1)
    print()

    print("=== After Discharge ===")
    patient1.show_info()
    doctor.show_info()


if __name__ == "__main__":
    main()
