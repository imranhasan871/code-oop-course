"""
Practice 13: Many-Many Association — Patient & Medications
Task: Patients are prescribed multiple medications, and each medication
      can be prescribed to many patients. Certain medication combinations
      are forbidden.

      Forbidden combinations:
        - Antibiotics & Statins
        - Muscle Relaxants & CNS Depressants
        - Anti-Inflammatories & Anticoagulants

How to run:
  python practice_13.py

Key Concepts:
  - Many-Many Association via a Prescription junction class
  - Conflict detection logic
  - Collection management
"""

from datetime import date


class Medication:
    """Medication with name and category."""

    def __init__(self, name: str, category: str):
        """Creates a new Medication."""
        self.name = name
        self.category = category

    def show_info(self):
        """Prints medication details."""
        print(f"  {self.name:<25} | Category: {self.category}")


class Patient:
    """Patient who can be prescribed multiple medications."""

    def __init__(self, patient_id: str, name: str):
        """Creates a new Patient."""
        self.patient_id = patient_id
        self.name = name

    def show_info(self):
        """Prints patient details."""
        print(f"  Patient ID : {self.patient_id}")
        print(f"  Name       : {self.name}")


class Prescription:
    """Represents a single prescription linking a patient to a medication."""

    def __init__(self, patient: Patient, medication: Medication, prescribed_date: date):
        """Creates a new Prescription."""
        self.patient = patient
        self.medication = medication
        self.prescribed_date = prescribed_date

    def show_info(self):
        """Prints prescription details."""
        print(f"  {self.patient.name:<12} | {self.medication.name:<25} | "
              f"Category: {self.medication.category:<22} | Date: {self.prescribed_date}")


# Forbidden medication category combinations
FORBIDDEN_COMBINATIONS = [
    ("Antibiotics", "Statins"),
    ("Muscle Relaxants", "CNS Depressants"),
    ("Anti-Inflammatories", "Anticoagulants"),
]


class PrescriptionManager:
    """Manages all prescriptions with conflict checking."""

    def __init__(self):
        """Creates a new PrescriptionManager."""
        self.prescriptions = []

    def _get_patient_categories(self, patient: Patient) -> set:
        """Returns the set of medication categories currently prescribed to a patient."""
        categories = set()
        for p in self.prescriptions:
            if p.patient.patient_id == patient.patient_id:
                categories.add(p.medication.category)
        return categories

    def _check_conflict(self, existing_categories: set, new_category: str) -> str:
        """Checks if adding new_category would conflict with existing categories.
        Returns a conflict message or empty string."""
        for cat_a, cat_b in FORBIDDEN_COMBINATIONS:
            if new_category == cat_a and cat_b in existing_categories:
                return f"{cat_a} cannot be combined with {cat_b}"
            if new_category == cat_b and cat_a in existing_categories:
                return f"{cat_b} cannot be combined with {cat_a}"
        return ""

    def prescribe(self, patient: Patient, medication: Medication, prescribed_date: date):
        """Prescribes a medication to a patient after conflict checking."""
        # Check for duplicate prescription
        for p in self.prescriptions:
            if p.patient.patient_id == patient.patient_id and p.medication.name == medication.name:
                print(f"  [Error] {medication.name} is already prescribed to {patient.name}.")
                return

        # Check for conflicts
        existing_categories = self._get_patient_categories(patient)
        conflict = self._check_conflict(existing_categories, medication.category)
        if conflict:
            print(f"  [Error] Cannot prescribe {medication.name} to {patient.name}: {conflict}.")
            return

        prescription = Prescription(patient, medication, prescribed_date)
        self.prescriptions.append(prescription)
        print(f"  [OK] {medication.name} prescribed to {patient.name}.")

    def get_patient_prescriptions(self, patient: Patient) -> list:
        """Returns all prescriptions for a given patient."""
        return [p for p in self.prescriptions if p.patient.patient_id == patient.patient_id]

    def get_medication_patients(self, medication: Medication) -> list:
        """Returns all prescriptions for a given medication."""
        return [p for p in self.prescriptions if p.medication.name == medication.name]

    def show_all(self):
        """Prints all prescriptions."""
        if not self.prescriptions:
            print("  No prescriptions recorded.")
            return
        print(f"  {'Patient':<12} | {'Medication':<25} | {'Category':<22} | Date")
        print("  " + "-" * 80)
        for p in self.prescriptions:
            p.show_info()
        print()


def main():
    # --- Create medications ---
    amoxicillin = Medication("Amoxicillin", "Antibiotics")
    ibuprofen = Medication("Ibuprofen", "Anti-Inflammatories")
    diazepam = Medication("Diazepam", "Muscle Relaxants")
    cough_syrup = Medication("Dextromethorphan", "Cough Syrup")
    insulin = Medication("Insulin", "Insulin")
    warfarin = Medication("Warfarin", "Anticoagulants")
    phenobarbital = Medication("Phenobarbital", "CNS Depressants")
    atorvastatin = Medication("Atorvastatin", "Statins")

    print("=== Available Medications ===")
    for med in [amoxicillin, ibuprofen, diazepam, cough_syrup,
                insulin, warfarin, phenobarbital, atorvastatin]:
        med.show_info()
    print()

    # --- Create patients ---
    patient1 = Patient("PAT-001", "Tareq")
    patient2 = Patient("PAT-002", "Afsana")
    patient3 = Patient("PAT-003", "Imtiaz")

    # --- Create prescription manager ---
    manager = PrescriptionManager()
    today = date.today()

    # --- Prescribe medications ---
    print("=== Prescribing Medications ===")

    # Patient 1: Tareq — Antibiotics + Cough Syrup (OK)
    manager.prescribe(patient1, amoxicillin, today)
    manager.prescribe(patient1, cough_syrup, today)
    # Tareq: Antibiotics + Statins (CONFLICT)
    manager.prescribe(patient1, atorvastatin, today)
    print()

    # Patient 2: Afsana — Muscle Relaxants + Insulin (OK)
    manager.prescribe(patient2, diazepam, today)
    manager.prescribe(patient2, insulin, today)
    # Afsana: Muscle Relaxants + CNS Depressants (CONFLICT)
    manager.prescribe(patient2, phenobarbital, today)
    print()

    # Patient 3: Imtiaz — Anti-Inflammatories + Cough Syrup (OK)
    manager.prescribe(patient3, ibuprofen, today)
    manager.prescribe(patient3, cough_syrup, today)
    # Imtiaz: Anti-Inflammatories + Anticoagulants (CONFLICT)
    manager.prescribe(patient3, warfarin, today)
    print()

    # --- Show all prescriptions ---
    print("=== All Prescriptions ===")
    manager.show_all()

    # --- Show prescriptions per patient ---
    for patient in [patient1, patient2, patient3]:
        print(f"=== Prescriptions for {patient.name} ===")
        for p in manager.get_patient_prescriptions(patient):
            p.show_info()
        print()

    # --- Show which patients are on a specific medication ---
    print("=== Patients on Dextromethorphan (Cough Syrup) ===")
    for p in manager.get_medication_patients(cough_syrup):
        print(f"  {p.patient.name} (prescribed: {p.prescribed_date})")
    print()


if __name__ == "__main__":
    main()
