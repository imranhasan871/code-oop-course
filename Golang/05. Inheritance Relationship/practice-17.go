/**
 * Practice 17: Vehicle Rental System — Inheritance
 * Task: A vehicle rental company rents out Cars, Bikes, and Trucks.
 *       All vehicles share Brand, Model, and Year of Manufacture.
 *       Each type has different rental pricing and discount rules.
 *
 * Rules:
 *   - Car  : $50/day. 10% discount if the car is older than 5 years.
 *   - Bike : $15/day. 15% discount if rental period is more than 7 days.
 *   - Truck: $100/day + $100/day additional maintenance fee.
 *
 * How to run:
 *   go run practice-17.go
 *
 * Key Concepts:
 *   - Struct embedding (Go's approach to inheritance)
 *   - Shared fields via embedded struct
 *   - Type-specific methods
 */

package main

import "fmt"

/** Base struct for all rental vehicles. */
type Vehicle17 struct {
	Brand string
	Model string
	Year  int
}

func (v *Vehicle17) ShowInfo() {
	fmt.Println("  Brand :", v.Brand)
	fmt.Println("  Model :", v.Model)
	fmt.Println("  Year  :", v.Year)
}

/** Car: $50/day, 10% discount if older than 5 years. */
type Car17 struct {
	Vehicle17
}

func (c *Car17) CalculateRentalCost(days int) float64 {
	cost := 50.0 * float64(days)
	if 2026-c.Year > 5 {
		cost *= 0.9
	}
	return cost
}

/** Bike: $15/day, 15% discount if rental > 7 days. */
type Bike17 struct {
	Vehicle17
}

func (b *Bike17) CalculateRentalCost(days int) float64 {
	cost := 15.0 * float64(days)
	if days > 7 {
		cost *= 0.85
	}
	return cost
}

/** Truck: $100/day + $100/day additional maintenance fee. */
type Truck17 struct {
	Vehicle17
}

func (t *Truck17) CalculateRentalCost(days int) float64 {
	return (100.0 + 100.0) * float64(days)
}

func main() {
	car := &Car17{Vehicle17{"Toyota", "Corolla", 2019}}
	bike := &Bike17{Vehicle17{"Yamaha", "R15", 2023}}
	truck := &Truck17{Vehicle17{"Volvo", "FH16", 2022}}

	days := 10

	fmt.Println("=== Car Rental ===")
	car.ShowInfo()
	fmt.Printf("  Rental (%d days): $%.2f\n", days, car.CalculateRentalCost(days))
	fmt.Println()

	fmt.Println("=== Bike Rental ===")
	bike.ShowInfo()
	fmt.Printf("  Rental (%d days): $%.2f\n", days, bike.CalculateRentalCost(days))
	fmt.Println()

	fmt.Println("=== Truck Rental ===")
	truck.ShowInfo()
	fmt.Printf("  Rental (%d days): $%.2f\n", days, truck.CalculateRentalCost(days))
	fmt.Println()

	fmt.Println("=== Discount Scenarios ===")

	oldCar := &Car17{Vehicle17{"Honda", "Civic", 2018}}
	fmt.Println("Old Car (2018):")
	oldCar.ShowInfo()
	fmt.Printf("  Rental (%d days): $%.2f\n", days, oldCar.CalculateRentalCost(days))
	fmt.Printf("  (10%% discount applied — car is %d years old)\n", 2026-oldCar.Year)
	fmt.Println()

	shortBike := &Bike17{Vehicle17{"Honda", "CBR", 2024}}
	shortDays := 5
	fmt.Printf("Bike for %d days (no discount):\n", shortDays)
	shortBike.ShowInfo()
	fmt.Printf("  Rental (%d days): $%.2f\n", shortDays, shortBike.CalculateRentalCost(shortDays))
	fmt.Println()

	longBike := &Bike17{Vehicle17{"Suzuki", "Gixxer", 2023}}
	longDays := 10
	fmt.Printf("Bike for %d days (15%% discount):\n", longDays)
	longBike.ShowInfo()
	fmt.Printf("  Rental (%d days): $%.2f\n", longDays, longBike.CalculateRentalCost(longDays))
}
