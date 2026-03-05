/**
 * Practice 06: Car Rental System (OOAD)
 * Task: Model a car rental company with a fleet of vehicles.
 *       Customers can rent and return cars. Track availability and maintenance.
 *
 * How to compile and run:
 *   javac Practice06.java
 *   java Practice06
 *
 * Key Concepts:
 *   - Multiple classes working together (Car, Customer)
 *   - Object references between classes
 *   - State management (available, rented, maintenance)
 */

public class Practice06 {

    /** Car represents a vehicle in the rental fleet. */
    static class Car {
        private String licensePlate;
        private String brand;
        private String model;
        private String carType;
        private double dailyRate;
        private boolean isAvailable;
        private boolean needsMaintenance;

        /** Creates a new Car that is available for rent. */
        public Car(String licensePlate, String brand, String model, String carType, double dailyRate) {
            this.licensePlate = licensePlate;
            this.brand = brand;
            this.model = model;
            this.carType = carType;
            this.dailyRate = dailyRate;
            this.isAvailable = true;
            this.needsMaintenance = false;
        }

        /** Marks the car as rented (not available). Returns true if successful. */
        public boolean rent() {
            if (!isAvailable) {
                System.out.printf("  [Error] %s %s (%s) is not available for rent.%n", brand, model, licensePlate);
                return false;
            }
            if (needsMaintenance) {
                System.out.printf("  [Error] %s %s (%s) needs maintenance and cannot be rented.%n", brand, model, licensePlate);
                return false;
            }
            isAvailable = false;
            return true;
        }

        /** Marks the car as available again. */
        public void returnCar() {
            isAvailable = true;
            System.out.printf("  [OK] %s %s (%s) has been returned.%n", brand, model, licensePlate);
        }

        /** Marks the car as needing maintenance. */
        public void sendToMaintenance() {
            needsMaintenance = true;
            isAvailable = false;
            System.out.printf("  [OK] %s %s (%s) sent to maintenance.%n", brand, model, licensePlate);
        }

        /** Marks maintenance as done and makes car available. */
        public void completeMaintenance() {
            needsMaintenance = false;
            isAvailable = true;
            System.out.printf("  [OK] %s %s (%s) maintenance completed. Now available.%n", brand, model, licensePlate);
        }

        /** Prints car details. */
        public void showInfo() {
            String status = "Available";
            if (needsMaintenance) {
                status = "In Maintenance";
            } else if (!isAvailable) {
                status = "Rented";
            }
            System.out.printf("  [%s] %s %s | Type: %s | Rate: %.2f/day | Plate: %s%n",
                    status, brand, model, carType, dailyRate, licensePlate);
        }
    }

    /** Customer represents a person who can rent a car. */
    static class Customer {
        private String name;
        private String phone;
        private Car rentedCar;

        /** Creates a new Customer with no rented car. */
        public Customer(String name, String phone) {
            this.name = name;
            this.phone = phone;
            this.rentedCar = null;
        }

        /** Allows the customer to rent an available car. */
        public void rentCar(Car car) {
            if (rentedCar != null) {
                System.out.printf("  [Error] %s already has a rented car (%s %s).%n",
                        name, rentedCar.brand, rentedCar.model);
                return;
            }
            if (car.rent()) {
                rentedCar = car;
                System.out.printf("  [OK] %s rented %s %s (%s) at %.2f/day.%n",
                        name, car.brand, car.model, car.licensePlate, car.dailyRate);
            }
        }

        /** Allows the customer to return their rented car. */
        public void returnCar() {
            if (rentedCar == null) {
                System.out.printf("  [Error] %s has no car to return.%n", name);
                return;
            }
            rentedCar.returnCar();
            System.out.printf("  [OK] %s returned the car.%n", name);
            rentedCar = null;
        }

        /** Prints customer details. */
        public void showInfo() {
            System.out.printf("  Customer: %s | Phone: %s", name, phone);
            if (rentedCar != null) {
                System.out.printf(" | Rented: %s %s (%s)", rentedCar.brand, rentedCar.model, rentedCar.licensePlate);
            } else {
                System.out.print(" | Rented: None");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // --- Create fleet of cars ---
        Car car1 = new Car("DHK-1001", "Toyota", "Corolla", "Sedan", 3500);
        Car car2 = new Car("DHK-1002", "Honda", "CR-V", "SUV", 5000);
        Car car3 = new Car("DHK-1003", "Toyota", "HiAce", "Van", 7000);

        System.out.println("=== Fleet Status ===");
        car1.showInfo();
        car2.showInfo();
        car3.showInfo();
        System.out.println();

        // --- Create customers ---
        Customer cust1 = new Customer("Tareq", "01700-000001");
        Customer cust2 = new Customer("Afsana", "01700-000002");

        // --- Rent cars ---
        System.out.println("=== Renting Cars ===");
        cust1.rentCar(car1);
        cust2.rentCar(car2);
        System.out.println();

        System.out.println("=== After Renting ===");
        car1.showInfo();
        car2.showInfo();
        car3.showInfo();
        cust1.showInfo();
        cust2.showInfo();
        System.out.println();

        // --- Try to rent already rented car ---
        System.out.println("=== Try to rent an already rented car ===");
        cust2.rentCar(car1);
        System.out.println();

        // --- Return car ---
        System.out.println("=== Tareq returns car ===");
        cust1.returnCar();
        cust1.showInfo();
        car1.showInfo();
        System.out.println();

        // --- Send car to maintenance ---
        System.out.println("=== Send car1 to maintenance ===");
        car1.sendToMaintenance();
        car1.showInfo();
        System.out.println();

        // --- Try renting car in maintenance ---
        System.out.println("=== Try renting car in maintenance ===");
        cust1.rentCar(car1);
        System.out.println();

        // --- Complete maintenance ---
        System.out.println("=== Complete maintenance ===");
        car1.completeMaintenance();
        car1.showInfo();
    }
}
