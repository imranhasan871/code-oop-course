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
 * How to compile and run:
 *   javac Practice17.java
 *   java Practice17
 *
 * Key Concepts:
 *   - Inheritance (IS-A relationship)
 *   - Base class and subclasses
 *   - Shared properties via parent class
 */

public class Practice17 {

    /** Base class for all rental vehicles. */
    static class Vehicle {
        protected String brand;
        protected String model;
        protected int year;

        public void showInfo() {
            System.out.println("  Brand : " + brand);
            System.out.println("  Model : " + model);
            System.out.println("  Year  : " + year);
        }
    }

    /** Car: $50/day, 10% discount if older than 5 years. */
    static class Car extends Vehicle {
        private static final double BASE_RATE = 50.0;

        public Car(String brand, String model, int year) {
            this.brand = brand;
            this.model = model;
            this.year = year;
        }

        public double calculateRentalCost(int days) {
            double cost = BASE_RATE * days;
            if (2026 - year > 5) {
                cost *= 0.9;
            }
            return cost;
        }
    }

    /** Bike: $15/day, 15% discount if rental > 7 days. */
    static class Bike extends Vehicle {
        private static final double BASE_RATE = 15.0;

        public Bike(String brand, String model, int year) {
            this.brand = brand;
            this.model = model;
            this.year = year;
        }

        public double calculateRentalCost(int days) {
            double cost = BASE_RATE * days;
            if (days > 7) {
                cost *= 0.85;
            }
            return cost;
        }
    }

    /** Truck: $100/day + $100/day additional maintenance fee. */
    static class Truck extends Vehicle {
        private static final double BASE_RATE = 100.0;
        private static final double ADDITIONAL_FEE = 100.0;

        public Truck(String brand, String model, int year) {
            this.brand = brand;
            this.model = model;
            this.year = year;
        }

        public double calculateRentalCost(int days) {
            return (BASE_RATE + ADDITIONAL_FEE) * days;
        }
    }

    public static void main(String[] args) {
        Car car = new Car("Toyota", "Corolla", 2019);
        Bike bike = new Bike("Yamaha", "R15", 2023);
        Truck truck = new Truck("Volvo", "FH16", 2022);

        int days = 10;

        System.out.println("=== Car Rental ===");
        car.showInfo();
        System.out.printf("  Rental (%d days): $%.2f%n", days, car.calculateRentalCost(days));
        System.out.println();

        System.out.println("=== Bike Rental ===");
        bike.showInfo();
        System.out.printf("  Rental (%d days): $%.2f%n", days, bike.calculateRentalCost(days));
        System.out.println();

        System.out.println("=== Truck Rental ===");
        truck.showInfo();
        System.out.printf("  Rental (%d days): $%.2f%n", days, truck.calculateRentalCost(days));
        System.out.println();

        System.out.println("=== Discount Scenarios ===");

        Car oldCar = new Car("Honda", "Civic", 2018);
        System.out.println("Old Car (2018):");
        oldCar.showInfo();
        System.out.printf("  Rental (%d days): $%.2f%n", days, oldCar.calculateRentalCost(days));
        System.out.printf("  (10%% discount applied — car is %d years old)%n", 2026 - oldCar.year);
        System.out.println();

        Bike shortBike = new Bike("Honda", "CBR", 2024);
        int shortDays = 5;
        System.out.printf("Bike for %d days (no discount):%n", shortDays);
        shortBike.showInfo();
        System.out.printf("  Rental (%d days): $%.2f%n", shortDays, shortBike.calculateRentalCost(shortDays));
        System.out.println();

        Bike longBike = new Bike("Suzuki", "Gixxer", 2023);
        int longDays = 10;
        System.out.printf("Bike for %d days (15%% discount):%n", longDays);
        longBike.showInfo();
        System.out.printf("  Rental (%d days): $%.2f%n", longDays, longBike.calculateRentalCost(longDays));
    }
}
