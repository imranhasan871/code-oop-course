/**
 * Practice 14: Smart Parking System
 * Task: Manage vehicle entry, slot assignment, payment, and exit.
 *       Calculate parking charge based on duration.
 *       Update slot allocation and availability in real time.
 *
 * How to compile and run:
 *   javac Practice14.java
 *   java Practice14
 *
 * Key Concepts:
 *   - Association between Vehicle, ParkingSlot, and ParkingLot
 *   - Duration-based charge calculation
 *   - Real-time availability tracking
 */

import java.time.LocalDateTime;
import java.time.Duration;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.List;

public class Practice14 {

    /** Vehicle with license plate and type. */
    static class Vehicle {
        private String licensePlate;
        private String vehicleType;

        public Vehicle(String licensePlate, String vehicleType) {
            this.licensePlate = licensePlate;
            this.vehicleType = vehicleType;
        }

        public String getLicensePlate() { return licensePlate; }
        public String getVehicleType() { return vehicleType; }
    }

    /** ParkingSlot that can hold one vehicle at a time. */
    static class ParkingSlot {
        private String slotId;
        private String slotType;
        private double ratePerHour;
        private Vehicle vehicle;
        private LocalDateTime entryTime;

        public ParkingSlot(String slotId, String slotType, double ratePerHour) {
            this.slotId = slotId;
            this.slotType = slotType;
            this.ratePerHour = ratePerHour;
        }

        public boolean isAvailable() { return vehicle == null; }
        public String getSlotId() { return slotId; }
        public String getSlotType() { return slotType; }
        public double getRatePerHour() { return ratePerHour; }
        public Vehicle getVehicle() { return vehicle; }
        public LocalDateTime getEntryTime() { return entryTime; }

        public void assignVehicle(Vehicle vehicle, LocalDateTime entryTime) {
            this.vehicle = vehicle;
            this.entryTime = entryTime;
        }

        public void releaseVehicle() {
            this.vehicle = null;
            this.entryTime = null;
        }

        public void showInfo() {
            String status = isAvailable()
                    ? "Available"
                    : "Occupied by " + vehicle.getLicensePlate();
            System.out.printf("  %-8s | Type: %-10s | Rate: %.0f/hr | %s%n",
                    slotId, slotType, ratePerHour, status);
        }
    }

    /** ParkingLot manages a collection of ParkingSlots. */
    static class ParkingLot {
        private String name;
        private List<ParkingSlot> slots = new ArrayList<>();
        private static final DateTimeFormatter FMT = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm");

        public ParkingLot(String name) {
            this.name = name;
        }

        public void addSlot(ParkingSlot slot) {
            slots.add(slot);
        }

        public ParkingSlot findAvailableSlot(String slotType) {
            for (ParkingSlot slot : slots) {
                if (slot.getSlotType().equals(slotType) && slot.isAvailable()) {
                    return slot;
                }
            }
            return null;
        }

        public ParkingSlot parkVehicle(Vehicle vehicle, LocalDateTime entryTime) {
            ParkingSlot slot = findAvailableSlot(vehicle.getVehicleType());
            if (slot == null) {
                System.out.printf("  [Error] No available %s slot for %s.%n",
                        vehicle.getVehicleType(), vehicle.getLicensePlate());
                return null;
            }
            slot.assignVehicle(vehicle, entryTime);
            System.out.printf("  [OK] %s parked in slot %s at %s.%n",
                    vehicle.getLicensePlate(), slot.getSlotId(), entryTime.format(FMT));
            return slot;
        }

        public double exitVehicle(String licensePlate, LocalDateTime exitTime) {
            for (ParkingSlot slot : slots) {
                if (!slot.isAvailable() && slot.getVehicle().getLicensePlate().equals(licensePlate)) {
                    Duration duration = Duration.between(slot.getEntryTime(), exitTime);
                    double hours = duration.toMinutes() / 60.0;
                    if (hours < 1) hours = 1; // Minimum 1 hour
                    double charge = hours * slot.getRatePerHour();
                    System.out.printf("  [OK] %s exiting slot %s.%n", licensePlate, slot.getSlotId());
                    System.out.printf("       Duration : %.1f hours%n", hours);
                    System.out.printf("       Rate     : %.0f/hr%n", slot.getRatePerHour());
                    System.out.printf("       Charge   : %.2f%n", charge);
                    slot.releaseVehicle();
                    return charge;
                }
            }
            System.out.printf("  [Error] Vehicle %s not found in parking lot.%n", licensePlate);
            return 0;
        }

        public void showStatus() {
            long available = slots.stream().filter(ParkingSlot::isAvailable).count();
            long occupied = slots.size() - available;
            System.out.println("  Parking Lot: " + name);
            System.out.printf("  Total Slots: %d | Available: %d | Occupied: %d%n",
                    slots.size(), available, occupied);
            System.out.println();
            for (ParkingSlot slot : slots) {
                slot.showInfo();
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        ParkingLot lot = new ParkingLot("City Center Parking");
        lot.addSlot(new ParkingSlot("A-01", "Car", 50));
        lot.addSlot(new ParkingSlot("A-02", "Car", 50));
        lot.addSlot(new ParkingSlot("A-03", "Car", 50));
        lot.addSlot(new ParkingSlot("B-01", "Bike", 20));
        lot.addSlot(new ParkingSlot("B-02", "Bike", 20));
        lot.addSlot(new ParkingSlot("C-01", "Truck", 100));

        System.out.println("=== Initial Parking Status ===");
        lot.showStatus();

        LocalDateTime now = LocalDateTime.of(2026, 3, 7, 9, 0);
        System.out.println("=== Parking Vehicles ===");
        lot.parkVehicle(new Vehicle("DHK-1234", "Car"), now);
        lot.parkVehicle(new Vehicle("DHK-5678", "Car"), now.plusMinutes(15));
        lot.parkVehicle(new Vehicle("DHK-9012", "Bike"), now.plusMinutes(30));
        lot.parkVehicle(new Vehicle("CTG-3456", "Truck"), now.plusHours(1));
        System.out.println();

        System.out.println("=== After Parking ===");
        lot.showStatus();

        System.out.println("=== Vehicle Exits ===");
        lot.exitVehicle("DHK-1234", now.plusHours(3));
        System.out.println();
        lot.exitVehicle("DHK-9012", now.plusHours(1).plusMinutes(30));
        System.out.println();
        lot.exitVehicle("CTG-3456", now.plusHours(5));
        System.out.println();

        System.out.println("=== Final Parking Status ===");
        lot.showStatus();

        System.out.println("=== Parking When Slots Are Limited ===");
        lot.parkVehicle(new Vehicle("SYL-7777", "Car"), now.plusHours(4));
        lot.parkVehicle(new Vehicle("SYL-8888", "Car"), now.plusHours(4));
        lot.parkVehicle(new Vehicle("SYL-9999", "Car"), now.plusHours(4));
    }
}
