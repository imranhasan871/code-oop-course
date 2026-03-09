/**
 * Practice 18: Vehicle Rental — Method Overriding & Constructor Chaining
 * Task: Enhance the vehicle rental system with method overriding and
 *       constructor chaining between base class and subclasses.
 *
 * Changes from Practice 17:
 *   - Constructor chaining: subclass constructors call super()
 *   - Method overriding: calculateRentalCost() in Vehicle, overridden in subclasses
 *   - toString() overriding in each class
 *
 * How to compile and run:
 *   javac Practice18.java
 *   java Practice18
 *
 * Key Concepts:
 *   - Method overriding (@Override)
 *   - Constructor chaining (super())
 *   - toString() overriding
 */

public class Practice18 {

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

        @Override
        public String toString() {
            return brand + " " + model + " (" + year + ")";
        }
    }

    /** Car: $50/day, 10% discount if older than 5 years. */
    static class Car extends Vehicle {
        private static final double BASE_RATE = 50.0;

        public Car(String brand, String model, int year) {
            super(brand, model, year);  // Constructor chaining
        }

        @Override
        public double calculateRentalCost(int days) {
            double cost = BASE_RATE * days;
            if (2026 - year > 5) {
                cost *= 0.9;
            }
            return cost;
        }

        @Override
        public String toString() {
            return "[Car] " + super.toString();
        }
    }

    /** Bike: $15/day, 15% discount if rental > 7 days. */
    static class Bike extends Vehicle {
        private static final double BASE_RATE = 15.0;

        public Bike(String brand, String model, int year) {
            super(brand, model, year);  // Constructor chaining
        }

        @Override
        public double calculateRentalCost(int days) {
            double cost = BASE_RATE * days;
            if (days > 7) {
                cost *= 0.85;
            }
            return cost;
        }

        @Override
        public String toString() {
            return "[Bike] " + super.toString();
        }
    }

    /** Truck: $100/day + $100/day additional fee. */
    static class Truck extends Vehicle {
        private static final double BASE_RATE = 100.0;
        private static final double ADDITIONAL_FEE = 100.0;

        public Truck(String brand, String model, int year) {
            super(brand, model, year);  // Constructor chaining
        }

        @Override
        public double calculateRentalCost(int days) {
            return (BASE_RATE + ADDITIONAL_FEE) * days;
        }

        @Override
        public String toString() {
            return "[Truck] " + super.toString();
        }
    }

    public static void main(String[] args) {
        Vehicle car = new Car("Toyota", "Corolla", 2019);
        Vehicle bike = new Bike("Yamaha", "R15", 2023);
        Vehicle truck = new Truck("Volvo", "FH16", 2022);

        int days = 10;

        System.out.println("=== Vehicle Rental — Method Overriding ===");
        System.out.println();

        for (Vehicle vehicle : new Vehicle[]{car, bike, truck}) {
            System.out.println("  Vehicle : " + vehicle);                                  // toString override
            System.out.printf("  Rental (%d days): $%.2f%n", days, vehicle.calculateRentalCost(days)); // Method override
            System.out.println();
        }

        System.out.println("=== String Representation (toString) ===");
        System.out.println("  Car   : " + car);
        System.out.println("  Bike  : " + bike);
        System.out.println("  Truck : " + truck);
        System.out.println();

        System.out.println("=== Discount Scenarios ===");

        Vehicle oldCar = new Car("Honda", "Civic", 2018);
        System.out.println("  " + oldCar);
        System.out.printf("  Age: %d years → 10%% discount applied%n", 2026 - oldCar.year);
        System.out.printf("  Rental (%d days): $%.2f%n", days, oldCar.calculateRentalCost(days));
        System.out.println();

        Vehicle shortBike = new Bike("Honda", "CBR", 2024);
        int shortDays = 5;
        System.out.println("  " + shortBike);
        System.out.printf("  Rental (%d days): $%.2f (no discount)%n", shortDays, shortBike.calculateRentalCost(shortDays));
        System.out.println();

        Vehicle longBike = new Bike("Suzuki", "Gixxer", 2023);
        int longDays = 10;
        System.out.println("  " + longBike);
        System.out.printf("  Rental (%d days): $%.2f (15%% discount)%n", longDays, longBike.calculateRentalCost(longDays));
    }
}
