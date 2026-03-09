/**
 * Practice 18: Vehicle Rental — Method Overriding & Constructor Chaining
 * Task: Enhance the vehicle rental system with method overriding and
 *       constructor chaining between base struct and embedded structs.
 *
 * Changes from Practice 17:
 *   - Vehicle interface with CalculateRentalCost() and String() methods
 *   - Each type implements the interface (Go's approach to method overriding)
 *   - Constructor functions initialize embedded struct (constructor chaining)
 *
 * How to run:
 *   go run practice-18.go
 *
 * Key Concepts:
 *   - Interface-based polymorphism (Go's method overriding)
 *   - Struct embedding (Go's constructor chaining)
 *   - String() method (Go's toString equivalent)
 */

package main

import "fmt"

/** Vehicle18 interface — defines the contract for all rental vehicles. */
type Vehicle18 interface {
	CalculateRentalCost(days int) float64
	String() string
	GetBrand() string
	GetModel() string
	GetYear() int
}

/** BaseVehicle18 holds shared fields. */
type BaseVehicle18 struct {
	Brand string
	Model string
	Year  int
}

func (v *BaseVehicle18) GetBrand() string { return v.Brand }
func (v *BaseVehicle18) GetModel() string { return v.Model }
func (v *BaseVehicle18) GetYear() int     { return v.Year }

func (v *BaseVehicle18) String() string {
	return fmt.Sprintf("%s %s (%d)", v.Brand, v.Model, v.Year)
}

/** Car18: $50/day, 10% discount if older than 5 years. */
type Car18 struct {
	BaseVehicle18 // Embedding (constructor chaining via struct init)
}

func NewCar18(brand, model string, year int) *Car18 {
	return &Car18{BaseVehicle18{brand, model, year}}
}

func (c *Car18) CalculateRentalCost(days int) float64 {
	cost := 50.0 * float64(days)
	if 2026-c.Year > 5 {
		cost *= 0.9
	}
	return cost
}

func (c *Car18) String() string {
	return "[Car] " + c.BaseVehicle18.String()
}

/** Bike18: $15/day, 15% discount if rental > 7 days. */
type Bike18 struct {
	BaseVehicle18
}

func NewBike18(brand, model string, year int) *Bike18 {
	return &Bike18{BaseVehicle18{brand, model, year}}
}

func (b *Bike18) CalculateRentalCost(days int) float64 {
	cost := 15.0 * float64(days)
	if days > 7 {
		cost *= 0.85
	}
	return cost
}

func (b *Bike18) String() string {
	return "[Bike] " + b.BaseVehicle18.String()
}

/** Truck18: $100/day + $100/day additional fee. */
type Truck18 struct {
	BaseVehicle18
}

func NewTruck18(brand, model string, year int) *Truck18 {
	return &Truck18{BaseVehicle18{brand, model, year}}
}

func (t *Truck18) CalculateRentalCost(days int) float64 {
	return (100.0 + 100.0) * float64(days)
}

func (t *Truck18) String() string {
	return "[Truck] " + t.BaseVehicle18.String()
}

func main() {
	var car Vehicle18 = NewCar18("Toyota", "Corolla", 2019)
	var bike Vehicle18 = NewBike18("Yamaha", "R15", 2023)
	var truck Vehicle18 = NewTruck18("Volvo", "FH16", 2022)

	days := 10

	fmt.Println("=== Vehicle Rental — Method Overriding ===")
	fmt.Println()

	for _, vehicle := range []Vehicle18{car, bike, truck} {
		fmt.Println("  Vehicle :", vehicle.String())                                          // String override
		fmt.Printf("  Rental (%d days): $%.2f\n", days, vehicle.CalculateRentalCost(days))    // Method override
		fmt.Println()
	}

	fmt.Println("=== String Representation (String) ===")
	fmt.Println("  Car   :", car.String())
	fmt.Println("  Bike  :", bike.String())
	fmt.Println("  Truck :", truck.String())
	fmt.Println()

	fmt.Println("=== Discount Scenarios ===")

	var oldCar Vehicle18 = NewCar18("Honda", "Civic", 2018)
	fmt.Println(" ", oldCar.String())
	fmt.Printf("  Age: %d years → 10%% discount applied\n", 2026-oldCar.GetYear())
	fmt.Printf("  Rental (%d days): $%.2f\n", days, oldCar.CalculateRentalCost(days))
	fmt.Println()

	var shortBike Vehicle18 = NewBike18("Honda", "CBR", 2024)
	shortDays := 5
	fmt.Println(" ", shortBike.String())
	fmt.Printf("  Rental (%d days): $%.2f (no discount)\n", shortDays, shortBike.CalculateRentalCost(shortDays))
	fmt.Println()

	var longBike Vehicle18 = NewBike18("Suzuki", "Gixxer", 2023)
	longDays := 10
	fmt.Println(" ", longBike.String())
	fmt.Printf("  Rental (%d days): $%.2f (15%% discount)\n", longDays, longBike.CalculateRentalCost(longDays))
}
