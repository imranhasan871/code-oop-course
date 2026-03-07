/**
 * Practice 11: 1-1 Association — Car & License Plate
 * Task: Manage car details and their associated license plates.
 *       Validate license plate expiration. Track owner, manufacturer,
 *       model, year. Car age must be 20 years or less for renewal.
 *
 * How to compile and run:
 *   javac Practice11.java
 *   java Practice11
 *
 * Key Concepts:
 *   - 1-1 Association (Car has a LicensePlate)
 *   - Date-based expiration validation
 *   - Age calculation for renewal eligibility
 */

import java.time.LocalDate;

public class Practice11 {

    /** LicensePlate with plate number, registration and expiration dates. */
    static class LicensePlate {
        private String plateNumber;
        private LocalDate registrationDate;
        private LocalDate expirationDate;

        public LicensePlate(String plateNumber, LocalDate registrationDate, LocalDate expirationDate) {
            this.plateNumber = plateNumber;
            this.registrationDate = registrationDate;
            this.expirationDate = expirationDate;
        }

        public boolean isValid() {
            return !LocalDate.now().isAfter(expirationDate);
        }

        public String getPlateNumber() { return plateNumber; }
        public LocalDate getExpirationDate() { return expirationDate; }
        public void setExpirationDate(LocalDate date) { this.expirationDate = date; }

        public void showInfo() {
            String status = isValid() ? "Valid" : "Expired";
            System.out.println("  Plate Number       : " + plateNumber);
            System.out.println("  Registration Date  : " + registrationDate);
            System.out.println("  Expiration Date    : " + expirationDate);
            System.out.println("  Status             : " + status);
        }
    }

    /** Car with owner, manufacturer, model, year, and a LicensePlate (1-1). */
    static class Car {
        private static final int MAX_AGE_FOR_RENEWAL = 20;

        private String owner;
        private String manufacturer;
        private String model;
        private int year;
        private LicensePlate licensePlate;

        public Car(String owner, String manufacturer, String model, int year, LicensePlate licensePlate) {
            this.owner = owner;
            this.manufacturer = manufacturer;
            this.model = model;
            this.year = year;
            this.licensePlate = licensePlate;
        }

        public int carAge() {
            return LocalDate.now().getYear() - year;
        }

        public boolean qualifiesForRenewal() {
            return carAge() <= MAX_AGE_FOR_RENEWAL;
        }

        public void renewRegistration(LocalDate newExpiration) {
            if (!qualifiesForRenewal()) {
                System.out.printf("  [Error] %s %s (%d) is %d years old. Maximum age for renewal is %d years.%n",
                        manufacturer, model, year, carAge(), MAX_AGE_FOR_RENEWAL);
                return;
            }
            licensePlate.setExpirationDate(newExpiration);
            System.out.printf("  [OK] Registration renewed for %s %s (%s). New expiration: %s%n",
                    manufacturer, model, licensePlate.getPlateNumber(), newExpiration);
        }

        public void showInfo() {
            System.out.println("  Owner              : " + owner);
            System.out.println("  Manufacturer       : " + manufacturer);
            System.out.println("  Model              : " + model);
            System.out.println("  Year               : " + year);
            System.out.println("  Car Age            : " + carAge() + " years");
            System.out.println("  Qualifies Renewal  : " + (qualifiesForRenewal() ? "Yes" : "No"));
            licensePlate.showInfo();
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // --- Create license plates ---
        LicensePlate plate1 = new LicensePlate("DHK-GA-1234", LocalDate.of(2023, 1, 1), LocalDate.of(2028, 1, 1));
        LicensePlate plate2 = new LicensePlate("DHK-KA-5678", LocalDate.of(2020, 6, 15), LocalDate.of(2025, 6, 15));
        LicensePlate plate3 = new LicensePlate("CTG-GA-9012", LocalDate.of(2015, 3, 10), LocalDate.of(2025, 3, 10));

        // --- Create cars ---
        Car car1 = new Car("Tareq", "Toyota", "Corolla", 2020, plate1);
        Car car2 = new Car("Afsana", "Honda", "Civic", 2018, plate2);
        Car car3 = new Car("Imtiaz", "Nissan", "Sunny", 2000, plate3);

        System.out.println("=== Car Information ===");
        car1.showInfo();
        car2.showInfo();
        car3.showInfo();

        // --- Validate license plates ---
        System.out.println("=== License Plate Validation ===");
        Car[] cars = {car1, car2, car3};
        for (Car car : cars) {
            String valid = car.licensePlate.isValid() ? "Valid" : "EXPIRED";
            System.out.printf("  %s %s (%s): %s%n",
                    car.manufacturer, car.model, car.licensePlate.getPlateNumber(), valid);
        }
        System.out.println();

        // --- Renew registration ---
        System.out.println("=== Registration Renewal ===");
        car1.renewRegistration(LocalDate.of(2033, 1, 1));
        car2.renewRegistration(LocalDate.of(2030, 6, 15));
        System.out.println();

        // --- Try to renew an old car ---
        System.out.println("=== Renew Old Car (should fail if too old) ===");
        Car oldCar = new Car("Robin", "Toyota", "Corona", 2003,
                new LicensePlate("DHK-GA-0001", LocalDate.of(2005, 1, 1), LocalDate.of(2025, 1, 1)));
        oldCar.showInfo();
        oldCar.renewRegistration(LocalDate.of(2030, 1, 1));
        System.out.println();

        // --- Very old car ---
        System.out.println("=== Very Old Car ===");
        Car ancientCar = new Car("Karim", "Datsun", "120Y", 1980,
                new LicensePlate("DHK-TA-0002", LocalDate.of(1985, 1, 1), LocalDate.of(2000, 1, 1)));
        ancientCar.showInfo();
        ancientCar.renewRegistration(LocalDate.of(2030, 1, 1));
    }
}
