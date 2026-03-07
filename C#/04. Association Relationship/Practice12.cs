/**
 * Practice 12: 1-Many Association — Doctor & Patients
 * Task: Manage doctors and patients. Doctors: schedule appointments,
 *       diagnose, prescribe treatments, discharge patients.
 *       Patients: view appointments, check treatments, track history.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice12.cs
 *
 * Key Concepts:
 *   - 1-Many Association (Doctor has many Patients)
 *   - Object references and collection management
 *   - Business logic across associated objects
 */

using System;
using System.Collections.Generic;

class Practice12
{
    /** Patient with personal info, medical history, appointments, and treatments. */
    class Patient
    {
        public string PatientId { get; }
        public string Name { get; }
        public List<string> MedicalHistory { get; } = new List<string>();
        public List<string> Appointments { get; } = new List<string>();
        public List<string> Treatments { get; } = new List<string>();
        public bool IsAdmitted { get; set; } = true;

        public Patient(string patientId, string name)
        {
            PatientId = patientId;
            Name = name;
        }

        public void ViewAppointments()
        {
            if (Appointments.Count == 0)
            {
                Console.WriteLine($"  {Name} has no appointments.");
                return;
            }
            Console.WriteLine($"  Appointments for {Name}:");
            foreach (var appt in Appointments)
                Console.WriteLine($"    - {appt}");
        }

        public bool HasOngoingTreatments() => Treatments.Count > 0;

        public void ViewMedicalHistory()
        {
            if (MedicalHistory.Count == 0)
            {
                Console.WriteLine($"  {Name} has no medical history records.");
                return;
            }
            Console.WriteLine($"  Medical history for {Name}:");
            foreach (var record in MedicalHistory)
                Console.WriteLine($"    - {record}");
        }

        public void ShowInfo()
        {
            string status = IsAdmitted ? "Admitted" : "Discharged";
            Console.WriteLine($"  Patient ID    : {PatientId}");
            Console.WriteLine($"  Name          : {Name}");
            Console.WriteLine($"  Status        : {status}");
            Console.WriteLine($"  Diagnoses     : {MedicalHistory.Count}");
            Console.WriteLine($"  Treatments    : {Treatments.Count}");
            Console.WriteLine($"  Appointments  : {Appointments.Count}");
            Console.WriteLine();
        }
    }

    /** Doctor who manages multiple patients (1-Many association). */
    class Doctor
    {
        public string DoctorId { get; }
        public string Name { get; }
        public string Specialization { get; }
        public List<Patient> Patients { get; } = new List<Patient>();

        public Doctor(string doctorId, string name, string specialization)
        {
            DoctorId = doctorId;
            Name = name;
            Specialization = specialization;
        }

        public void AddPatient(Patient patient)
        {
            Patients.Add(patient);
            Console.WriteLine($"  [OK] {patient.Name} assigned to Dr. {Name}.");
        }

        public void ScheduleAppointment(Patient patient, DateTime date, string description)
        {
            if (!Patients.Contains(patient))
            {
                Console.WriteLine($"  [Error] {patient.Name} is not under Dr. {Name}'s care.");
                return;
            }
            patient.Appointments.Add($"{date:yyyy-MM-dd}: {description}");
            Console.WriteLine($"  [OK] Appointment scheduled for {patient.Name} on {date:yyyy-MM-dd}: {description}");
        }

        public void Diagnose(Patient patient, string diagnosis)
        {
            if (!Patients.Contains(patient))
            {
                Console.WriteLine($"  [Error] {patient.Name} is not under Dr. {Name}'s care.");
                return;
            }
            patient.MedicalHistory.Add(diagnosis);
            Console.WriteLine($"  [OK] Dr. {Name} diagnosed {patient.Name}: {diagnosis}");
        }

        public void PrescribeTreatment(Patient patient, string treatment)
        {
            if (!Patients.Contains(patient))
            {
                Console.WriteLine($"  [Error] {patient.Name} is not under Dr. {Name}'s care.");
                return;
            }
            if (patient.MedicalHistory.Count == 0)
            {
                Console.WriteLine($"  [Error] {patient.Name} has no diagnosis. Diagnose before prescribing.");
                return;
            }
            patient.Treatments.Add(treatment);
            Console.WriteLine($"  [OK] Dr. {Name} prescribed '{treatment}' to {patient.Name}.");
        }

        public void DischargePatient(Patient patient)
        {
            if (!Patients.Contains(patient))
            {
                Console.WriteLine($"  [Error] {patient.Name} is not under Dr. {Name}'s care.");
                return;
            }
            patient.IsAdmitted = false;
            patient.Treatments.Clear();
            Patients.Remove(patient);
            Console.WriteLine($"  [OK] {patient.Name} has been discharged by Dr. {Name}.");
        }

        public void ShowInfo()
        {
            Console.WriteLine($"  Doctor ID     : {DoctorId}");
            Console.WriteLine($"  Name          : Dr. {Name}");
            Console.WriteLine($"  Specialization: {Specialization}");
            Console.WriteLine($"  Patients      : {Patients.Count}");
            foreach (var p in Patients)
            {
                string status = p.IsAdmitted ? "Admitted" : "Discharged";
                Console.WriteLine($"    - {p.Name} ({p.PatientId}) [{status}]");
            }
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        var doctor = new Doctor("DOC-001", "Karim", "General Medicine");
        var patient1 = new Patient("PAT-001", "Tareq");
        var patient2 = new Patient("PAT-002", "Afsana");
        var patient3 = new Patient("PAT-003", "Imtiaz");

        Console.WriteLine("=== Assigning Patients ===");
        doctor.AddPatient(patient1);
        doctor.AddPatient(patient2);
        doctor.AddPatient(patient3);
        Console.WriteLine();

        Console.WriteLine("=== Doctor Info ===");
        doctor.ShowInfo();

        Console.WriteLine("=== Scheduling Appointments ===");
        doctor.ScheduleAppointment(patient1, new DateTime(2026, 3, 10), "General Checkup");
        doctor.ScheduleAppointment(patient2, new DateTime(2026, 3, 12), "Follow-up Visit");
        doctor.ScheduleAppointment(patient3, new DateTime(2026, 3, 15), "Blood Test Review");
        Console.WriteLine();

        Console.WriteLine("=== Diagnosing Patients ===");
        doctor.Diagnose(patient1, "Mild Fever");
        doctor.Diagnose(patient2, "Vitamin D Deficiency");
        doctor.Diagnose(patient3, "Fractured Wrist");
        Console.WriteLine();

        Console.WriteLine("=== Prescribing Treatments ===");
        doctor.PrescribeTreatment(patient1, "Paracetamol 500mg — 3 times daily");
        doctor.PrescribeTreatment(patient2, "Vitamin D supplements — once daily");
        doctor.PrescribeTreatment(patient3, "Cast and rest for 6 weeks");
        Console.WriteLine();

        Console.WriteLine("=== Patient Details ===");
        patient1.ShowInfo();
        patient1.ViewAppointments();
        patient1.ViewMedicalHistory();
        Console.WriteLine();

        patient2.ShowInfo();
        patient2.ViewAppointments();
        patient2.ViewMedicalHistory();
        Console.WriteLine();

        Console.WriteLine("=== Discharging Patient ===");
        doctor.DischargePatient(patient1);
        Console.WriteLine();

        Console.WriteLine("=== After Discharge ===");
        patient1.ShowInfo();
        doctor.ShowInfo();
    }
}
