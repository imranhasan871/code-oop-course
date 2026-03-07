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
 * How to run (using dotnet-script):
 *   dotnet script Practice13.cs
 *
 * Key Concepts:
 *   - Many-Many Association via a Prescription junction class
 *   - Conflict detection logic
 *   - Collection management
 */

using System;
using System.Collections.Generic;
using System.Linq;

class Practice13
{
    class Medication
    {
        public string Name { get; }
        public string Category { get; }

        public Medication(string name, string category)
        {
            Name = name;
            Category = category;
        }

        public void ShowInfo()
        {
            Console.WriteLine($"  {Name,-25} | Category: {Category}");
        }
    }

    class Patient
    {
        public string PatientId { get; }
        public string Name { get; }

        public Patient(string patientId, string name)
        {
            PatientId = patientId;
            Name = name;
        }
    }

    class Prescription
    {
        public Patient Patient { get; }
        public Medication Medication { get; }
        public DateTime PrescribedDate { get; }

        public Prescription(Patient patient, Medication medication, DateTime prescribedDate)
        {
            Patient = patient;
            Medication = medication;
            PrescribedDate = prescribedDate;
        }

        public void ShowInfo()
        {
            Console.WriteLine($"  {Patient.Name,-12} | {Medication.Name,-25} | " +
                              $"Category: {Medication.Category,-22} | Date: {PrescribedDate:yyyy-MM-dd}");
        }
    }

    static readonly string[][] ForbiddenCombinations =
    {
        new[] { "Antibiotics", "Statins" },
        new[] { "Muscle Relaxants", "CNS Depressants" },
        new[] { "Anti-Inflammatories", "Anticoagulants" },
    };

    class PrescriptionManager
    {
        private List<Prescription> prescriptions = new List<Prescription>();

        private HashSet<string> GetPatientCategories(Patient patient)
        {
            var categories = new HashSet<string>();
            foreach (var p in prescriptions)
            {
                if (p.Patient.PatientId == patient.PatientId)
                    categories.Add(p.Medication.Category);
            }
            return categories;
        }

        private string CheckConflict(HashSet<string> existing, string newCategory)
        {
            foreach (var pair in ForbiddenCombinations)
            {
                if (newCategory == pair[0] && existing.Contains(pair[1]))
                    return $"{pair[0]} cannot be combined with {pair[1]}";
                if (newCategory == pair[1] && existing.Contains(pair[0]))
                    return $"{pair[1]} cannot be combined with {pair[0]}";
            }
            return "";
        }

        public void Prescribe(Patient patient, Medication medication, DateTime date)
        {
            // Check duplicate
            if (prescriptions.Any(p => p.Patient.PatientId == patient.PatientId
                                       && p.Medication.Name == medication.Name))
            {
                Console.WriteLine($"  [Error] {medication.Name} is already prescribed to {patient.Name}.");
                return;
            }
            // Check conflict
            var existing = GetPatientCategories(patient);
            string conflict = CheckConflict(existing, medication.Category);
            if (!string.IsNullOrEmpty(conflict))
            {
                Console.WriteLine($"  [Error] Cannot prescribe {medication.Name} to {patient.Name}: {conflict}.");
                return;
            }
            prescriptions.Add(new Prescription(patient, medication, date));
            Console.WriteLine($"  [OK] {medication.Name} prescribed to {patient.Name}.");
        }

        public List<Prescription> GetPatientPrescriptions(Patient patient)
        {
            return prescriptions.Where(p => p.Patient.PatientId == patient.PatientId).ToList();
        }

        public List<Prescription> GetMedicationPatients(Medication medication)
        {
            return prescriptions.Where(p => p.Medication.Name == medication.Name).ToList();
        }

        public void ShowAll()
        {
            if (prescriptions.Count == 0)
            {
                Console.WriteLine("  No prescriptions recorded.");
                return;
            }
            Console.WriteLine($"  {"Patient",-12} | {"Medication",-25} | {"Category",-22} | Date");
            Console.WriteLine("  " + new string('-', 80));
            foreach (var p in prescriptions)
                p.ShowInfo();
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        var amoxicillin = new Medication("Amoxicillin", "Antibiotics");
        var ibuprofen = new Medication("Ibuprofen", "Anti-Inflammatories");
        var diazepam = new Medication("Diazepam", "Muscle Relaxants");
        var coughSyrup = new Medication("Dextromethorphan", "Cough Syrup");
        var insulin = new Medication("Insulin", "Insulin");
        var warfarin = new Medication("Warfarin", "Anticoagulants");
        var phenobarbital = new Medication("Phenobarbital", "CNS Depressants");
        var atorvastatin = new Medication("Atorvastatin", "Statins");

        Console.WriteLine("=== Available Medications ===");
        foreach (var m in new[] { amoxicillin, ibuprofen, diazepam, coughSyrup,
                                   insulin, warfarin, phenobarbital, atorvastatin })
            m.ShowInfo();
        Console.WriteLine();

        var patient1 = new Patient("PAT-001", "Tareq");
        var patient2 = new Patient("PAT-002", "Afsana");
        var patient3 = new Patient("PAT-003", "Imtiaz");

        var manager = new PrescriptionManager();
        var today = DateTime.Today;

        Console.WriteLine("=== Prescribing Medications ===");
        manager.Prescribe(patient1, amoxicillin, today);
        manager.Prescribe(patient1, coughSyrup, today);
        manager.Prescribe(patient1, atorvastatin, today); // CONFLICT
        Console.WriteLine();

        manager.Prescribe(patient2, diazepam, today);
        manager.Prescribe(patient2, insulin, today);
        manager.Prescribe(patient2, phenobarbital, today); // CONFLICT
        Console.WriteLine();

        manager.Prescribe(patient3, ibuprofen, today);
        manager.Prescribe(patient3, coughSyrup, today);
        manager.Prescribe(patient3, warfarin, today); // CONFLICT
        Console.WriteLine();

        Console.WriteLine("=== All Prescriptions ===");
        manager.ShowAll();

        foreach (var patient in new[] { patient1, patient2, patient3 })
        {
            Console.WriteLine($"=== Prescriptions for {patient.Name} ===");
            foreach (var p in manager.GetPatientPrescriptions(patient))
                p.ShowInfo();
            Console.WriteLine();
        }

        Console.WriteLine("=== Patients on Dextromethorphan (Cough Syrup) ===");
        foreach (var p in manager.GetMedicationPatients(coughSyrup))
            Console.WriteLine($"  {p.Patient.Name} (prescribed: {p.PrescribedDate:yyyy-MM-dd})");
        Console.WriteLine();
    }
}
