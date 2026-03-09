/**
 * Practice 19: Vehicle Rental — Runtime Polymorphism
 * Task: Display brand, model, year, and rental price for one car,
 *       two bikes, and one truck if each is rented for 10 days.
 *
 * How to run:
 *   go run practice-19.go
 *
 * Key Concepts:
 *   - Interface-based polymorphism (Go's runtime polymorphism)
 *   - Storing different types in a single interface slice
 *   - Method dispatch at runtime
 */

package main

import (
	"fmt"
	"reflect"
)

/** Vehicle19 interface — defines the contract for all rental vehicles. */
type Vehicle19 interface {
	CalculateRentalCost(days int) float64
	GetVehicleType() string
	String() string
	GetBrand() string
	GetModel() string
	GetYear() int
}

/** BaseVehicle19 holds shared fields. */
type BaseVehicle19 struct {
	Brand string
	Model string
	Year  int
}

func (v *BaseVehicle19) GetBrand() string { return v.Brand }
func (v *BaseVehicle19) GetModel() string { return v.Model }
func (v *BaseVehicle19) GetYear() int     { return v.Year }

func (v *BaseVehicle19) String() string {
	return fmt.Sprintf("%s %s (%d)", v.Brand, v.Model, v.Year)
}

/** Car19: $50/day, 10% discount if older than 5 years. */
type Car19 struct {
	BaseVehicle19
}

func NewCar19(brand, model string, year int) *Car19 {
	return &Car19{BaseVehicle19{brand, model, year}}
}

func (c *Car19) CalculateRentalCost(days int) float64 {
	cost := 50.0 * float64(days)
	if 2026-c.Year > 5 {
		cost *= 0.9
	}
	return cost
}

func (c *Car19) GetVehicleType() string { return "Car" }
func (c *Car19) String() string         { return "[Car] " + c.BaseVehicle19.String() }

/** Bike19: $15/day, 15% discount if rental > 7 days. */
type Bike19 struct {
	BaseVehicle19
}

func NewBike19(brand, model string, year int) *Bike19 {
	return &Bike19{BaseVehicle19{brand, model, year}}
}

func (b *Bike19) CalculateRentalCost(days int) float64 {
	cost := 15.0 * float64(days)
	if days > 7 {
		cost *= 0.85
	}
	return cost
}

func (b *Bike19) GetVehicleType() string { return "Bike" }
func (b *Bike19) String() string         { return "[Bike] " + b.BaseVehicle19.String() }

/** Truck19: $100/day + $100/day additional fee. */
type Truck19 struct {
	BaseVehicle19
}

func NewTruck19(brand, model string, year int) *Truck19 {
	return &Truck19{BaseVehicle19{brand, model, year}}
}

func (t *Truck19) CalculateRentalCost(days int) float64 {
	return (100.0 + 100.0) * float64(days)
}

func (t *Truck19) GetVehicleType() string { return "Truck" }
func (t *Truck19) String() string         { return "[Truck] " + t.BaseVehicle19.String() }

func main() {
	// Interface slice — all types stored as Vehicle19 (polymorphism)
	vehicles := []Vehicle19{
		NewCar19("Toyota", "Corolla", 2019),  // 1 Car
		NewBike19("Yamaha", "R15", 2023),     // 2 Bikes
		NewBike19("Honda", "CBR", 2024),
		NewTruck19("Volvo", "FH16", 2022),    // 1 Truck
	}

	rentalDays := 10

	fmt.Printf("=== Vehicle Rental Report (%d Days) ===\n", rentalDays)
	fmt.Println()
	fmt.Printf("  %-8s | %-10s | %-10s | %-6s | %12s\n",
		"Type", "Brand", "Model", "Year", "Rental Cost")
	fmt.Println("  " + "------------------------------------------------------------")

	totalCost := 0.0
	for _, vehicle := range vehicles {
		cost := vehicle.CalculateRentalCost(rentalDays) // Runtime polymorphism
		totalCost += cost
		fmt.Printf("  %-8s | %-10s | %-10s | %-6d | $%10.2f\n",
			vehicle.GetVehicleType(), vehicle.GetBrand(), vehicle.GetModel(),
			vehicle.GetYear(), cost)
	}

	fmt.Println("  " + "------------------------------------------------------------")
	fmt.Printf("  %-8s | %-10s | %-10s | %-6s | $%10.2f\n",
		"Total", "", "", "", totalCost)
	fmt.Println()

	// Demonstrate runtime polymorphism
	fmt.Println("=== Runtime Polymorphism ===")
	fmt.Println()
	for _, vehicle := range vehicles {
		fmt.Println(" ", vehicle.String())
		fmt.Printf("    Type at runtime : %s\n", reflect.TypeOf(vehicle).Elem().Name())
		fmt.Printf("    Rental (10 days): $%.2f\n", vehicle.CalculateRentalCost(10))
		fmt.Println()
	}
}
