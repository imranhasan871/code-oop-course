/**
 * Practice 14: Smart Parking System
 * Task: Manage vehicle entry, slot assignment, payment, and exit.
 *       Calculate parking charge based on duration.
 *       Update slot allocation and availability in real time.
 *
 * How to run:
 *   go run practice-14.go
 *
 * Key Concepts:
 *   - Association between Vehicle, ParkingSlot, and ParkingLot
 *   - Duration-based charge calculation
 *   - Real-time availability tracking
 */

package main

import (
	"fmt"
	"time"
)

/** Vehicle with license plate and type. */
type Vehicle14 struct {
	LicensePlate string
	VehicleType  string
}

/** ParkingSlot that can hold one vehicle at a time. */
type ParkingSlot14 struct {
	SlotId      string
	SlotType    string
	RatePerHour float64
	Vehicle     *Vehicle14
	EntryTime   time.Time
}

func (s *ParkingSlot14) IsAvailable() bool {
	return s.Vehicle == nil
}

func (s *ParkingSlot14) ShowInfo() {
	status := "Available"
	if !s.IsAvailable() {
		status = "Occupied by " + s.Vehicle.LicensePlate
	}
	fmt.Printf("  %-8s | Type: %-10s | Rate: %.0f/hr | %s\n",
		s.SlotId, s.SlotType, s.RatePerHour, status)
}

/** ParkingLot manages a collection of ParkingSlots. */
type ParkingLot14 struct {
	Name  string
	Slots []*ParkingSlot14
}

func NewParkingLot14(name string) *ParkingLot14 {
	return &ParkingLot14{Name: name}
}

func (lot *ParkingLot14) AddSlot(slot *ParkingSlot14) {
	lot.Slots = append(lot.Slots, slot)
}

func (lot *ParkingLot14) FindAvailableSlot(slotType string) *ParkingSlot14 {
	for _, slot := range lot.Slots {
		if slot.SlotType == slotType && slot.IsAvailable() {
			return slot
		}
	}
	return nil
}

func (lot *ParkingLot14) ParkVehicle(vehicle *Vehicle14, entryTime time.Time) *ParkingSlot14 {
	slot := lot.FindAvailableSlot(vehicle.VehicleType)
	if slot == nil {
		fmt.Printf("  [Error] No available %s slot for %s.\n", vehicle.VehicleType, vehicle.LicensePlate)
		return nil
	}
	slot.Vehicle = vehicle
	slot.EntryTime = entryTime
	fmt.Printf("  [OK] %s parked in slot %s at %s.\n",
		vehicle.LicensePlate, slot.SlotId, entryTime.Format("2006-01-02 15:04"))
	return slot
}

func (lot *ParkingLot14) ExitVehicle(licensePlate string, exitTime time.Time) float64 {
	for _, slot := range lot.Slots {
		if !slot.IsAvailable() && slot.Vehicle.LicensePlate == licensePlate {
			duration := exitTime.Sub(slot.EntryTime)
			hours := duration.Hours()
			if hours < 1 {
				hours = 1
			}
			charge := hours * slot.RatePerHour
			fmt.Printf("  [OK] %s exiting slot %s.\n", licensePlate, slot.SlotId)
			fmt.Printf("       Duration : %.1f hours\n", hours)
			fmt.Printf("       Rate     : %.0f/hr\n", slot.RatePerHour)
			fmt.Printf("       Charge   : %.2f\n", charge)
			slot.Vehicle = nil
			slot.EntryTime = time.Time{}
			return charge
		}
	}
	fmt.Printf("  [Error] Vehicle %s not found in parking lot.\n", licensePlate)
	return 0
}

func (lot *ParkingLot14) ShowStatus() {
	available := 0
	for _, slot := range lot.Slots {
		if slot.IsAvailable() {
			available++
		}
	}
	occupied := len(lot.Slots) - available
	fmt.Println("  Parking Lot:", lot.Name)
	fmt.Printf("  Total Slots: %d | Available: %d | Occupied: %d\n", len(lot.Slots), available, occupied)
	fmt.Println()
	for _, slot := range lot.Slots {
		slot.ShowInfo()
	}
	fmt.Println()
}

func main() {
	lot := NewParkingLot14("City Center Parking")
	lot.AddSlot(&ParkingSlot14{SlotId: "A-01", SlotType: "Car", RatePerHour: 50})
	lot.AddSlot(&ParkingSlot14{SlotId: "A-02", SlotType: "Car", RatePerHour: 50})
	lot.AddSlot(&ParkingSlot14{SlotId: "A-03", SlotType: "Car", RatePerHour: 50})
	lot.AddSlot(&ParkingSlot14{SlotId: "B-01", SlotType: "Bike", RatePerHour: 20})
	lot.AddSlot(&ParkingSlot14{SlotId: "B-02", SlotType: "Bike", RatePerHour: 20})
	lot.AddSlot(&ParkingSlot14{SlotId: "C-01", SlotType: "Truck", RatePerHour: 100})

	fmt.Println("=== Initial Parking Status ===")
	lot.ShowStatus()

	now := time.Date(2026, 3, 7, 9, 0, 0, 0, time.Local)
	fmt.Println("=== Parking Vehicles ===")
	lot.ParkVehicle(&Vehicle14{"DHK-1234", "Car"}, now)
	lot.ParkVehicle(&Vehicle14{"DHK-5678", "Car"}, now.Add(15*time.Minute))
	lot.ParkVehicle(&Vehicle14{"DHK-9012", "Bike"}, now.Add(30*time.Minute))
	lot.ParkVehicle(&Vehicle14{"CTG-3456", "Truck"}, now.Add(1*time.Hour))
	fmt.Println()

	fmt.Println("=== After Parking ===")
	lot.ShowStatus()

	fmt.Println("=== Vehicle Exits ===")
	lot.ExitVehicle("DHK-1234", now.Add(3*time.Hour))
	fmt.Println()
	lot.ExitVehicle("DHK-9012", now.Add(90*time.Minute))
	fmt.Println()
	lot.ExitVehicle("CTG-3456", now.Add(5*time.Hour))
	fmt.Println()

	fmt.Println("=== Final Parking Status ===")
	lot.ShowStatus()

	fmt.Println("=== Parking When Slots Are Limited ===")
	lot.ParkVehicle(&Vehicle14{"SYL-7777", "Car"}, now.Add(4*time.Hour))
	lot.ParkVehicle(&Vehicle14{"SYL-8888", "Car"}, now.Add(4*time.Hour))
	lot.ParkVehicle(&Vehicle14{"SYL-9999", "Car"}, now.Add(4*time.Hour))
}
