/**
 * Practice 11: 1-1 Association — Car & License Plate
 * Task: Manage car details and their associated license plates.
 *       Validate license plate expiration. Track owner, manufacturer,
 *       model, year. Car age must be 20 years or less for renewal.
 *
 * How to run:
 *   go run practice-11.go
 *
 * Key Concepts:
 *   - 1-1 Association (Car has a LicensePlate)
 *   - Date-based expiration validation
 *   - Age calculation for renewal eligibility
 */

package main

import (
	"fmt"
	"time"
)

/** LicensePlate with plate number, registration and expiration dates. */
type LicensePlate struct {
	PlateNumber      string
	RegistrationDate time.Time
	ExpirationDate   time.Time
}

func (lp *LicensePlate) IsValid() bool {
	return !time.Now().After(lp.ExpirationDate)
}

func (lp *LicensePlate) ShowInfo() {
	status := "Valid"
	if !lp.IsValid() {
		status = "Expired"
	}
	fmt.Println("  Plate Number       :", lp.PlateNumber)
	fmt.Println("  Registration Date  :", lp.RegistrationDate.Format("2006-01-02"))
	fmt.Println("  Expiration Date    :", lp.ExpirationDate.Format("2006-01-02"))
	fmt.Println("  Status             :", status)
}

/** Car with owner, manufacturer, model, year, and a LicensePlate (1-1). */
type Car struct {
	Owner        string
	Manufacturer string
	Model        string
	Year         int
	LicensePlate *LicensePlate
}

const maxAgeForRenewal = 20

func (c *Car) CarAge() int {
	return time.Now().Year() - c.Year
}

func (c *Car) QualifiesForRenewal() bool {
	return c.CarAge() <= maxAgeForRenewal
}

func (c *Car) RenewRegistration(newExpiration time.Time) {
	if !c.QualifiesForRenewal() {
		fmt.Printf("  [Error] %s %s (%d) is %d years old. Maximum age for renewal is %d years.\n",
			c.Manufacturer, c.Model, c.Year, c.CarAge(), maxAgeForRenewal)
		return
	}
	c.LicensePlate.ExpirationDate = newExpiration
	fmt.Printf("  [OK] Registration renewed for %s %s (%s). New expiration: %s\n",
		c.Manufacturer, c.Model, c.LicensePlate.PlateNumber, newExpiration.Format("2006-01-02"))
}

func (c *Car) ShowInfo() {
	qualifies := "Yes"
	if !c.QualifiesForRenewal() {
		qualifies = "No"
	}
	fmt.Println("  Owner              :", c.Owner)
	fmt.Println("  Manufacturer       :", c.Manufacturer)
	fmt.Println("  Model              :", c.Model)
	fmt.Println("  Year               :", c.Year)
	fmt.Printf("  Car Age            : %d years\n", c.CarAge())
	fmt.Println("  Qualifies Renewal  :", qualifies)
	c.LicensePlate.ShowInfo()
	fmt.Println()
}

func main() {
	plate1 := &LicensePlate{"DHK-GA-1234", time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2028, 1, 1, 0, 0, 0, 0, time.Local)}
	plate2 := &LicensePlate{"DHK-KA-5678", time.Date(2020, 6, 15, 0, 0, 0, 0, time.Local), time.Date(2025, 6, 15, 0, 0, 0, 0, time.Local)}
	plate3 := &LicensePlate{"CTG-GA-9012", time.Date(2015, 3, 10, 0, 0, 0, 0, time.Local), time.Date(2025, 3, 10, 0, 0, 0, 0, time.Local)}

	car1 := &Car{"Tareq", "Toyota", "Corolla", 2020, plate1}
	car2 := &Car{"Afsana", "Honda", "Civic", 2018, plate2}
	car3 := &Car{"Imtiaz", "Nissan", "Sunny", 2000, plate3}

	fmt.Println("=== Car Information ===")
	car1.ShowInfo()
	car2.ShowInfo()
	car3.ShowInfo()

	fmt.Println("=== License Plate Validation ===")
	for _, car := range []*Car{car1, car2, car3} {
		valid := "Valid"
		if !car.LicensePlate.IsValid() {
			valid = "EXPIRED"
		}
		fmt.Printf("  %s %s (%s): %s\n", car.Manufacturer, car.Model, car.LicensePlate.PlateNumber, valid)
	}
	fmt.Println()

	fmt.Println("=== Registration Renewal ===")
	car1.RenewRegistration(time.Date(2033, 1, 1, 0, 0, 0, 0, time.Local))
	car2.RenewRegistration(time.Date(2030, 6, 15, 0, 0, 0, 0, time.Local))
	fmt.Println()

	fmt.Println("=== Renew Old Car (should fail if too old) ===")
	oldCar := &Car{"Robin", "Toyota", "Corona", 2003,
		&LicensePlate{"DHK-GA-0001", time.Date(2005, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)}}
	oldCar.ShowInfo()
	oldCar.RenewRegistration(time.Date(2030, 1, 1, 0, 0, 0, 0, time.Local))
	fmt.Println()

	fmt.Println("=== Very Old Car ===")
	ancientCar := &Car{"Karim", "Datsun", "120Y", 1980,
		&LicensePlate{"DHK-TA-0002", time.Date(1985, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)}}
	ancientCar.ShowInfo()
	ancientCar.RenewRegistration(time.Date(2030, 1, 1, 0, 0, 0, 0, time.Local))
}
