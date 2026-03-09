/**
 * Practice 19: Vehicle Rental — Runtime Polymorphism
 * Task: Display brand, model, year, and rental price for one car,
 *       two bikes, and one truck if each is rented for 10 days.
 *
 * How to compile and run:
 *   javac Practice19.java
 *   java Practice19
 *
 * Key Concepts:
 *   - Upcasting (subclass → base class reference)
 *   - Runtime polymorphism (correct method called at runtime)
 *   - Iterating an array of base class references
 */

public class Practice19 {

    /** Base class for all rental vehicles. */
    static class Vehicle {
        protected String brand;
        protected String model;
        protected int year;

        public Vehicle(String brand, String model, int year) {
            this.brand = brand;
            this.model = model;
            this.year = year;
        }

        public double calculateRentalCost(int days) {
            return 0.0;
        }

        public String getVehicleType() {
            return "Vehicle";
        }

        @Override
        public String toString() {
            return brand + " " + model + " (" + year + ")";
        }
    }

    /** Car: $50/day, 10% discount if older than 5 years. */
    static class Car extends Vehicle {
        private static final double BASE_RATE = 50.0;

        public Car(String brand, String model, int year) {
            super(brand, model, year);
        }

        @Override
        public double calculateRentalCost(int days) {
            double cost = BASE_RATE * days;
            if (2026 - year > 5) cost *= 0.9;
            return cost;
        }

        @Override
        public String getVehicleType() { return "Car"; }

        @Override
        public String toString() { return "[Car] " + super.toString(); }
    }

    /** Bike: $15/day, 15% discount if rental > 7 days. */
    static class Bike extends Vehicle {
        private static final double BASE_RATE = 15.0;

        public Bike(String brand, String model, int year) {
            super(brand, model, year);
        }

        @Override
        public double calculateRentalCost(int days) {
            double cost = BASE_RATE * days;
            if (days > 7) cost *= 0.85;
            return cost;
        }

        @Override
        public String getVehicleType() { return "Bike"; }

        @Override
        public String toString() { return "[Bike] " + super.toString(); }
    }

    /** Truck: $100/day + $100/day additional fee. */
    static class Truck extends Vehicle {
        private static final double BASE_RATE = 100.0;
        private static final double ADDITIONAL_FEE = 100.0;

        public Truck(String brand, String model, int year) {
            super(brand, model, year);
        }

        @Override
        public double calculateRentalCost(int days) {
            return (BASE_RATE + ADDITIONAL_FEE) * days;
        }

        @Override
        public String getVehicleType() { return "Truck"; }

        @Override
        public String toString() { return "[Truck] " + super.toString(); }
    }

    public static void main(String[] args) {
        // Upcasting: all stored as Vehicle references
        Vehicle[] vehicles = {
            new Car("Toyota", "Corolla", 2019),     // 1 Car
            new Bike("Yamaha", "R15", 2023),         // 2 Bikes
            new Bike("Honda", "CBR", 2024),
            new Truck("Volvo", "FH16", 2022),        // 1 Truck
        };

        int rentalDays = 10;

        // Runtime polymorphism
        System.out.printf("=== Vehicle Rental Report (%d Days) ===%n", rentalDays);
        System.out.println();
        System.out.printf("  %-8s | %-10s | %-10s | %-6s | %12s%n",
                "Type", "Brand", "Model", "Year", "Rental Cost");
        System.out.println("  " + "-".repeat(60));

        double totalCost = 0;
        for (Vehicle vehicle : vehicles) {
            double cost = vehicle.calculateRentalCost(rentalDays); // Runtime polymorphism
            totalCost += cost;
            System.out.printf("  %-8s | %-10s | %-10s | %-6d | $%10.2f%n",
                    vehicle.getVehicleType(), vehicle.brand, vehicle.model,
                    vehicle.year, cost);
        }

        System.out.println("  " + "-".repeat(60));
        System.out.printf("  %-8s | %-10s | %-10s | %-6s | $%10.2f%n",
                "Total", "", "", "", totalCost);
        System.out.println();

        // Demonstrate runtime polymorphism
        System.out.println("=== Runtime Polymorphism ===");
        System.out.println();
        for (Vehicle vehicle : vehicles) {
            System.out.println("  " + vehicle);
            System.out.println("    Type at runtime : " + vehicle.getClass().getSimpleName());
            System.out.printf("    Rental (10 days): $%.2f%n", vehicle.calculateRentalCost(10));
            System.out.println();
        }
    }
}
