/**
 * Practice 06: Car Rental System (OOAD)
 * Task: Model a car rental company with a fleet of vehicles.
 *       Customers can rent and return cars. Track availability and maintenance.
 *
 * How to run:
 *   go run practice-06.go
 *
 * Key Concepts:
 *   - Multiple structs working together (Car, Customer)
 *   - Pointer references between structs
 *   - State management (available, rented, maintenance)
 */

package main

import "fmt"

/** Car represents a vehicle in the rental fleet. */
type Car struct {
	LicensePlate     string
	Brand            string
	Model            string
	CarType          string  // sedan, suv, van
	DailyRate        float64
	IsAvailable      bool
	NeedsMaintenance bool
}

/** NewCar creates a new Car that is available for rent. */
func NewCar(licensePlate, brand, model, carType string, dailyRate float64) Car {
	return Car{
		LicensePlate:     licensePlate,
		Brand:            brand,
		Model:            model,
		CarType:          carType,
		DailyRate:        dailyRate,
		IsAvailable:      true,
		NeedsMaintenance: false,
	}
}

/** Rent marks the car as rented (not available). */
func (c *Car) Rent() bool {
	if !c.IsAvailable {
		fmt.Printf("  [Error] %s %s (%s) is not available for rent.\n", c.Brand, c.Model, c.LicensePlate)
		return false
	}
	if c.NeedsMaintenance {
		fmt.Printf("  [Error] %s %s (%s) needs maintenance and cannot be rented.\n", c.Brand, c.Model, c.LicensePlate)
		return false
	}
	c.IsAvailable = false
	return true
}

/** Return marks the car as available again. */
func (c *Car) Return() {
	c.IsAvailable = true
	fmt.Printf("  [OK] %s %s (%s) has been returned.\n", c.Brand, c.Model, c.LicensePlate)
}

/** SendToMaintenance marks the car as needing maintenance. */
func (c *Car) SendToMaintenance() {
	c.NeedsMaintenance = true
	c.IsAvailable = false
	fmt.Printf("  [OK] %s %s (%s) sent to maintenance.\n", c.Brand, c.Model, c.LicensePlate)
}

/** CompleteMaintenance marks maintenance as done and makes car available. */
func (c *Car) CompleteMaintenance() {
	c.NeedsMaintenance = false
	c.IsAvailable = true
	fmt.Printf("  [OK] %s %s (%s) maintenance completed. Now available.\n", c.Brand, c.Model, c.LicensePlate)
}

/** ShowInfo prints car details. */
func (c *Car) ShowInfo() {
	status := "Available"
	if c.NeedsMaintenance {
		status = "In Maintenance"
	} else if !c.IsAvailable {
		status = "Rented"
	}
	fmt.Printf("  [%s] %s %s | Type: %s | Rate: %.2f/day | Plate: %s\n",
		status, c.Brand, c.Model, c.CarType, c.DailyRate, c.LicensePlate)
}

/** Customer represents a person who can rent a car. */
type Customer struct {
	Name      string
	Phone     string
	RentedCar *Car
}

/** NewCustomer creates a new Customer with no rented car. */
func NewCustomer(name, phone string) Customer {
	return Customer{Name: name, Phone: phone, RentedCar: nil}
}

/** RentCar allows the customer to rent an available car. */
func (cu *Customer) RentCar(car *Car) {
	if cu.RentedCar != nil {
		fmt.Printf("  [Error] %s already has a rented car (%s %s).\n",
			cu.Name, cu.RentedCar.Brand, cu.RentedCar.Model)
		return
	}
	if car.Rent() {
		cu.RentedCar = car
		fmt.Printf("  [OK] %s rented %s %s (%s) at %.2f/day.\n",
			cu.Name, car.Brand, car.Model, car.LicensePlate, car.DailyRate)
	}
}

/** ReturnCar allows the customer to return their rented car. */
func (cu *Customer) ReturnCar() {
	if cu.RentedCar == nil {
		fmt.Printf("  [Error] %s has no car to return.\n", cu.Name)
		return
	}
	cu.RentedCar.Return()
	fmt.Printf("  [OK] %s returned the car.\n", cu.Name)
	cu.RentedCar = nil
}

/** ShowInfo prints customer details. */
func (cu *Customer) ShowInfo() {
	fmt.Printf("  Customer: %s | Phone: %s", cu.Name, cu.Phone)
	if cu.RentedCar != nil {
		fmt.Printf(" | Rented: %s %s (%s)", cu.RentedCar.Brand, cu.RentedCar.Model, cu.RentedCar.LicensePlate)
	} else {
		fmt.Print(" | Rented: None")
	}
	fmt.Println()
}

func main() {
	// --- Create fleet of cars ---
	car1 := NewCar("DHK-1001", "Toyota", "Corolla", "Sedan", 3500)
	car2 := NewCar("DHK-1002", "Honda", "CR-V", "SUV", 5000)
	car3 := NewCar("DHK-1003", "Toyota", "HiAce", "Van", 7000)

	fmt.Println("=== Fleet Status ===")
	car1.ShowInfo()
	car2.ShowInfo()
	car3.ShowInfo()
	fmt.Println()

	// --- Create customers ---
	cust1 := NewCustomer("Tareq", "01700-000001")
	cust2 := NewCustomer("Afsana", "01700-000002")

	// --- Rent cars ---
	fmt.Println("=== Renting Cars ===")
	cust1.RentCar(&car1)
	cust2.RentCar(&car2)
	fmt.Println()

	fmt.Println("=== After Renting ===")
	car1.ShowInfo()
	car2.ShowInfo()
	car3.ShowInfo()
	cust1.ShowInfo()
	cust2.ShowInfo()
	fmt.Println()

	// --- Try to rent already rented car ---
	fmt.Println("=== Try to rent an already rented car ===")
	cust2.RentCar(&car1)
	fmt.Println()

	// --- Return car ---
	fmt.Println("=== Tareq returns car ===")
	cust1.ReturnCar()
	cust1.ShowInfo()
	car1.ShowInfo()
	fmt.Println()

	// --- Send car to maintenance ---
	fmt.Println("=== Send car1 to maintenance ===")
	car1.SendToMaintenance()
	car1.ShowInfo()
	fmt.Println()

	// --- Try renting car in maintenance ---
	fmt.Println("=== Try renting car in maintenance ===")
	cust1.RentCar(&car1)
	fmt.Println()

	// --- Complete maintenance ---
	fmt.Println("=== Complete maintenance ===")
	car1.CompleteMaintenance()
	car1.ShowInfo()
}
